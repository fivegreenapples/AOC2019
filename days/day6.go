package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day6Part1(in string) string {
	orbitData := utils.Lines(in)

	orbits := map[string][]string{}
	for _, o := range orbitData {
		bodies := strings.Split(o, ")")

		center := bodies[0]
		orbital := bodies[1]

		if currentOrbitals, found := orbits[center]; found {
			currentOrbitals = append(currentOrbitals, orbital)
			orbits[center] = currentOrbitals
		} else {
			orbits[center] = []string{orbital}
		}
	}

	orbitCount := 0
	depth := 0
	currentCenters := []string{"COM"}
	for len(currentCenters) > 0 {

		newCenters := []string{}
		for _, current := range currentCenters {
			orbitCount += depth
			if orbitals, found := orbits[current]; found {
				newCenters = append(newCenters, orbitals...)
			}

		}
		currentCenters = newCenters
		depth++
	}

	return strconv.Itoa(orbitCount)
}

func (r *Runner) Day6Part2(in string) string {
	orbitData := utils.Lines(in)

	orbits := map[string]string{}
	for _, o := range orbitData {
		bodies := strings.Split(o, ")")

		center := bodies[0]
		orbital := bodies[1]

		orbits[orbital] = center
	}

	// create path from YOU to COM and SAN to COM
	youPath := []string{}
	body := "YOU"
	for body != "COM" {
		if nextBody, found := orbits[body]; found {
			youPath = append(youPath, nextBody)
			body = nextBody
		} else {
			panic(fmt.Errorf("can't find %s in orbit map", body))
		}
	}
	sanPath := []string{}
	body = "SAN"
	for body != "COM" {
		if nextBody, found := orbits[body]; found {
			sanPath = append(sanPath, nextBody)
			body = nextBody
		} else {
			panic(fmt.Errorf("can't find %s in orbit map", body))
		}
	}

	// Reverse the paths so we can walk up from COM to find the first diff
	youPath = utils.StringSliceReverse(youPath)
	sanPath = utils.StringSliceReverse(sanPath)

	if r.verbose {
		fmt.Println("YOU path:", youPath)
		fmt.Println("SAN path:", sanPath)
	}

	// find depth of common path from COM
	depth := 0
	for {
		if youPath[depth] != sanPath[depth] {
			break
		}
		depth++
	}

	// common paths up to depth
	// So minimum orbital transfer is remaining lengths in each path
	orbitalTransfers := len(youPath) - depth + len(sanPath) - depth

	return strconv.Itoa(orbitalTransfers)
}
