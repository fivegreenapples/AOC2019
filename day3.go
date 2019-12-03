package main

import (
	"fmt"
	"math"
	"strconv"
)

func init() {
	registerPart1(3, day3Part1)
	registerPart2(3, day3Part2)
}

func day3Part1(in string, verbose bool) string {

	wires := lines(in)
	wire1 := csvToStrings(wires[0])
	wire2 := csvToStrings(wires[1])

	wire1Pixels := pixellate(wire1)
	wire2Pixels := pixellate(wire2)

	if verbose {
		renderWires(wire1Pixels, wire2Pixels)
	}

	collisions := findCollisions(wire1Pixels, wire2Pixels)
	minDist := math.MaxInt64
	for _, c := range collisions {
		mh := c.manhattan()
		if mh > 0 && mh < minDist {
			minDist = mh
		}
	}

	return strconv.Itoa(minDist)
}

func day3Part2(in string, verbose bool) string {

	wires := lines(in)
	wire1 := csvToStrings(wires[0])
	wire2 := csvToStrings(wires[1])

	wire1Pixels := pixellateWithStepCount(wire1)
	wire2Pixels := pixellateWithStepCount(wire2)

	collisions := findCollisionsAndStepCounts(wire1Pixels, wire2Pixels)

	minSteps := math.MaxInt64
	for c, count := range collisions {

		if c.manhattan() == 0 {
			continue
		}

		if verbose {
			fmt.Println(c, count)
		}
		if count < minSteps {
			minSteps = count
		}
	}

	return strconv.Itoa(minSteps)
}

func findCollisionsAndStepCounts(wire1, wire2 map[coord]int) map[coord]int {
	collisions := map[coord]int{}
	for c, steps1 := range wire1 {
		if steps2, hit := wire2[c]; hit {
			collisions[c] = steps1 + steps2
		}
	}
	return collisions
}

func findCollisions(wire1, wire2 map[coord]bool) []coord {
	collisions := []coord{}
	for c := range wire1 {
		if _, hit := wire2[c]; hit {
			collisions = append(collisions, c)
		}
	}
	return collisions
}

func renderWires(wire1, wire2 map[coord]bool) {

	// find extents
	max, min := coord{math.MinInt64, math.MinInt64}, coord{math.MaxInt64, math.MaxInt64}
	for c := range wire1 {
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.x < min.x {
			min.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
	}
	for c := range wire2 {
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.x < min.x {
			min.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
	}

	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			p1 := wire1[coord{x, y}]
			p2 := wire2[coord{x, y}]
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

func pixellate(wire []string) map[coord]bool {

	pixels := map[coord]bool{}
	current := coord{0, 0}
	pixels[current] = true
	for _, next := range wire {

		unitDelta, len := convertToDeltaDetails(next)

		for len > 0 {
			current = current.add(unitDelta)
			pixels[current] = true
			len--
		}
	}

	return pixels
}

func pixellateWithStepCount(wire []string) map[coord]int {

	pixels := map[coord]int{}
	current := coord{0, 0}
	pixels[current] = 0
	steps := 0
	for _, next := range wire {

		unitDelta, len := convertToDeltaDetails(next)

		for len > 0 {
			steps++
			current = current.add(unitDelta)
			if _, seen := pixels[current]; !seen {
				// only overwrite steps count if we haven't visited here before
				pixels[current] = steps
			}
			len--
		}
	}

	return pixels
}

func convertToDeltaDetails(in string) (unitDelta coord, length int) {

	dir := in[0:1]
	len := in[1:]
	lenInt, err := strconv.Atoi(len)
	if err != nil {
		panic(fmt.Errorf("failed converting %s to vector", in))
	}

	switch dir {
	case "U":
		return coord{0, 1}, lenInt
	case "D":
		return coord{0, -1}, lenInt
	case "L":
		return coord{-1, 0}, lenInt
	case "R":
		return coord{1, 0}, lenInt
	}

	panic(fmt.Errorf("failed converting %s to vector", in))
}
