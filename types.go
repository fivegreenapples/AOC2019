package main

import "math"

type coord struct {
	x int
	y int
}

func (c coord) manhattan() int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

func (c coord) add(cc coord) coord {
	return coord{
		c.x + cc.x,
		c.y + cc.y,
	}
}
