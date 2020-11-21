package days

import (
	"fmt"
	"math"
	"strconv"
	"sync"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

type ampInput struct {
	value  int
	phases map[int]bool
}

func (r *Runner) Day7Part1(in string) string {
	controlSoftware := utils.CsvToInts(in)
	controlVM := intcode.New(controlSoftware)
	controlVM.SetDebug(r.verbose)

	phasePermutations := permutations([]int{0, 1, 2, 3, 4})
	maxOutput := math.MinInt64
	for _, phases := range phasePermutations {

		currentInput := 0

		for _, phaseSetting := range phases {
			vmIn := []int{phaseSetting, currentInput}

			_, out := controlVM.RunSlice(vmIn)

			currentInput = out[0]
		}

		if currentInput > maxOutput {
			maxOutput = currentInput
		}

	}
	return strconv.Itoa(maxOutput)
}

func (r *Runner) Day7Part2(in string) string {

	controlSoftware := utils.CsvToInts(in)
	controlVM := intcode.New(controlSoftware)
	controlVM.SetDebug(r.verbose)

	phasePermutations := permutations([]int{5, 6, 7, 8, 9})
	var maxThrusterInput int
	for _, phases := range phasePermutations {

		if r.verbose {
			fmt.Println("Phaes:", phases)
		}

		eOutputAInput := make(chan int, 1)
		aOutputBInput := make(chan int, 1)
		bOutputCInput := make(chan int, 1)
		cOutputDInput := make(chan int, 1)
		dOutputEInput := make(chan int, 1)

		wait := sync.WaitGroup{}
		wait.Add(5)

		go func() {
			controlVM.Run(eOutputAInput, aOutputBInput)
			wait.Done()
		}()
		go func() {
			controlVM.Run(aOutputBInput, bOutputCInput)
			wait.Done()
		}()
		go func() {
			controlVM.Run(bOutputCInput, cOutputDInput)
			wait.Done()
		}()
		go func() {
			controlVM.Run(cOutputDInput, dOutputEInput)
			wait.Done()
		}()
		go func() {
			controlVM.Run(dOutputEInput, eOutputAInput)
			wait.Done()
		}()

		// Send phase settings to each core
		eOutputAInput <- phases[0]
		aOutputBInput <- phases[1]
		bOutputCInput <- phases[2]
		cOutputDInput <- phases[3]
		dOutputEInput <- phases[4]

		// Send initial input to core A
		eOutputAInput <- 0
		wait.Wait()

		finalOutput := <-eOutputAInput

		if finalOutput > maxThrusterInput {
			maxThrusterInput = finalOutput
		}
	}

	// fmt.Println(permutations([]int{1, 2, 3, 4, 5}))
	return strconv.Itoa(maxThrusterInput)
}

func permutations(options []int) [][]int {
	if len(options) == 0 {
		panic("can't do permutations of zero elements")
	}
	if len(options) == 1 {
		return [][]int{
			{
				options[0],
			},
		}
	}

	ret := [][]int{}
	for _, opt := range options {

		// copy original options to new slice filtering out current option
		subOptions := []int{}
		for _, subOpt := range options {
			if subOpt != opt {
				subOptions = append(subOptions, subOpt)
			}
		}

		subPerms := permutations(subOptions)
		for _, perm := range subPerms {
			thisPermutation := append([]int{opt}, perm...)
			ret = append(ret, thisPermutation)
		}
	}

	return ret
}
