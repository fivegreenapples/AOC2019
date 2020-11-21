package days

import (
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day10Part1(in string) string {

	asteroidLocs, xExtent, yExtent := getAsteroidLocations(in)
	maxVisibleAsteroids, maxVisibleCoord, maxVisibleAsteroidsLocs := findBestLocation(asteroidLocs, xExtent, yExtent)

	if r.verbose {
		fmt.Printf("Coord %d,%d with %d other asteroids detected\n", maxVisibleCoord.X, maxVisibleCoord.Y, maxVisibleAsteroids)
		for y := 0; y <= yExtent; y++ {
			for x := 0; x <= xExtent; x++ {

				pt := utils.Coord{X: x, Y: y}

				if pt == maxVisibleCoord {
					fmt.Print("O")
				} else if maxVisibleAsteroidsLocs[pt] {
					fmt.Print("*")
				} else if !asteroidLocs[pt] {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Print("\n")
		}
	}

	return strconv.Itoa(maxVisibleAsteroids)
}

func (r *Runner) Day10Part2(in string) string {
	// use part 1 to get the monitoring station
	asteroidLocs, xExtent, yExtent := getAsteroidLocations(in)
	maxVisibleAsteroids, monitoringCoord, maxVisibleAsteroidsLocs := findBestLocation(asteroidLocs, xExtent, yExtent)

	// convert all asteroid locations to angle and distance from monitoring station
	// group into buckets by angle
	angles := []int{}
	buckets := map[int][]utils.Coord{}
	for pt, isAsteroid := range asteroidLocs {
		if !isAsteroid {
			continue
		}
		if pt == monitoringCoord {
			continue
		}
		vec := pt.Sub(monitoringCoord)

		// Use some trickery to get the angle going clockwise from north as 0 to 2pi
		angle := (math.Pi + math.Atan2(float64(-vec.X), float64(vec.Y)))
		// round to hundredths of a degree, so we can have defnitive buckets without worries about floats
		angleMetric := int(math.Round(angle * 18000 / math.Pi))
		// Make sure due north has a zero angle
		angleMetric = angleMetric % 36000

		// Store pt in bucket
		if currentPts, exists := buckets[angleMetric]; !exists {
			angles = append(angles, angleMetric)
			buckets[angleMetric] = []utils.Coord{pt}
		} else {
			currentPts = append(currentPts, pt)
			buckets[angleMetric] = currentPts
		}
	}

	// loop over the buckets and sort asteroids by distance away from
	// monitoring station farthest first
	for angl, asteroids := range buckets {
		sort.Slice(asteroids, func(i, j int) bool {
			distI := asteroids[i].Sub(monitoringCoord).Manhattan()
			distJ := asteroids[j].Sub(monitoringCoord).Manhattan()
			return distJ < distI
		})
		buckets[angl] = asteroids
	}
	// Sort angles
	sort.Ints(angles)

	// now range over angles, lasering each asteroid until the 200th
	asteroidNum := 0
	laseredOrder := map[utils.Coord]int{}
	twoHundredth := utils.Coord{}
	for asteroidNum < 200 {
		for a := 0; a < len(angles); a++ {
			angle := angles[a]
			availAsteroids := buckets[angle]
			if len(availAsteroids) == 0 {
				continue
			}
			asteroidNum++
			laseredOrder[availAsteroids[len(availAsteroids)-1]] = asteroidNum
			if asteroidNum == 200 {
				twoHundredth = availAsteroids[len(availAsteroids)-1]
			}

			// put back in the bucket all but one of the asteroids (that one gets fried)
			buckets[angle] = availAsteroids[:len(availAsteroids)-1]
		}
	}

	if r.verbose {
		fmt.Printf("Coord %d,%d with %d other asteroids detected\n", monitoringCoord.X, monitoringCoord.Y, maxVisibleAsteroids)
		for y := 0; y <= yExtent; y++ {
			for x := 0; x <= xExtent; x++ {

				pt := utils.Coord{X: x, Y: y}
				order := laseredOrder[pt]

				if pt == monitoringCoord {
					fmt.Print("O")
				} else if order > 0 && order < 10 {
					fmt.Print(order % 10)
				} else if maxVisibleAsteroidsLocs[pt] {
					fmt.Print("*")
				} else if !asteroidLocs[pt] {
					fmt.Print(".")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Print("\n")
		}
	}
	return fmt.Sprintf("%d", twoHundredth.X*100+twoHundredth.Y)
}

func getAsteroidLocations(in string) (map[utils.Coord]bool, int, int) {
	mapRows := utils.Lines(in)
	asteroidLocs := map[utils.Coord]bool{}
	maxExtent := utils.Coord{X: 0, Y: 0}
	for y, row := range mapRows {
		for x, pt := range row {
			asteroidLocs[utils.Coord{X: x, Y: y}] = pt == '#'
			maxExtent = utils.Coord{X: x, Y: y}
		}
	}
	return asteroidLocs, maxExtent.X, maxExtent.Y
}

func findBestLocation(asteroidLocs map[utils.Coord]bool, xExtent, yExtent int) (int, utils.Coord, map[utils.Coord]bool) {
	maxVisibleAsteroids := 0
	maxVisibleCoord := utils.Coord{X: 0, Y: 0}
	maxVisibleAsteroidsLocs := map[utils.Coord]bool{}
	for y := 0; y <= yExtent; y++ {
		for x := 0; x <= xExtent; x++ {

			pt := utils.Coord{X: x, Y: y}

			if !asteroidLocs[pt] {
				continue
			}

			visibleAsteroids := calcVisibleAsteroids(pt, asteroidLocs)
			if len(visibleAsteroids) > maxVisibleAsteroids {
				maxVisibleAsteroids = len(visibleAsteroids)
				maxVisibleCoord = pt
				maxVisibleAsteroidsLocs = visibleAsteroids
			}

		}
	}

	return maxVisibleAsteroids, maxVisibleCoord, maxVisibleAsteroidsLocs
}

func calcVisibleAsteroids(from utils.Coord, asteroidLocs map[utils.Coord]bool) map[utils.Coord]bool {

	processedLocs := map[utils.Coord]bool{}
	visibleLocs := map[utils.Coord]bool{}

	for pt, isAsteroid := range asteroidLocs {

		if !isAsteroid {
			continue
		}

		if pt == from {
			continue
		}

		if processedLocs[pt] {
			continue
		}

		vec := pt.Sub(from)
		vec = vec.Simplify()

		current := from
		foundFirstAsteroid := false
		for {
			current = current.Add(vec)
			foundAsteroid, onMap := asteroidLocs[current]
			if !onMap {
				break
			}
			if foundAsteroid && !foundFirstAsteroid {
				foundFirstAsteroid = true
				visibleLocs[current] = true
			}
			processedLocs[current] = true
		}

	}
	return visibleLocs
}
