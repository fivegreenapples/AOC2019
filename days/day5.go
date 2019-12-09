package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day5Part1(in string) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(r.verbose)

	_, out := vm.RunSlice([]int{1})

	return strconv.Itoa(out[len(out)-1])
}

func (r *Runner) Day5Part2(in string) string {

	testProgram := utils.CsvToInts(in)
	vm := intcode.New(testProgram)
	vm.SetDebug(r.verbose)

	_, out := vm.RunSlice([]int{5})

	return strconv.Itoa(out[0])
}
