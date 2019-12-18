package intcode

import "fmt"

import "sync"

import "github.com/fivegreenapples/AOC2019/utils"

type VM struct {
	src   []int
	debug bool
}

type Core struct {
	ram map[int]int
}

func New(program []int) *VM {
	return &VM{
		src: program,
	}
}

func NewFromString(program string) *VM {
	return New(utils.CsvToInts(program))
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
		ram: make(map[int]int, len(vm.src)),
	}
	for i, data := range vm.src {
		core.ram[i] = data
	}

	pc := 0
	relativeBase := 0
	var opcode, operation int
	var modeA, modeB, modeC int
	var pA, pB, pC int

	for {

		opcode = core.ram[pc]
		operation = opcode % 100
		modeA = (opcode / 100) % 10
		modeB = (opcode / 1000) % 10
		modeC = (opcode / 10000) % 10

		if vm.debug {
			// fmt.Println(opcode, operation)
		}

		switch operation {
		case 1, 2, 7, 8:
			// pC is always an address of where to write
			if modeC == 0 {
				pC = core.ram[pc+3]
			} else if modeC == 1 {
				panic("immediate mode not supported for opcode pC")
			} else if modeC == 2 {
				pC = relativeBase + core.ram[pc+3]
			}
			fallthrough
		case 5, 6:
			if modeB == 0 {
				pB = core.ram[core.ram[pc+2]]
			} else if modeB == 1 {
				pB = core.ram[pc+2]
			} else if modeB == 2 {
				pB = core.ram[relativeBase+core.ram[pc+2]]
			}
			fallthrough
		case 4, 9:
			if modeA == 0 {
				pA = core.ram[core.ram[pc+1]]
			} else if modeA == 1 {
				pA = core.ram[pc+1]
			} else if modeA == 2 {
				pA = core.ram[relativeBase+core.ram[pc+1]]
			}
		case 3:
			// for opcode 3 (input), pA becomes the address of where to write
			if modeA == 0 {
				pA = core.ram[pc+1]
			} else if modeA == 1 {
				panic("immediate mode not supported for opcode 3")
			} else if modeA == 2 {
				pA = relativeBase + core.ram[pc+1]
			}
		}

		switch operation {
		case 1:
			core.ram[pC] = pA + pB
			pc += 4
		case 2:
			core.ram[pC] = pA * pB
			pc += 4
		case 3:
			in := <-input
			core.ram[pA] = in
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
				core.ram[pC] = 1
			} else {
				core.ram[pC] = 0
			}
			pc += 4
		case 8:
			if pA == pB {
				core.ram[pC] = 1
			} else {
				core.ram[pC] = 0
			}
			pc += 4
		case 9:
			relativeBase += pA
			pc += 2
		case 99:
			return &core
		default:
			panic(fmt.Errorf("unhandled operation %d (in opcode %d) at position %d", operation, opcode, pc))
		}

	}

}
