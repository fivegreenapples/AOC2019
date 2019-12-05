package utils

import "math"

type Coord struct {
	X int
	Y int
}

func (c Coord) Manhattan() int {
	return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)))
}

func (c Coord) Add(cc Coord) Coord {
	return Coord{
		c.X + cc.X,
		c.Y + cc.Y,
	}
}
