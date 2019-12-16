package days

import (
	"fmt"
	"testing"

	"github.com/fivegreenapples/AOC2019/utils"
)

func TestDay16FFT(t *testing.T) {

	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fft := fftList(input)
	fft.runPhase()
	if testing.Verbose() {
		fmt.Println("After 1 phase:", fft)
	}
	fft.runPhase()
	if testing.Verbose() {
		fmt.Println("After 2 phases:", fft)
	}
	fft.runPhase()
	if testing.Verbose() {
		fmt.Println("After 3 phases:", fft)
	}
	fft.runPhase()
	if testing.Verbose() {
		fmt.Println("After 4 phases:", fft)
	}

	result := utils.IntsToCSVSep([]int(fft), "")
	expected := "01029498"
	if result != expected {
		t.Errorf("day16 FFT test failed. Expected %s, got %s", expected, result)
	}
}

func TestDay16Part1(t *testing.T) {

	testInputs := map[string]string{
		"80871224585914546619083218645595": "24176176",
		"19617804207202209144916044189917": "73745418",
		"69317163492948606335995924319873": "52432133",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(16, 1, in)
		if out != expectedOut {
			t.Errorf("day16 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}

func TestDay16Part2(t *testing.T) {

	testInputs := map[string]string{
		"03036732577212944063491565474664": "84462026",
		"02935109699940807407585447034323": "78725270",
		"03081770884921959731165446850517": "53553731",
	}

	dayRunner := NewRunner(testing.Verbose())

	for in, expectedOut := range testInputs {
		out, _ := dayRunner.Run(16, 2, in)
		if out != expectedOut {
			t.Errorf("day16 failed with %s. Expected %s, got %s", in, expectedOut, out)
		}
	}

}
