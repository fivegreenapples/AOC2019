package main

import (
	"fmt"
	"strconv"
	"strings"
)

func init() {
	registerPart1(4, day4Part1)
	registerPart2(4, day4Part2)
}

func day4Part1(in string, verbose bool) string {
	passRange := stringsToInts(strings.Split(in, "-"))
	min := password(passRange[0])
	max := password(passRange[1])

	min.makeValidMultipleSameDigits()

	numPasswords := 0
	for p := min; p < max; {
		numPasswords++
		if verbose {
			fmt.Println(numPasswords, p)
		}
		p.add(1)
		p.makeValidMultipleSameDigits()
	}

	return strconv.Itoa(numPasswords)
}

func day4Part2(in string, verbose bool) string {
	passRange := stringsToInts(strings.Split(in, "-"))
	min := password(passRange[0])
	max := password(passRange[1])

	min.makeValidDoubleDigit()

	numPasswords := 0
	for p := min; p < max; {
		numPasswords++
		if verbose {
			fmt.Println(numPasswords, p)
		}
		p.add(1)
		p.makeValidDoubleDigit()
	}

	return strconv.Itoa(numPasswords)
}
