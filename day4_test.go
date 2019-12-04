package main

import "testing"

func TestDay4Part1(t *testing.T) {

	testInputs := map[int]int{
		111111: 111111,
		223450: 223455,
		223459: 223459,
		123789: 123799,
		123788: 123788,
		367479: 367777,
	}

	for in, expectedValid := range testInputs {
		p := password(in)
		p.makeValidMultipleSameDigits()

		if expectedValid != int(p) {
			t.Errorf("password validation failed. Expected next valid on or after %v to be %v. Got %v", in, expectedValid, p)
		}
	}

}
func TestDay4Part2(t *testing.T) {

	testInputs := map[int]int{
		112233: 112233,
		123444: 123445,
		111122: 111122,
	}

	for in, expectedValid := range testInputs {
		p := password(in)
		p.makeValidDoubleDigit()

		if expectedValid != int(p) {
			t.Errorf("password validation failed. Expected next valid on or after %v to be %v. Got %v", in, expectedValid, p)
		}
	}

}
