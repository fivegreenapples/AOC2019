package days

import "testing"

func TestDay6Part1(t *testing.T) {

	testInputs := map[string]string{
		`COM)B
		 B)C
		 C)D
		 D)E
		 E)F
		 B)G
		 G)H
		 D)I
		 E)J
		 J)K
		 K)L`: "42",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 1, in)
		if out != expectedOut {
			t.Errorf("day6 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay6Part2(t *testing.T) {

	testInputs := map[string]string{
		`COM)B
		 B)C
		 C)D
		 D)E
		 E)F
		 B)G
		 G)H
		 D)I
		 E)J
		 J)K
		 K)L
		 K)YOU
		 I)SAN`: "4",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(6, 2, in)
		if out != expectedOut {
			t.Errorf("day6 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
