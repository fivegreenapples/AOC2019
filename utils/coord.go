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
func (c Coord) Sub(cc Coord) Coord {
	return Coord{
		c.X - cc.X,
		c.Y - cc.Y,
	}
}

func (c Coord) Simplify() Coord {

	absX := AbsInt(c.X)
	absY := AbsInt(c.Y)

	if absX == 0 && absY == 0 {
		return c
	}
	if absX == 0 {
		return Coord{0, c.Y / absY}
	}
	if absY == 0 {
		return Coord{c.X / absX, 0}
	}

	lcf := LargestCommonFactor(absX, absY)
	return Coord{c.X / lcf, c.Y / lcf}

}
