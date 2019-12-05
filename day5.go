package main

import (
	"strings"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func init() {
	registerPart1(5, day5Part1)
	registerPart2(5, day5Part2)
}

func day5Part1(in string, verbose bool) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(verbose)

	input := strings.NewReader("1")
	var output strings.Builder

	vm.Run(input, &output)

	return output.String()
}

func day5Part2(in string, verbose bool) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(verbose)

	input := strings.NewReader("5")
	var output strings.Builder

	vm.Run(input, &output)

	return output.String()
}
