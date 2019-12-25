package days

import "fmt"

import "math/bits"

import "strconv"

func (r *Runner) Day24Part1(in string) string {

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
	// lines := utils.Lines(int)
	return ""
}
