package main

import "fmt"
import "github.com/fivegreenapples/AOC2019/utils"

func init() {
	registerPart1(1, day1Part1)
	registerPart2(1, day1Part2)
}

func day1Part1(in string, verbose bool) string {
	modules := utils.LinesAsInts(in)

	fuel := 0
	for _, m := range modules {
		fuel += fuelForMass(m)
	}

	return fmt.Sprintf("%d", fuel)
}

func day1Part2(in string, verbose bool) string {
	modules := utils.LinesAsInts(in)

	fuel := 0
	for _, m := range modules {
		fuel += totalFuelForMass(m)
	}

	return fmt.Sprintf("%d", fuel)
}

func fuelForMass(mass int) int {
	return (mass / 3) - 2
}

func totalFuelForMass(mass int) int {
	fuel := 0
	f := fuelForMass(mass)
	for f > 0 {
		fuel += f
		f = fuelForMass(f)
	}
	return fuel
}
