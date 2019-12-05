package days

import "fmt"
import "github.com/fivegreenapples/AOC2019/utils"

func (r *Runner) Day1Part1(in string) string {
	modules := utils.LinesAsInts(in)

	fuel := 0
	for _, m := range modules {
		fuel += fuelForMass(m)
	}

	return fmt.Sprintf("%d", fuel)
}

func (r *Runner) Day1Part2(in string) string {
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
