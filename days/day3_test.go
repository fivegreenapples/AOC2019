package days

import "testing"

func TestDay3Part1(t *testing.T) {

	testInputs := map[string]string{
		"R8,U5,L5,D3\nU7,R6,D4,L4": "6",
		"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83":               "159",
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7": "135",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 1, in)
		if out != expectedOut {
			t.Errorf("day3 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay3Part2(t *testing.T) {

	testInputs := map[string]string{
		"R8,U5,L5,D3\nU7,R6,D4,L4": "30",
		"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83":               "610",
		"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7": "410",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(3, 2, in)
		if out != expectedOut {
			t.Errorf("day3 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
