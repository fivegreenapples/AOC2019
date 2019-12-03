package main

import "testing"

func TestDay4Part1(t *testing.T) {

	testInputs := map[string]string{
		"A": "A",
	}

	for in, expectedOut := range testInputs {
		out := day4Part1(in, testing.Verbose())
		if out != expectedOut {
			t.Errorf("day4 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
