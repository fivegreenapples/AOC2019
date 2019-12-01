package main

import "fmt"

func day1Part1(in string) string {
	modules := linesAsInts(in)

	fuel := 0
	for _, m := range modules {
		fuel += fuelForMass(m)
	}

	return fmt.Sprintf("%d", fuel)
}

func day1Part2(in string) string {
	modules := linesAsInts(in)

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
