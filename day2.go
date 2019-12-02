package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day2Part1(in string) string {
	program := csvToInts(in)
	result := d2Execute(program)
	return fmt.Sprintf("%d", result)
}

func d2Execute(program []int) int {

	pc := 0

	for {
		// fmt.Println(printProgram(program, pc))

		opcode := program[pc]
		switch opcode {
		case 1:
			program[program[pc+3]] = program[program[pc+1]] + program[program[pc+2]]
		case 2:
			program[program[pc+3]] = program[program[pc+1]] * program[program[pc+2]]
		case 99:
			return program[0]
		default:
			panic(fmt.Errorf("unhandled opcode: %d at position %d", opcode, pc))
		}

		pc += 4
	}

}

func day2Part2(in string) string {
	program := csvToInts(in)

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
