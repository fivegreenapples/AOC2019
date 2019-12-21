package utils

import "math"

type Coord3d struct {
	X int
	Y int
	Z int
}

func (c Coord3d) Manhattan() int {
	return int(math.Abs(float64(c.X)) + math.Abs(float64(c.Y)) + math.Abs(float64(c.Z)))
}

func (c Coord3d) Add(cc Coord3d) Coord3d {
	return Coord3d{
		c.X + cc.X,
		c.Y + cc.Y,
		c.Z + cc.Z,
	}
}
func (c Coord3d) Sub(cc Coord3d) Coord3d {
	return Coord3d{
		c.X - cc.X,
		c.Y - cc.Y,
		c.Z - cc.Z,
	}
}

func (c Coord3d) TwoD() Coord {
	return Coord{
		c.X,
		c.Y,
	}
}

func ExtentsOf3DIntMap(in map[Coord3d]int) (min, max Coord3d) {
	min = Coord3d{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Coord3d{math.MinInt64, math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.Z < min.Z {
			min.Z = pt.Z
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
		if pt.Z > max.Z {
			max.Z = pt.Z
		}
	}
	return min, max
}
func ExtentsOf3DBoolMap(in map[Coord3d]bool) (min, max Coord3d) {
	min = Coord3d{math.MaxInt64, math.MaxInt64, math.MaxInt64}
	max = Coord3d{math.MinInt64, math.MinInt64, math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.Z < min.Z {
			min.Z = pt.Z
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
		if pt.Z > max.Z {
			max.Z = pt.Z
		}
	}
	return min, max
}
