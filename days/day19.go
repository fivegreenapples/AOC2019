package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/intcode"
)

func (r *Runner) Day19Part1(in string) string {
	drone := intcode.NewFromString(in)

	tractoredPoints := 0
	var beamForm strings.Builder
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			_, statuses := drone.RunSlice([]int{x, y})
			status := statuses[0]
			if status == 0 {
				fmt.Fprint(&beamForm, ".")
			} else if status == 1 {
				fmt.Fprint(&beamForm, "#")
				tractoredPoints++
			} else {
				fmt.Fprint(&beamForm, "?")
			}
		}
		fmt.Fprintln(&beamForm)
	}
	fmt.Fprintln(&beamForm)

	if r.verbose {
		fmt.Print(beamForm.String())
	}

	return strconv.Itoa(tractoredPoints)
}

func (r *Runner) Day19Part2(in string) string {
	drone := intcode.NewFromString(in)

	// find y for h bounds of length >=20 to get good estimate for gradient
	var y20, y20XMin, y20XMax int
	for {
		y20XMin, y20XMax = hBoundsForY(drone, y20)
		if y20XMax-y20XMin >= 20 {
			break
		}
		y20++
	}

	leftGradient := float64(y20) / float64(y20XMin)
	rightGradient := float64(y20) / float64(y20XMax)
	if r.verbose {
		fmt.Println("y20:", y20, "y20XMin:", y20XMin, "y20XMax:", y20XMax)
		fmt.Println("leftGradient:", leftGradient, "rightGradient:", rightGradient)
	}

	// Find estimate for row containing the top left pt of 100x100 square inside tractor beam
	// imagine 100x100 square having a top line and bottom line
	// top line is topY, and bottom line is (topY+99)
	// we need (topXMax - bottomXMin) == 99
	// topXMax = topY / rightGradient
	// bottomXMin = (topY+99) / leftGradient
	// So:
	// (topY / rightGradient) - ((topY+99) / leftGradient) == 99
	// topY*leftGradient - (topY+99)*rightGradient == 99 * leftGradient * rightGradient
	// topY*(leftGradient - rightGradient) - 99*rightGradient == 99 * leftGradient * rightGradient
	// topY == 99*(leftGradient*rightGradient + rightGradient) / (leftGradient - rightGradient)

	topY := 99.0 * (leftGradient*rightGradient + rightGradient) / (leftGradient - rightGradient)
	yEst := int(topY)

	// So get bounds at this estimate, and for 100th line further down for santa's ship
	yEstXMin, yEstXMax := hBoundsForY(drone, yEst)
	yEstPlus99XMin, yEstPlus99XMax := hBoundsForY(drone, yEst+99)
	yEstOverlap := yEstXMax - yEstPlus99XMin

	if r.verbose {
		fmt.Println("First estimate:")
		fmt.Println("yEst:", yEst)
		fmt.Println("yEstXMin:", yEstXMin, "yEstXMax:", yEstXMax)
		fmt.Println("yEstPlus99XMin:", yEstPlus99XMin, "yEstPlus99XMax:", yEstPlus99XMax)
		fmt.Println("yEstOverlap:", yEstOverlap)
		fmt.Println()
	}

	// check overlap and iterate
	for yEstOverlap != 99 {

		if r.verbose {
			fmt.Println("Iterating as overlap not 99")
		}

		if yEstOverlap < 99 {
			yEst++
		} else {
			yEst--
		}
		if r.verbose {
			fmt.Println("Trying", yEst)
		}

		yEstXMin, yEstXMax = hBoundsForY(drone, yEst)
		yEstPlus99XMin, yEstPlus99XMax = hBoundsForY(drone, yEst+99)
		yEstOverlap = yEstXMax - yEstPlus99XMin
	}

	// Finally work backwards until less than 99 and use last good
	for {
		tryY := yEst - 1
		if r.verbose {
			fmt.Println("Checking for lower y value. Trying", tryY)
		}

		_, tryXMax := hBoundsForY(drone, tryY)
		tryPlus99XMin, _ := hBoundsForY(drone, tryY+99)
		tryOverlap := tryXMax - tryPlus99XMin
		if tryOverlap < 99 {
			if r.verbose {
				fmt.Println("Last try too small. Correct y ==", yEst)
			}
			break
		}
		yEst = tryY
	}

	xEst, _ := hBoundsForY(drone, yEst+99)

	// So now we know that yEst is the topY proper and the coord of the top left is (xEst,yEst)
	return strconv.Itoa(10000*xEst + yEst)
}

func hBoundsForY(drone *intcode.VM, y int) (xMin, xMax int) {
	var x int
	xMin = -1
	for {
		_, statuses := drone.RunSlice([]int{x, y})
		status := statuses[0]

		if status == 1 && xMin < 0 {
			xMin = x
		}
		if status == 0 && xMin >= 0 {
			xMax = x - 1
			return
		}

		x++
		if y < 10 && x > 100 {
			// exception to account for sparse beams near emitter
			return 0, 0
		}
	}
}
