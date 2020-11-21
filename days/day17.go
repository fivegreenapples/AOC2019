package days

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day17Part1(in string) string {
	ascii := intcode.NewFromString(in)

	input := make(chan int)
	output := make(chan int)

	viewMap := map[utils.Coord]int{}

	go func() {
		ascii.Run(input, output)
		close(output)
	}()

	x, y := 0, 0
	calibrationSum := 0
	for {
		cur, valid := <-output
		if !valid {
			break
		}
		viewMap[utils.Coord{X: x, Y: y}] = cur

		if x >= 3 && y >= 2 {
			if viewMap[utils.Coord{X: x - 2, Y: y}] == '#' &&
				viewMap[utils.Coord{X: x - 1, Y: y}] == '#' &&
				viewMap[utils.Coord{X: x, Y: y}] == '#' &&
				viewMap[utils.Coord{X: x - 1, Y: y - 1}] == '#' {

				// Must be intersection at {x - 1, y}
				viewMap[utils.Coord{X: x - 1, Y: y}] = 'O'
				if r.verbose {
					fmt.Print("\bO")
				}
				calibrationSum += (x - 1) * y

			}
		}

		if r.verbose {
			fmt.Print(string(rune(viewMap[utils.Coord{X: x, Y: y}])))
		}

		if cur == '\n' {
			y++
			x = 0
		} else {
			x++
		}

	}

	return strconv.Itoa(calibrationSum)

}

func (r *Runner) Day17Part2(in string) string {
	program := utils.CsvToInts(in)
	program[0] = 2
	ascii := intcode.New(program)

	input := make(chan int)
	output := make(chan int)

	go func() {

		/*
			L4 L4 L10 R4 (A)
			R4 L4 L4 R8 R10 (B)
			L4 L4 L10 R4 (A)
			R4 L10 R10 (C)
			L4 L4 L10 R4 (A)
			R4 L10 R10 (C)
			R4 L4 L4 R8 R10 (B)
			R4 L10 R10 (C)
			R4 L10 R10 (C)
			R4 L4 L4 R8 R10 (B)
		*/
		mainMovement := "A,B,A,C,A,C,B,C,C,B"
		funcA := "L,4,L,4,L,10,R,4"
		funcB := "R,4,L,4,L,4,R,8,R,10"
		funcC := "R,4,L,10,R,10"
		videoFeed := "n"
		if r.verbose {
			videoFeed = "y"
		}

		instructions := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", mainMovement, funcA, funcB, funcC, videoFeed)
		for _, char := range instructions {
			input <- int(char)
		}

	}()

	go func() {
		ascii.Run(input, output)
		close(output)
	}()

	var dust int
	var last int
	for {
		current, valid := <-output
		if !valid {
			break
		}
		if last == '\n' && current == '\n' {
			time.Sleep(time.Millisecond * 50)
		}
		last = current

		if r.verbose && current <= 127 {
			fmt.Print(string(rune(current)))
		}
		if current > 127 {
			dust = current
		}
	}

	return strconv.Itoa(dust)

}
