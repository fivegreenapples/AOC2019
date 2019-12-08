package intcode

import "fmt"

import "sync"

type VM struct {
	src   []int
	debug bool
}

type Core struct {
	ram []int
}

func New(program []int) *VM {
	return &VM{
		src: program,
	}
}

func (c *Core) Read(addr int) int {
	return c.ram[addr]
}

func (vm *VM) SetDebug(val bool) {
	vm.debug = val
}

func (vm *VM) RunSlice(input []int) (*Core, []int) {

	inputBuffer := make(chan int, len(input))
	outputBuffer := make(chan int)

	// fill input buffer
	for _, v := range input {
		inputBuffer <- v
	}

	// Create wait group used to signal end of goroutines
	wait := sync.WaitGroup{}
	wait.Add(1)
	wait.Add(1)
	var core *Core
	go func() {
		core = vm.Run(inputBuffer, outputBuffer)
		close(outputBuffer)
		wait.Done()
	}()

	// create output slice and start goroutine to fill it from the outputBuffer
	output := []int{}
	go func() {
		for val := range outputBuffer {
			output = append(output, val)
		}
		wait.Done()
	}()

	wait.Wait()

	return core, output

}

func (vm *VM) Run(input chan int, output chan int) *Core {

	core := Core{
		ram: make([]int, len(vm.src)),
	}
	copy(core.ram, vm.src)

	pc := 0
	var opcode, operation int
	var modeA, modeB int
	var pA, pB int

	for {

		opcode = core.ram[pc]
		operation = opcode % 100
		modeA = (opcode / 100) % 10
		modeB = (opcode / 1000) % 10

		if vm.debug {
			// fmt.Println(opcode, operation)
		}

		switch operation {
		case 4:
			pA = core.ram[pc+1]
			if modeA == 0 {
				pA = core.ram[pA]
			}
		case 1, 2, 5, 6, 7, 8:
			pA = core.ram[pc+1]
			pB = core.ram[pc+2]
			if modeA == 0 {
				pA = core.ram[pA]
			}
			if modeB == 0 {
				pB = core.ram[pB]
			}
		}

		switch operation {
		case 1:
			core.ram[core.ram[pc+3]] = pA + pB
			pc += 4
		case 2:
			core.ram[core.ram[pc+3]] = pA * pB
			pc += 4
		// case 3:
		// 	var opIn int
		// 	fmt.Fscanf(input, "%d\n", &opIn)
		// 	core.ram[core.ram[pc+1]] = opIn
		// 	pc += 2
		// case 4:
		// 	fmt.Fprintf(output, "%d\n", pA)
		// 	pc += 2
		case 3:
			in := <-input
			core.ram[core.ram[pc+1]] = in
			if vm.debug {
				fmt.Println("Read input: ", in)
			}
			pc += 2
		case 4:
			if vm.debug {
				fmt.Println("Sending output: ", pA)
			}
			output <- pA
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
				core.ram[core.ram[pc+3]] = 1
			} else {
				core.ram[core.ram[pc+3]] = 0
			}
			pc += 4
		case 8:
			if pA == pB {
				core.ram[core.ram[pc+3]] = 1
			} else {
				core.ram[core.ram[pc+3]] = 0
			}
			pc += 4
		case 99:
			return &core
		default:
			panic(fmt.Errorf("unhandled operation %d (in opcode %d) at position %d", operation, opcode, pc))
		}

	}

}
