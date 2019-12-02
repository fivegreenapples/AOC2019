package main

import "testing"

func TestDay2Part1(t *testing.T) {

	testInputs := map[string]string{
		"1,9,10,3,2,3,11,0,99,30,40,50": "3500",
		"1,0,0,0,99":                    "2",
		"2,3,0,3,99":                    "2",
		"2,4,4,5,99,0":                  "2",
		"1,1,1,4,99,5,6,0,99":           "30",
	}

	for in, expectedOut := range testInputs {
		out := day2Part1(in)
		if out != expectedOut {
			t.Errorf("day1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}