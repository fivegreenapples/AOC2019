package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	day := flag.Int("d", 0, "Day of Advent")
	input := flag.String("i", "", "Input file")
	flag.Parse()

	if *day <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	var puzzleInput string
	if *input != "" {
		puzzleInputBytes, err := ioutil.ReadFile(*input)
		if err != nil {
			fmt.Printf("Error: couldn't read input file: %v\n", err)
			os.Exit(2)
		}
		puzzleInput = string(puzzleInputBytes)
	}

	result1, result2 := "", ""
	switch *day {
	case 1:
		result1 = day1Part1(puzzleInput)
		result2 = day1Part2(puzzleInput)
	default:
		fmt.Printf("Error: day %d not yet implemented\n", *day)
		os.Exit(2)

	}

	fmt.Println(result1)
	fmt.Println(result2)
}
