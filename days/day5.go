package days

import (
	"strings"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day5Part1(in string) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(r.verbose)

	input := strings.NewReader("1")
	var output strings.Builder

	vm.Run(input, &output)

	return output.String()
}

func (r *Runner) Day5Part2(in string) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(r.verbose)

	input := strings.NewReader("5")
	var output strings.Builder

	vm.Run(input, &output)

	return output.String()
}
