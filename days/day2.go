package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day2Part1(in string) string {
	program := utils.CsvToInts(in)
	program[1] = 12
	program[2] = 2
	result := d2Execute(program)
	return fmt.Sprintf("%d", result)
}

func d2Execute(program []int) int {

	vm := intcode.New(program)
	vm.Run(nil, nil)
	return vm.Read(0)

}

func (r *Runner) Day2Part2(in string) string {
	program := utils.CsvToInts(in)

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {

			testProgram := make([]int, len(program))
			copy(testProgram, program)
			testProgram[1] = noun
			testProgram[2] = verb

			result := d2Execute(testProgram)
			if result == 19690720 {
				return fmt.Sprintf("%d", (noun*100)+verb)
			}
		}
	}

	panic("no noun verb combo worked")
}

func printProgram(program []int, pc int) string {
	str := ""
	for i, p := range program {

		if i == pc {
			str += "["
		}
		str += strconv.Itoa(p)
		if i == pc {
			str += "]"
		}
		str += ","
	}

	str = strings.TrimRight(str, ",")

	if pc >= len(program) {
		str += fmt.Sprintf(" [?=%d]", pc)
	}

	return str
}
