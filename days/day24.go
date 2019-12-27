package days

import "fmt"

import "math/bits"

import "strconv"

func (r *Runner) Day24Part1(in string) string {
	// we represent a bug as a 1 bit in a uint64 where the position of the bit
	// is an index into the 5 by 5 grid

	var bioVal uint64

	idx := 0
	for _, char := range in {

		if char == '#' {
			bioVal += 1 << idx
			idx++
		} else if char == '.' {
			idx++
		}
	}

	seenValues := map[uint64]bool{
		bioVal: true,
	}

	for {
		var newVal uint64
		for i := 0; i < 25; i++ {

			nIdx := i - 5
			eIdx := i + 1
			sIdx := i + 5
			wIdx := i - 1

			if i%5 == 0 {
				wIdx = -1
			}
			if (i+1)%5 == 0 {
				eIdx = -1
			}

			var surroundVal uint64
			if nIdx >= 0 {
				surroundVal += (1 << nIdx) & bioVal
			}
			if eIdx >= 0 {
				surroundVal += (1 << eIdx) & bioVal
			}
			if sIdx >= 0 {
				surroundVal += (1 << sIdx) & bioVal
			}
			if wIdx >= 0 {
				surroundVal += (1 << wIdx) & bioVal
			}

			numSurroundingBugs := bits.OnesCount64(surroundVal)
			isCurrentBug := ((1 << i) & bioVal) > 0

			if isCurrentBug {
				isCurrentBug = numSurroundingBugs == 1
			} else {
				isCurrentBug = numSurroundingBugs == 1 || numSurroundingBugs == 2
			}

			if isCurrentBug {
				newVal += (1 << i)
			}

		}
		bioVal = newVal
		if seenValues[bioVal] {
			break
		}
		seenValues[bioVal] = true
		if r.verbose {
			fmt.Printf("%b\n", bioVal)
		}
	}

	return strconv.Itoa(int(bioVal))
}

func (r *Runner) Day24Part2(in string) string {
	// we represent a bug as a 1 bit in a uint64 where the position of the bit
	// is an index into the 5 by 5 grid

	var bioVal uint64
	idx := 0
	for _, char := range in {

		if char == '#' {
			bioVal += 1 << idx
			idx++
		} else if char == '.' {
			idx++
		}
	}

	minDepth, maxDepth := -1, 1
	valsByDepth := map[int]uint64{
		-1: 0,
		0:  bioVal,
		1:  0,
	}

	for minute := 1; minute <= 200; minute++ {
		newValsByDepth := map[int]uint64{}
		for depth, thisVal := range valsByDepth {

			outerVal := valsByDepth[depth-1]
			innerVal := valsByDepth[depth+1]

			var newVal uint64
			for i := 0; i <= 24; i++ {

				if i == 12 {
					// ignore grid center as this is the inner grid
					continue
				}

				var numSurroundingBugs int
				// north
				{
					var nVal uint64
					if i <= 4 {
						nVal = (1 << 7) & outerVal
					} else if i == 17 {
						nVal = (1 << 20) & innerVal
						nVal += (1 << 21) & innerVal
						nVal += (1 << 22) & innerVal
						nVal += (1 << 23) & innerVal
						nVal += (1 << 24) & innerVal
					} else {
						nVal = (1 << (i - 5)) & thisVal
					}
					numSurroundingBugs += bits.OnesCount64(nVal)
				}
				// east
				{
					var eVal uint64
					if (i+1)%5 == 0 {
						eVal = (1 << 13) & outerVal
					} else if i == 11 {
						eVal = (1 << 0) & innerVal
						eVal += (1 << 5) & innerVal
						eVal += (1 << 10) & innerVal
						eVal += (1 << 15) & innerVal
						eVal += (1 << 20) & innerVal
					} else {
						eVal = (1 << (i + 1)) & thisVal
					}
					numSurroundingBugs += bits.OnesCount64(eVal)
				}
				// south
				{
					var sVal uint64
					if i >= 20 {
						sVal = (1 << 17) & outerVal
					} else if i == 7 {
						sVal = (1 << 0) & innerVal
						sVal += (1 << 1) & innerVal
						sVal += (1 << 2) & innerVal
						sVal += (1 << 3) & innerVal
						sVal += (1 << 4) & innerVal
					} else {
						sVal = (1 << (i + 5)) & thisVal
					}
					numSurroundingBugs += bits.OnesCount64(sVal)
				}
				// west
				{
					var wVal uint64
					if i%5 == 0 {
						wVal = (1 << 11) & outerVal
					} else if i == 13 {
						wVal = (1 << 4) & innerVal
						wVal += (1 << 9) & innerVal
						wVal += (1 << 14) & innerVal
						wVal += (1 << 19) & innerVal
						wVal += (1 << 24) & innerVal
					} else {
						wVal = (1 << (i - 1)) & thisVal
					}
					numSurroundingBugs += bits.OnesCount64(wVal)
				}

				isCurrentBug := ((1 << i) & thisVal) > 0
				if isCurrentBug {
					isCurrentBug = numSurroundingBugs == 1
				} else {
					isCurrentBug = numSurroundingBugs == 1 || numSurroundingBugs == 2
				}

				if isCurrentBug {
					newVal += (1 << i)
				}
			}

			newValsByDepth[depth] = newVal
		}
		valsByDepth = newValsByDepth

		// Extend depth as necessary
		if valsByDepth[minDepth] > 0 {
			valsByDepth[minDepth-1] = 0
			minDepth--
		}
		if valsByDepth[maxDepth] > 0 {
			valsByDepth[maxDepth+1] = 0
			maxDepth++
		}

	}

	if r.verbose {
		for d := minDepth; d <= maxDepth; d++ {
			if valsByDepth[d] == 0 {
				continue
			}
			fmt.Printf("Depth %d:\n", d)
			d24RenderGrid(valsByDepth[d])
			fmt.Println()
		}
	}

	totalBugs := 0
	for _, val := range valsByDepth {
		totalBugs += bits.OnesCount64(val)
	}
	return strconv.Itoa(totalBugs)
}

func d24RenderGrid(val uint64) {
	for i := 0; i <= 24; i++ {
		if (1<<i)&val > 0 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		if (i+1)%5 == 0 {
			fmt.Println()
		}
	}
}
