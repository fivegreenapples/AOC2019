package utils

import (
	"strconv"
	"strings"
)

func Lines(in string) []string {
	lines := strings.Split(in, "\n")
	trimmed := []string{}
	for _, l := range lines {
		trimmed = append(trimmed, strings.TrimSpace(l))
	}
	return trimmed
}

func LinesAsInts(in string) []int {
	return StringsToInts(Lines(in))
}

func CsvToInts(in string) []int {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, ",")
	bits := strings.Split(in, ",")
	return StringsToInts(bits)
}
func CsvToStrings(in string) []string {
	in = strings.TrimSpace(in)
	in = strings.Trim(in, ",")
	bits := strings.Split(in, ",")
	return bits
}

func StringsToInts(inStrings []string) []int {
	ints := []int{}
	for _, in := range inStrings {
		in := strings.TrimSpace(in)
		thisInt, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}
		ints = append(ints, thisInt)
	}
	return ints
}

func StringToDigits(in string) []int {
	digits := make([]int, len(in))
	for i, d := range strings.Split(in, "") {
		digits[i] = MustAtoi(d)
	}
	return digits
}

func DigitsToString(in []int) string {
	stringDigits := make([]string, len(in))
	for i, d := range in {
		stringDigits[i] = strconv.Itoa(d)
	}
	return strings.Join(stringDigits, "")
}

func MustAtoi(in string) int {
	ret, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return ret
}

func StringSliceReverse(in []string) []string {
	for left, right := 0, len(in)-1; left < right; left, right = left+1, right-1 {
		in[left], in[right] = in[right], in[left]
	}
	return in
}
