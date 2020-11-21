package days

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day11Part1(in string) string {
	painted := goRobot(in, 0)
	return strconv.Itoa(len(painted))
}

func (r *Runner) Day11Part2(in string) string {
	painted := goRobot(in, 1)
	picture := &strings.Builder{}

	min, max := utils.ExtentsOfIntMap(painted)
	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x <= max.X; x++ {
			if 1 == painted[utils.Coord{X: x, Y: y}] {
				fmt.Fprint(picture, "#")
			} else {
				fmt.Fprint(picture, " ")
			}
		}
		fmt.Fprint(picture, "\n")
	}

	return picture.String()

}

func goRobot(in string, startColour int) map[utils.Coord]int {
	robotProgram := utils.CsvToInts(in)
	robot := intcode.New(robotProgram)

	input := make(chan int, 1)
	output := make(chan int)
	curPos := utils.Coord{X: 0, Y: 0}
	curBearing := utils.Coord{X: 0, Y: 1}
	painted := map[utils.Coord]int{}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for {
			colour, valid := <-output
			if !valid {
				break
			}
			direction, valid := <-output
			if !valid {
				break
			}

			painted[curPos] = colour

			if direction == 0 { // turn left 90
				curBearing.X, curBearing.Y = -curBearing.Y, curBearing.X
			} else if direction == 1 { // turn right 90
				curBearing.X, curBearing.Y = curBearing.Y, -curBearing.X
			} else {
				panic("odd direction")
			}

			curPos = curPos.Add(curBearing)

			input <- painted[curPos]
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		robot.Run(input, output)
		close(output)
		wg.Done()
	}()

	input <- startColour

	wg.Wait()

	return painted
}
