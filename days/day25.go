package days

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/fivegreenapples/AOC2019/intcode"
)

func (r *Runner) Day25Part1(in string) string {
	droid := intcode.NewFromString(in)

	input := make(chan int)
	output := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for char := range output {
			fmt.Print(string(rune(char)))
		}
		wg.Done()
	}()

	go func() {

		instructions := []string{
			"west",
			"take whirled peas",
			"east",
			"south",
			"west",
			"take bowl of rice",
			"east",
			"east",
			"take mutex",
			"east",
			"take astronaut ice cream",
			"east",
			"take ornament",
			"west",
			"south",
			"take tambourine",
			"north",
			"west",
			"south",
			"east",
			"take mug",
			"west",
			"south",
			"west",
			"south",
			"take easter egg",
			"west",
			"drop bowl of rice",        // heavier
			"drop easter egg",          // heavier
			"drop tambourine",          // heavier
			"drop astronaut ice cream", // heavier
			"drop mug",                 // lighter
			"drop mutex",               // heavier
			"drop whirled peas",        // lighter
			"drop ornament",            // lighter
		}

		items := map[int]string{
			0b00000001: "bowl of rice",
			0b00000010: "tambourine",
			0b00000100: "astronaut ice cream",
			0b00001000: "easter egg",
			0b00010000: "mug",
			0b00100000: "mutex",
			0b01000000: "whirled peas",
			0b10000000: "ornament",
		}

		for c := 0; c <= 255; c++ {
			thisItems := []string{}
			for k, v := range items {
				if c&k > 0 {
					thisItems = append(thisItems, v)
				}
			}

			for _, thisItem := range thisItems {
				instructions = append(instructions, "take "+thisItem)
			}
			instructions = append(instructions, "west")
			instructions = append(instructions, "inv")
			for _, thisItem := range thisItems {
				instructions = append(instructions, "drop "+thisItem)
			}
		}

		for _, instr := range instructions {
			for _, ch := range instr {
				input <- int(ch)
			}
			input <- '\n'
		}

		for {
			reader := bufio.NewReader(os.Stdin)
			text, _ := reader.ReadString('\n')
			for _, char := range text {
				input <- int(char)
			}
		}
	}()

	wg.Add(1)
	go func() {
		droid.Run(input, output)
		close(output)
		wg.Done()
	}()

	wg.Wait()
	return ""
}

func (r *Runner) Day25Part2(in string) string {
	return ""
}
