package days

import "testing"

func TestDay12TenSteps(t *testing.T) {

	testInputs := map[string]int{
		`<x=-1, y=0, z=2>
		<x=2, y=-10, z=-7>
		<x=4, y=-8, z=8>
		<x=3, y=5, z=-1>`: 179,
	}

	for in, expectedOut := range testInputs {
		out := calcEnergyAfterSteps(in, 10)
		if out != expectedOut {
			t.Errorf("day12 failed with 10 step calculation. Expected %d, got %d", expectedOut, out)
		}
	}

}
func TestDay12HundredSteps(t *testing.T) {

	testInputs := map[string]int{
		`<x=-8, y=-10, z=0>
		<x=5, y=5, z=10>
		<x=2, y=-7, z=3>
		<x=9, y=-8, z=-3>`: 1940,
	}

	for in, expectedOut := range testInputs {
		out := calcEnergyAfterSteps(in, 100)
		if out != expectedOut {
			t.Errorf("day12 failed with 100 step calculation. Expected %d, got %d", expectedOut, out)
		}
	}

}

func TestDay12Part2(t *testing.T) {

	testInputs := map[string]string{
		`<x=-1, y=0, z=2>
		<x=2, y=-10, z=-7>
		<x=4, y=-8, z=8>
		<x=3, y=5, z=-1>`: "1 2 3",
		`<x=-8, y=-10, z=0>
		<x=5, y=5, z=10>
		<x=2, y=-7, z=3>
		<x=9, y=-8, z=-3>`: "1 2 3",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(12, 2, in)
		if out != expectedOut {
			t.Errorf("day12 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
