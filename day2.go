package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func init() {
	registerPart1(2, day2Part1)
	registerPart2(2, day2Part2)
}

func day2Part1(in string, verbose bool) string {
	program := utils.CsvToInts(in)
	result := d2Execute(program)
	return fmt.Sprintf("%d", result)
}

func d2Execute(program []int) int {

	vm := intcode.New(program)
	vm.Run(nil, nil)
	return vm.Read(0)

}

func day2Part2(in string, verbose bool) string {
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
