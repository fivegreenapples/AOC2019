package days

import (
	"fmt"
	"math"
	"strconv"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day3Part1(in string) string {

	wires := utils.Lines(in)
	wire1 := utils.CsvToStrings(wires[0])
	wire2 := utils.CsvToStrings(wires[1])

	wire1Pixels := pixellate(wire1)
	wire2Pixels := pixellate(wire2)

	if r.verbose {
		renderWires(wire1Pixels, wire2Pixels)
	}

	collisions := findCollisions(wire1Pixels, wire2Pixels)
	minDist := math.MaxInt64
	for _, c := range collisions {
		mh := c.Manhattan()
		if mh > 0 && mh < minDist {
			minDist = mh
		}
	}

	return strconv.Itoa(minDist)
}

func (r *Runner) Day3Part2(in string) string {

	wires := utils.Lines(in)
	wire1 := utils.CsvToStrings(wires[0])
	wire2 := utils.CsvToStrings(wires[1])

	wire1Pixels := pixellateWithStepCount(wire1)
	wire2Pixels := pixellateWithStepCount(wire2)

	collisions := findCollisionsAndStepCounts(wire1Pixels, wire2Pixels)

	minSteps := math.MaxInt64
	for c, count := range collisions {

		if c.Manhattan() == 0 {
			continue
		}

		if r.verbose {
			fmt.Println(c, count)
		}
		if count < minSteps {
			minSteps = count
		}
	}

	return strconv.Itoa(minSteps)
}

func findCollisionsAndStepCounts(wire1, wire2 map[utils.Coord]int) map[utils.Coord]int {
	collisions := map[utils.Coord]int{}
	for c, steps1 := range wire1 {
		if steps2, hit := wire2[c]; hit {
			collisions[c] = steps1 + steps2
		}
	}
	return collisions
}

func findCollisions(wire1, wire2 map[utils.Coord]bool) []utils.Coord {
	collisions := []utils.Coord{}
	for c := range wire1 {
		if _, hit := wire2[c]; hit {
			collisions = append(collisions, c)
		}
	}
	return collisions
}

func renderWires(wire1, wire2 map[utils.Coord]bool) {

	// find extents
	max, min := utils.Coord{math.MinInt64, math.MinInt64}, utils.Coord{math.MaxInt64, math.MaxInt64}
	for c := range wire1 {
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
	}
	for c := range wire2 {
		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
	}

	for y := max.Y; y >= min.Y; y-- {
		for x := min.X; x <= max.X; x++ {
			p1 := wire1[utils.Coord{x, y}]
			p2 := wire2[utils.Coord{x, y}]
			if x == 0 && y == 0 {
				fmt.Print("O")
			} else if p1 && p2 {
				fmt.Print("X")
			} else if p1 {
				fmt.Print("1")
			} else if p2 {
				fmt.Print("2")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}

	fmt.Println(max, min)

}

func pixellate(wire []string) map[utils.Coord]bool {

	pixels := map[utils.Coord]bool{}
	current := utils.Coord{0, 0}
	pixels[current] = true
	for _, next := range wire {

		unitDelta, len := convertToDeltaDetails(next)

		for len > 0 {
			current = current.Add(unitDelta)
			pixels[current] = true
			len--
		}
	}

	return pixels
}

func pixellateWithStepCount(wire []string) map[utils.Coord]int {

	pixels := map[utils.Coord]int{}
	current := utils.Coord{0, 0}
	pixels[current] = 0
	steps := 0
	for _, next := range wire {

		unitDelta, len := convertToDeltaDetails(next)

		for len > 0 {
			steps++
			current = current.Add(unitDelta)
			if _, seen := pixels[current]; !seen {
				// only overwrite steps count if we haven't visited here before
				pixels[current] = steps
			}
			len--
		}
	}

	return pixels
}

func convertToDeltaDetails(in string) (unitDelta utils.Coord, length int) {

	dir := in[0:1]
	len := in[1:]
	lenInt, err := strconv.Atoi(len)
	if err != nil {
		panic(fmt.Errorf("failed converting %s to vector", in))
	}

	switch dir {
	case "U":
		return utils.Coord{0, 1}, lenInt
	case "D":
		return utils.Coord{0, -1}, lenInt
	case "L":
		return utils.Coord{-1, 0}, lenInt
	case "R":
		return utils.Coord{1, 0}, lenInt
	}

	panic(fmt.Errorf("failed converting %s to vector", in))
}
