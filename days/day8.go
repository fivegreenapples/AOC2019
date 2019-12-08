package days

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

func (r *Runner) Day8Part1(in string) string {

	// image is 25 x 6
	minZeroes := math.MaxInt32
	var minZeroLayer map[rune]int
	currentLayerOfInterest := map[rune]int{}
	for i, char := range in {
		charCount := currentLayerOfInterest[char]
		currentLayerOfInterest[char] = charCount + 1

		if (i+1)%150 == 0 {
			if r.verbose {
				fmt.Println(currentLayerOfInterest)
			}
			lastZeroes := currentLayerOfInterest['0']
			if lastZeroes < minZeroes {
				minZeroes = lastZeroes
				minZeroLayer = currentLayerOfInterest
			}
			currentLayerOfInterest = map[rune]int{}
		}
	}
	if r.verbose {
		fmt.Println(minZeroLayer)
	}
	oneDigits := minZeroLayer['1']
	twoDigits := minZeroLayer['2']
	return strconv.Itoa(oneDigits * twoDigits)
}

func (r *Runner) Day8Part2(in string) string {
	// image is 25 x 6
	image := bytes.Repeat([]byte{'2'}, 150)
	for i, char := range []byte(in) {
		if char == '2' {
			continue
		}
		if image[i%150] != '2' {
			continue
		}
		if char == '0' {
			// map black to space for easier reading in a (black) terminal
			image[i%150] = ' '
		} else {
			image[i%150] = char
		}
	}

	output := ""
	for p := 0; p < 150; p += 25 {
		output += fmt.Sprintf("%s\n", image[p:p+25])
	}

	return output
}
