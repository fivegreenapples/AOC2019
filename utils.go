package main

import "strings"

import "strconv"

func lines(in string) []string {
	lines := strings.Split(in, "\n")
	trimmed := []string{}
	for _, l := range lines {
		trimmed = append(trimmed, strings.TrimSpace(l))
	}
	return trimmed
}

func linesAsInts(in string) []int {
	return stringsToInts(lines(in))
}

func csvToInts(in string) []int {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, ",")
	bits := strings.Split(in, ",")
	return stringsToInts(bits)
}

func stringsToInts(inStrings []string) []int {
	ints := []int{}
	for _, in := range inStrings {
		thisInt, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		ints = append(ints, thisInt)
	}
	return ints
}
