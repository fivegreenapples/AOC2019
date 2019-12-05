package intcode

import "io"

import "fmt"

type VM struct {
	src   []int
	ram   []int
	debug bool
}

func New(program []int) *VM {
	return &VM{
		src: program,
	}
}

func (vm *VM) Read(addr int) int {
	return vm.ram[addr]
}

func (vm *VM) SetDebug(val bool) {
	vm.debug = val
}

func (vm *VM) Run(input io.Reader, output io.Writer) {

	vm.ram = make([]int, len(vm.src))
	copy(vm.ram, vm.src)

	pc := 0
	var opcode, operation int
	var modeA, modeB int
	var pA, pB int

	for {

		opcode = vm.ram[pc]
		operation = opcode % 100
		modeA = (opcode / 100) % 10
		modeB = (opcode / 1000) % 10

		if vm.debug {
			fmt.Println(opcode, operation)
		}

		switch operation {
		case 4:
			pA = vm.ram[pc+1]
			if modeA == 0 {
				pA = vm.ram[pA]
			}
		case 1, 2, 5, 6, 7, 8:
			pA = vm.ram[pc+1]
			pB = vm.ram[pc+2]
			if modeA == 0 {
				pA = vm.ram[pA]
			}
			if modeB == 0 {
				pB = vm.ram[pB]
			}
		}

		switch operation {
		case 1:
			vm.ram[vm.ram[pc+3]] = pA + pB
			pc += 4
		case 2:
			vm.ram[vm.ram[pc+3]] = pA * pB
			pc += 4
		case 3:
			var opIn int
			fmt.Fscanf(input, "%d\n", &opIn)
			vm.ram[vm.ram[pc+1]] = opIn
			pc += 2
		case 4:
			fmt.Fprintf(output, "%d\n", pA)
			pc += 2
		case 5:
			if pA != 0 {
				pc = pB
			} else {
				pc += 3
			}
		case 6:
			if pA == 0 {
				pc = pB
			} else {
				pc += 3
			}
		case 7:
			if pA < pB {
				vm.ram[vm.ram[pc+3]] = 1
			} else {
				vm.ram[vm.ram[pc+3]] = 0
			}
			pc += 4
		case 8:
			if pA == pB {
				vm.ram[vm.ram[pc+3]] = 1
			} else {
				vm.ram[vm.ram[pc+3]] = 0
			}
			pc += 4
		case 99:
			return
		default:
			panic(fmt.Errorf("unhandled operation %d (in opcode %d) at position %d", operation, opcode, pc))
		}

	}

}
