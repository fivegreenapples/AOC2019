package days

import (
	"fmt"
	"sync"

	"github.com/fivegreenapples/AOC2019/intcode"
)

func (r *Runner) Day21Part1(in string) string {
	springdroid := intcode.NewFromString(in)

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		springdroid.Run(input, output)
		close(output)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for char := range output {
			if char <= 127 {
				fmt.Print(string(char))
			} else {
				fmt.Print(char)
			}
		}
		wg.Done()
	}()

	//
	// Following go routine used for manual input
	//
	// go func() {
	// 	for {
	// 		reader := bufio.NewReader(os.Stdin)
	// 		text, _ := reader.ReadString('\n')
	// 		for _, ch := range text {
	// 			input <- int(ch)
	// 		}
	// 	}
	// }()

	// strategy:
	// - jump if there is any hole in front
	// - except don't jump if would land in hole i.e. no jump if D == 0
	// - in other words walk forward if can see 4 ground ahead
	// XXX0 - 0
	// 0001 - 1
	// 0011 - 1
	// 0011 - 1
	// 0101 - 1
	// 0111 - 1
	// 1001 - 1
	// 1011 - 1
	// 1101 - 1
	// 1111 - 0
	// => !(A && B && C) && D
	// => (!A || !B || !C) && D
	//
	// Resulting porgram is:
	// NOT A T
	// OR T J
	// NOT B T
	// OR T J
	// NOT C T
	// OR T J
	// AND D J

	inputInstructions := []string{
		"NOT A T",
		"OR T J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",
		"AND D J",
		"WALK",
	}
	for _, instr := range inputInstructions {
		for _, ch := range instr {
			input <- int(ch)
		}
		input <- '\n'
	}

	wg.Wait()
	return ""
}

func (r *Runner) Day21Part2(in string) string {
	springdroid := intcode.NewFromString(in)

	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		springdroid.Run(input, output)
		close(output)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for char := range output {
			if char <= 127 {
				fmt.Print(string(char))
			} else {
				fmt.Print(char)
			}
		}
		wg.Done()
	}()

	//
	// Following go routine used for manual input
	//
	// go func() {
	// 	for {
	// 		reader := bufio.NewReader(os.Stdin)
	// 		text, _ := reader.ReadString('\n')
	// 		for _, ch := range text {
	// 			input <- int(ch)
	// 		}
	// 	}
	// }()

	// strategy:
	// - jump if there is any hole in front
	// - except don't jump if would land in hole i.e. no jump if D == 0
	// - in other words walk forward if can see 4 ground ahead
	// - above is same as part 1. part 2 is amended with:
	// - as above but only jump if possible to jump immediately again, or by stepping one then jump, or if there is a hole immediately in front
	// XXX0 - 0
	// 0001 - 1
	// 0011 - 1
	// 0011 - 1
	// 0101 - 1
	// 0111 - 1
	// 1001 - 1
	// 1011 - 1
	// 1101 - 1
	// 1111 - 0
	// => !(A && B && C) && D
	// additions for part 2:
	// => (!A || !B || !C) && D && (H || E&&I || !A)
	// => (!A || !B || !C) && D && !(!(H || E&&I) && A)
	//
	// Resulting program is:
	// NOT A T
	// OR T J
	// NOT B T
	// OR T J
	// NOT C T
	// OR T J
	// AND D J
	// NOT E T
	// NOT T T // double NOT puts E in T
	// AND I T
	// OR H T
	// NOT T T
	// AND A T
	// NOT T T
	// AND T J

	inputInstructions := []string{
		"NOT A T",
		"OR T J",
		"NOT B T",
		"OR T J",
		"NOT C T",
		"OR T J",
		"AND D J",
		"NOT E T",
		"NOT T T",
		"AND I T",
		"OR H T",
		"NOT T T",
		"AND A T",
		"NOT T T",
		"AND T J",
		"RUN",
	}
	for _, instr := range inputInstructions {
		for _, ch := range instr {
			input <- int(ch)
		}
		input <- '\n'
	}

	wg.Wait()
	return ""
}
