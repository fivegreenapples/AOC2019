package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var part1Registry map[int]func(string, bool) string
var part2Registry map[int]func(string, bool) string

func registerPart1(day int, impl func(string, bool) string) {
	if part1Registry == nil {
		part1Registry = make(map[int]func(string, bool) string)
	}
	part1Registry[day] = impl
}
func registerPart2(day int, impl func(string, bool) string) {
	if part2Registry == nil {
		part2Registry = make(map[int]func(string, bool) string)
	}
	part2Registry[day] = impl
}

func main() {

	day := flag.Int("d", 0, "Day of Advent")
	verbose := flag.Bool("v", false, "Verbosity")
	part := flag.Int("p", 0, "Part")
	input := flag.String("i", "", "Input file")
	flag.Parse()

	if *day <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	var puzzleInput string

	if *input == "" {
		// with no input, we attempt to auto find the input file.
		tryFile := "inputs/day" + strconv.Itoa(*day) + ".txt"
		_, err := os.Stat(tryFile)
		if err == nil {
			*input = tryFile
		}
	}

	if *input != "" {
		puzzleInputBytes, err := ioutil.ReadFile(*input)
		if err != nil {
			fmt.Printf("Error: couldn't read input file: %v\n", err)
			os.Exit(2)
		}
		puzzleInput = string(puzzleInputBytes)
	}

	impl1, found1 := part1Registry[*day]
	impl2, found2 := part2Registry[*day]

	if !found1 {
		fmt.Printf("Day %d not yet implemented\n", *day)
		os.Exit(3)
	}

	if *part == 0 || *part == 1 {
		fmt.Println(impl1(puzzleInput, *verbose))
	}
	if found2 && (*part == 0 || *part == 2) {
		fmt.Println(impl2(puzzleInput, *verbose))
	}
}
