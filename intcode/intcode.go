package intcode

import "io"

import "fmt"

type VM struct {
	program []int
	pc      int
}

func New(program []int) *VM {
	return &VM{
		program: program,
	}
}

func (vm *VM) Read(addr int) int {
	return vm.program[addr]
}

func (vm *VM) Run(input io.Reader, output io.Writer) {

	vm.pc = 0
	var opcode, operation int
	var modeA, modeB int
	var pA, pB int

	for {

		opcode = vm.program[vm.pc]
		operation = opcode % 100
		modeA = (opcode / 100) % 10
		modeB = (opcode / 1000) % 10

		switch operation {
		case 1, 2:
			pA = vm.program[vm.pc+1]
			pB = vm.program[vm.pc+2]
			if modeA == 0 {
				pA = vm.program[pA]
			}
			if modeB == 0 {
				pB = vm.program[pB]
			}
		}

		switch operation {
		case 1:
			vm.program[vm.program[vm.pc+3]] = pA + pB
			vm.pc += 4
		case 2:
			vm.program[vm.program[vm.pc+3]] = pA * pB
			vm.pc += 4
		case 99:
			return
		default:
			panic(fmt.Errorf("unhandled operation %d (in opcode %d) at position %d", operation, opcode, vm.pc))
		}

	}

}
