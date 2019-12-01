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
	lineStrings := lines(in)
	ints := []int{}
	for _, l := range lineStrings {
		thisInt, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		ints = append(ints, thisInt)
	}
	return ints
}
