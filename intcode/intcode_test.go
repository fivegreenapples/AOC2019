package intcode

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/fivegreenapples/AOC2019/utils"
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

		vm := New(program)
		vm.Run(nil, nil)

		out := vm.Read(0)

		if out != expectedOut {
			t.Errorf("failed with %s. Expected %d, got %d", in, expectedOut, out)
		}
	}

}

func TestInputOutput(t *testing.T) {

	// following program should return input as output
	program := []int{3, 0, 4, 0, 99}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	rand.Seed(time.Now().UnixNano())
	iterations := 20
	for iterations > 0 {
		iterations--

		inputString := strconv.Itoa(rand.Int())
		var output strings.Builder

		vm.Run(strings.NewReader(inputString), &output)

		if inputString != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected same as input but got %s", inputString, output.String())
		}

	}

}

//
func TestParameterModes(t *testing.T) {

	// following program should put 99 into addr 4
	program := []int{1002, 4, 3, 4, 33}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	vm.Run(nil, nil)
	out := vm.Read(4)

	if out != 99 {
		t.Errorf("expected 99 at address 4, got %d", out)
	}

}

func TestEqualPositionMode(t *testing.T) {

	program := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-1": "0",
		"1":  "0",
		"8":  "1",
		"18": "0",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestEqualImmediateMode(t *testing.T) {

	program := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-1": "0",
		"1":  "0",
		"8":  "1",
		"18": "0",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestLessThanPositionMode(t *testing.T) {

	program := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-1": "1",
		"1":  "1",
		"8":  "0",
		"18": "0",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestLessThanImmediateMode(t *testing.T) {

	program := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-1": "1",
		"1":  "1",
		"8":  "0",
		"18": "0",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestJumpPositionMode(t *testing.T) {

	program := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-123": "1",
		"-13":  "1",
		"0":    "0",
		"13":   "1",
		"123":  "1",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestJumpImmediateMode(t *testing.T) {

	program := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-123": "1",
		"-13":  "1",
		"0":    "0",
		"13":   "1",
		"123":  "1",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}

func TestComparisons(t *testing.T) {

	program := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	vm := New(program)
	vm.SetDebug(testing.Verbose())

	testInputs := map[string]string{
		"-123":    "999",
		"-13":     "999",
		"0":       "999",
		"3":       "999",
		"8":       "1000",
		"13":      "1001",
		"133":     "1001",
		"1323":    "1001",
		"1324234": "1001",
	}

	for in, expectedOut := range testInputs {

		var output strings.Builder
		vm.Run(strings.NewReader(in), &output)

		if expectedOut != strings.TrimSpace(output.String()) {
			t.Errorf("failed with %s. Expected %s but got %s", in, expectedOut, output.String())
		}

	}

}
