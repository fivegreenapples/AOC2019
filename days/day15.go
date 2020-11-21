package days

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day15Part1(in string) string {
	_, minSteps, _ := r.d15ExploreWithDroid(in)
	return strconv.Itoa(minSteps)
}
func (r *Runner) Day15Part2(in string) string {
	curVerbose := r.verbose
	r.verbose = false
	mappedArea, _, o2Loc := r.d15ExploreWithDroid(in)
	r.verbose = curVerbose

	tips := []utils.Coord{o2Loc}
	minutes := 0

	for {
		if r.verbose {
			renderArea(mappedArea, utils.Origin)
			fmt.Printf("After %d minutes\n", minutes)
			time.Sleep(time.Millisecond * 50)
		}

		newTips := []utils.Coord{}
		for _, tip := range tips {

			for _, dir := range []direction{north, south, east, west} {
				nextTip := tip.Add(deltaForDirection(dir))
				if nextPos, exists := mappedArea[nextTip]; exists && nextPos.typ == floor {
					nextPos.typ = oxygen
					newTips = append(newTips, nextTip)
				}
			}

		}
		tips = newTips
		if len(tips) == 0 {
			break
		}

		minutes++
	}

	return strconv.Itoa(minutes)
}
func (r *Runner) d15ExploreWithDroid(in string) (mappedArea map[utils.Coord]*pos, minStepsToO2 int, o2Loc utils.Coord) {
	remoteProgram := utils.CsvToInts(in)
	droid := intcode.New(remoteProgram)

	input := make(chan int)
	output := make(chan int)

	go droid.Run(input, output)

	area := map[utils.Coord]*pos{
		utils.Origin: {
			typ: floor,
		},
	}
	currentLoc, prevLoc, oxygenLoc := utils.Origin, utils.Origin, utils.Origin
	breadcrumbs := []direction{}
	minStepsToOxygen := math.MaxInt64

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			if r.verbose && prevLoc != currentLoc {
				renderArea(area, currentLoc)
				time.Sleep(time.Millisecond * 50)
			}

			prevLoc = currentLoc

			nextDirection, nextDirMove := area[currentLoc].nextUnexploredDirection()
			if nextDirection == none {
				// retrace steps and loop
				if len(breadcrumbs) == 0 {
					// we have explored everywhere
					if r.verbose {
						renderArea(area, currentLoc)
					}
					wg.Done()
					return
				}

				// get the direction we last moved and reverse it (pop a direction off the breadcrumb stack)
				var lastMovedDirection direction
				lastMovedDirection, breadcrumbs = breadcrumbs[len(breadcrumbs)-1], breadcrumbs[:len(breadcrumbs)-1]
				goBackDirection := reverseDirection(lastMovedDirection)
				// move the current location based on this
				currentLoc = currentLoc.Add(deltaForDirection(goBackDirection))
				// tell the droid where we want to go...
				input <- int(goBackDirection)
				// ...and wait for response
				<-output

				// we don't care about the response so we just continue the loop
				continue
			}

			// send command
			input <- int(nextDirection)
			// wait for response
			status := <-output

			area[currentLoc].markDirAsExplored(nextDirection)
			newLoc := currentLoc.Add(nextDirMove)
			posAtLoc := area[newLoc]
			if posAtLoc == nil {
				posAtLoc = &pos{
					typ: unknown,
				}
				posAtLoc.setHaveComeFrom(nextDirection)
				area[newLoc] = posAtLoc
			}

			switch status {
			case 0: // hit a wall but droid didn't move
				posAtLoc.typ = wall
			case 1: // droid moved over floor
				posAtLoc.typ = floor
				breadcrumbs = append(breadcrumbs, nextDirection)
				currentLoc = newLoc
			case 2: // droid moved to oxygen tank
				posAtLoc.typ = oxygen
				breadcrumbs = append(breadcrumbs, nextDirection)
				currentLoc = newLoc
			}

			if posAtLoc.typ == oxygen {
				// Note, this algorithm lets the droid explore everywhere (partly
				// because it's required for part 2) and we use the min route to
				// oxygen as the smallest of the various routes that the droid found
				// to get there. But I would fully expect under some maps that this
				// would not find the right minimum. But it worked for my puzzle
				// input - in fact the droid only found one route to the oxygen.
				if len(breadcrumbs) < minStepsToOxygen {
					minStepsToOxygen = len(breadcrumbs)
				}
				oxygenLoc = currentLoc
			}

		}
	}()

	wg.Wait()
	return area, minStepsToOxygen, oxygenLoc
}

type tile string

const (
	wall    tile = "#"
	floor        = "."
	oxygen       = "o"
	unknown      = " "
)

type direction int

const (
	none  direction = 0
	north           = 1
	south           = 2
	east            = 4
	west            = 3
)

type pos struct {
	typ       tile
	beenNorth bool
	beenEast  bool
	beenSouth bool
	beenWest  bool
}

func (p pos) nextUnexploredDirection() (direction, utils.Coord) {
	if !p.beenNorth {
		return north, deltaForDirection(north)
	}
	if !p.beenEast {
		return east, deltaForDirection(east)
	}
	if !p.beenSouth {
		return south, deltaForDirection(south)
	}
	if !p.beenWest {
		return west, deltaForDirection(west)
	}
	return none, deltaForDirection(none)
}

func (p *pos) markDirAsExplored(d direction) {
	switch d {
	case north:
		p.beenNorth = true
	case east:
		p.beenEast = true
	case south:
		p.beenSouth = true
	case west:
		p.beenWest = true
	}
}

func (p *pos) setHaveComeFrom(d direction) {
	switch d {
	case north:
		p.beenSouth = true
	case east:
		p.beenWest = true
	case south:
		p.beenNorth = true
	case west:
		p.beenEast = true
	}
}

func renderArea(area map[utils.Coord]*pos, droidPos utils.Coord) {

	min, max := extentsOfMap(area)
	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			thisPos := utils.Coord{X: x, Y: y}
			if thisPos == utils.Origin {
				fmt.Print("X")
			} else if thisPos == droidPos {
				fmt.Print("D")
			} else if loc, found := area[thisPos]; found {
				fmt.Print(loc.typ)
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func extentsOfMap(in map[utils.Coord]*pos) (min, max utils.Coord) {
	min = utils.Coord{X: math.MaxInt64, Y: math.MaxInt64}
	max = utils.Coord{X: math.MinInt64, Y: math.MinInt64}
	for pt := range in {
		if pt.X < min.X {
			min.X = pt.X
		}
		if pt.Y < min.Y {
			min.Y = pt.Y
		}
		if pt.X > max.X {
			max.X = pt.X
		}
		if pt.Y > max.Y {
			max.Y = pt.Y
		}
	}
	return min, max
}

func reverseDirection(d direction) direction {
	switch d {
	case north:
		return south
	case east:
		return west
	case south:
		return north
	case west:
		return east
	}
	return none
}
func deltaForDirection(d direction) utils.Coord {
	switch d {
	case north:
		return utils.Coord{X: 0, Y: -1}
	case east:
		return utils.Coord{X: 1, Y: 0}
	case south:
		return utils.Coord{X: 0, Y: 1}
	case west:
		return utils.Coord{X: -1, Y: 0}
	}
	return utils.Origin
}
