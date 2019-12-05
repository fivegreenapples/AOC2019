package intcode

import (
	"github.com/fivegreenapples/AOC2019/utils"
	"testing"
)

func TestAddMultiplyPositionMode(t *testing.T) {

	testInputs := map[string]int{
		"1,9,10,3,2,3,11,0,99,30,40,50": 3500,
		"1,0,0,0,99":                    2,
		"2,3,0,3,99":                    2,
		"2,4,4,5,99,0":                  2,
		"1,1,1,4,99,5,6,0,99":           30,
	}

	for in, expectedOut := range testInputs {

		program := utils.CsvToInts(in)

		computer := New(program)
		computer.Run(nil, nil)

		out := computer.Read(0)

		if out != expectedOut {
			t.Errorf("failed with %s. Expected %d, got %d", in, expectedOut, out)
		}
	}

}
