package main

import (
	"flag"
	"fmt"
	"github.com/fivegreenapples/AOC2019/days"
	"io/ioutil"
	"os"
	"strconv"
)

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

	runner := days.NewRunner(*verbose)

	if *part == 0 || *part == 1 {
		pt1Out, pt1err := runner.Run(*day, 1, puzzleInput)
		if pt1err != nil {
			fmt.Printf("%s\n", pt1err.Error())
			os.Exit(3)
		}
		fmt.Println(pt1Out)
	}
	if *part == 0 || *part == 2 {
		pt2Out, pt2err := runner.Run(*day, 2, puzzleInput)
		if pt2err == nil {
			fmt.Println(pt2Out)
		} else if *part == 2 {
			// only display error if we asked for part 2
			fmt.Printf("%s\n", pt2err.Error())
			os.Exit(3)
		}
	}
}
