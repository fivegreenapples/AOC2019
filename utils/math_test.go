package utils

import "testing"

func TestPrimeFactos(t *testing.T) {

	testInputs := map[int]string{
		1:  "1",
		2:  "2",
		3:  "3",
		4:  "2,2",
		5:  "5",
		6:  "2,3",
		7:  "7",
		8:  "2,2,2",
		9:  "3,3",
		10: "2,5",
		11: "11",
		12: "2,2,3",
		13: "13",
		14: "2,7",
		15: "3,5",
		16: "2,2,2,2",
		17: "17",
		18: "2,3,3",
	}

	for in, expectedOut := range testInputs {
		out := IntsToCSV(PrimeFactors(in))
		if out != expectedOut {
			t.Errorf("prime factors failed with %d. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestSimplifyCoord(t *testing.T) {

	testInputs := map[Coord]Coord{
		Coord{6, 9}:   Coord{2, 3},
		Coord{5, 0}:   Coord{1, 0},
		Coord{0, 5}:   Coord{0, 1},
		Coord{-5, 0}:  Coord{-1, 0},
		Coord{0, -5}:  Coord{0, -1},
		Coord{1, 1}:   Coord{1, 1},
		Coord{1, 8}:   Coord{1, 8},
		Coord{2, 3}:   Coord{2, 3},
		Coord{-2, 3}:  Coord{-2, 3},
		Coord{2, 10}:  Coord{1, 5},
		Coord{8, 28}:  Coord{2, 7},
		Coord{8, -28}: Coord{2, -7},
	}

	for in, expectedOut := range testInputs {
		out := in.Simplify()
		if out != expectedOut {
			t.Errorf("Simplify Coord failed with %v. Expected %v, got %v", in, expectedOut, out)
		}
	}

}
