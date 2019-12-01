package main

import "testing"

func TestDay1Part1(t *testing.T) {

	testInputs := map[string]string{
		"12":     "2",
		"14":     "2",
		"1969":   "654",
		"100756": "33583",
	}

	for in, expectedOut := range testInputs {
		out := day1Part1(in)
		if out != expectedOut {
			t.Errorf("day1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay1Part2(t *testing.T) {

	testInputs := map[string]string{
		"14":     "2",
		"1969":   "966",
		"100756": "50346",
	}

	for in, expectedOut := range testInputs {
		out := day1Part2(in)
		if out != expectedOut {
			t.Errorf("day1 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
