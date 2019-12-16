package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day16Part1(in string) string {

	lists := utils.MultilineCsvToInts(in, "")

	list := fftList(lists[0])
	for i := 0; i < 100; i++ {
		list.runPhase()
	}
	return utils.IntsToCSVSep([]int(list[:8]), "")
}

func (r *Runner) Day16Part2(in string) string {
	// repeat input 10,000 times as a string
	longIn := strings.Repeat(in, 10000)
	lists := utils.MultilineCsvToInts(longIn, "")
	list := fftList(lists[0])
	offset, _ := strconv.Atoi(in[:7])

	// under a particular circumstance we can apply a custom short circuit
	if 2*offset >= len(list) {
		// This means that the input list will be covered just by the expanded zeroes and ones of
		// the FFT pattern. It means we can inplement a simplified algorithm that involves a running
		// total starting from the last element and working back. i.e. any element from the offset
		// is the sum of all future elements less the one before it.

		for ph := 1; ph <= 100; ph++ {
			runningTotal := 0
			for i := len(list) - 1; i >= offset; i-- {
				runningTotal += list[i]
				runningTotal = runningTotal % 10
				list[i] = runningTotal
			}
		}

	} else {
		for ph := 1; ph <= 100; ph++ {
			list.runPhaseForOffset(offset)
			if r.verbose {
				fmt.Println("Done phase:", ph)
			}
		}
	}

	message := utils.IntsToCSVSep(list[offset:offset+8], "")

	return message
}

type fftList []int

func (l *fftList) runPhase() {

	pattern := []int{0, 1, 0, -1}

	curValues := make([]int, len(*l))
	copy(curValues, *l)

	for lIdx := range *l {
		newVal := 0

		patternRepeats := lIdx + 1

		patIdx := 0
		listIdx := 0
		thisPatternRepeats := patternRepeats - 1
		for listIdx < len(curValues) {

			patValue := pattern[patIdx]

			if patValue == 0 {
				// skip values
				listIdx += thisPatternRepeats
				thisPatternRepeats = patternRepeats
				patIdx = (patIdx + 1) % len(pattern)
			} else {
				if thisPatternRepeats > 0 {
					newVal += patValue * curValues[listIdx]
					listIdx++
					thisPatternRepeats--
				} else {
					patIdx = (patIdx + 1) % len(pattern)
					thisPatternRepeats = patternRepeats
				}
			}
		}
		(*l)[lIdx] = utils.AbsInt(newVal % 10)
	}

}

func (l *fftList) runPhaseForOffset(offset int) {

	pattern := []int{0, 1, 0, -1}

	curValues := make([]int, len(*l))
	copy(curValues, *l)

	for lIdx := offset; lIdx < len(*l); lIdx++ {
		newVal := 0

		patternRepeats := lIdx + 1

		patIdx := 0
		listIdx := 0
		thisPatternRepeats := patternRepeats - 1
		// fmt.Println("pattern repeats is", thisPatternRepeats)
		for listIdx < len(curValues) {

			patValue := pattern[patIdx]

			if patValue == 0 {
				// skip values
				listIdx += thisPatternRepeats
				thisPatternRepeats = patternRepeats
				patIdx = (patIdx + 1) % len(pattern)
			} else {
				if thisPatternRepeats > 0 {
					newVal += patValue * curValues[listIdx]
					listIdx++
					thisPatternRepeats--
				} else {
					patIdx = (patIdx + 1) % len(pattern)
					thisPatternRepeats = patternRepeats
				}
			}
		}
		(*l)[lIdx] = utils.AbsInt(newVal % 10)
	}

}
