package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day9Part1(in string) string {
	boostProgram := utils.CsvToInts(in)
	vm := intcode.New(boostProgram)
	vm.SetDebug(r.verbose)

	_, out := vm.RunSlice([]int{1})

	return strconv.Itoa(out[0])
}

func (r *Runner) Day9Part2(in string) string {
	boostProgram := utils.CsvToInts(in)
	vm := intcode.New(boostProgram)
	vm.SetDebug(r.verbose)

	_, out := vm.RunSlice([]int{2})

	return strconv.Itoa(out[0])
}
