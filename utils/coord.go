package utils

import "math"

type Coord struct {
	X int
	Y int
}

var Origin = Coord{X: 0, Y: 0}

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

func ExtentsOfIntMap(in map[Coord]int) (min, max Coord) {
	min = Coord{math.MaxInt64, math.MaxInt64}
	max = Coord{math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}
func ExtentsOfBoolMap(in map[Coord]bool) (min, max Coord) {
	min = Coord{math.MaxInt64, math.MaxInt64}
	max = Coord{math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}
