package days

import (
	"fmt"
	"strconv"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day20Part1(in string) string {

	maze, entrance, exit := d20ProcessMaze(in)
	if r.verbose {
		d20RenderMaze(maze, entrance, exit)
	}

	shortest := d20FindShortestRoute(maze, entrance, exit)
	return strconv.Itoa(shortest)
}

func (r *Runner) Day20Part2(in string) string {
	maze, entrance, exit := d20ProcessMaze(in)
	if r.verbose {
		d20RenderMaze(maze, entrance, exit)
	}

	shortest := d20FindShortestRouteRecursive(maze, entrance, exit)
	return strconv.Itoa(shortest)
}

func d20RenderMaze(maze map[utils.Coord]string, entrance, exit utils.Coord) {
	min, max := utils.ExtentsOfStringMap(maze)

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {
			thisCoord := utils.Coord{x, y}
			switch thisCoord {
			case entrance:
				fmt.Print("I")
			case exit:
				fmt.Print("O")
			default:
				fmt.Print(maze[utils.Coord{x, y}])
			}
		}
		fmt.Println()
	}

}
func d20ProcessMaze(in string) (maze map[utils.Coord]string, entrance, exit utils.Coord) {

	maze = map[utils.Coord]string{}
	letters := map[utils.Coord]bool{}
	var x, y int
	for _, char := range in {

		if char == '\n' {
			y++
			x = 0
			continue
		}

		maze[utils.Coord{x, y}] = string(char)

		if char >= 'A' && char <= 'Z' {
			letters[utils.Coord{x, y}] = true
		}

		x++
	}

	for coord := range letters {

		thisLetter := maze[coord]
		if len(thisLetter) == 0 || len(thisLetter) > 1 {
			continue // indicates it has already been processed
		}

		var portalCoord, otherCoord, openTileCoord utils.Coord
		var portal string

		n := maze[coord.Add(utils.Coord{0, -1})]
		s := maze[coord.Add(utils.Coord{0, 1})]
		ss := maze[coord.Add(utils.Coord{0, 2})]
		if s >= "A" && s <= "Z" {
			portal = thisLetter + s

			if n == "." {
				portalCoord = coord
				otherCoord = coord.Add(utils.Coord{0, 1})
				openTileCoord = coord.Add(utils.Coord{0, -1})
			} else if ss == "." {
				portalCoord = coord.Add(utils.Coord{0, 1})
				otherCoord = coord
				openTileCoord = coord.Add(utils.Coord{0, 2})
			} else {
				panic("unexplained maze feature")
			}
		}

		w := maze[coord.Add(utils.Coord{-1, 0})]
		e := maze[coord.Add(utils.Coord{1, 0})]
		ee := maze[coord.Add(utils.Coord{2, 0})]
		if e >= "A" && e <= "Z" {
			portal = thisLetter + e

			if w == "." {
				portalCoord = coord
				otherCoord = coord.Add(utils.Coord{1, 0})
				openTileCoord = coord.Add(utils.Coord{-1, 0})
			} else if ee == "." {
				portalCoord = coord.Add(utils.Coord{1, 0})
				otherCoord = coord
				openTileCoord = coord.Add(utils.Coord{2, 0})
			} else {
				panic("unexplained maze feature")
			}
		}

		maze[portalCoord] = portal
		maze[otherCoord] = " "

		if portal == "AA" {
			maze[portalCoord] = " "
			entrance = openTileCoord
		} else if portal == "ZZ" {
			maze[portalCoord] = " "
			exit = openTileCoord
		}

	}

	return maze, entrance, exit
}

func d20FindPortals(maze map[utils.Coord]string) map[string]map[utils.Coord]utils.Coord {

	portals := map[string]map[utils.Coord]utils.Coord{}

	for c, val := range maze {

		if len(val) == 2 {
			if val == "AA" || val == "ZZ" {
				// special case for entrance and exit which aren't portals
				portals[val] = map[utils.Coord]utils.Coord{
					c: utils.Origin,
				}
			} else {
				curMap := portals[val]
				if curMap == nil {
					// if no existing portal then insert a holding map with a known key for this coord
					curMap = map[utils.Coord]utils.Coord{
						utils.Origin: c,
					}
				} else {
					// existing map will only have one key, that being the known key used above
					// find what it references, delete and it and add entries for the portal
					other := curMap[utils.Origin]
					delete(curMap, utils.Origin)
					curMap[other] = c
					curMap[c] = other
				}
				portals[val] = curMap
			}
		}

	}

	return portals
}

func d20FindShortestRoute(maze map[utils.Coord]string, entrance, exit utils.Coord) int {
	portals := d20FindPortals(maze)

	tips := []utils.Coord{
		entrance,
	}

	seenPoints := map[utils.Coord]bool{
		entrance: true,
	}

	movementOptions := []utils.Coord{
		utils.Coord{0, -1},
		utils.Coord{1, 0},
		utils.Coord{0, 1},
		utils.Coord{-1, 0},
	}

	length := 0
	for len(tips) > 0 {
		newTips := []utils.Coord{}
		length++
		for _, c := range tips {

			for _, move := range movementOptions {
				next := c.Add(move)

				if next == exit {
					return length
				}

				if seenPoints[next] {
					continue
				}
				if maze[next] == "." {
					seenPoints[next] = true
					newTips = append(newTips, next)
				} else if len(maze[next]) == 2 {
					// change next to be the other side of the portal and then find nearest open tile
					next = portals[maze[next]][next]
					for _, move2 := range movementOptions {
						next2 := next.Add(move2)
						if seenPoints[next2] {
							continue
						}
						if maze[next2] == "." {
							seenPoints[next2] = true
							newTips = append(newTips, next2)
							break
						}
					}

				}
			}

		}
		tips = newTips
	}
	return -1
}

func d20FindShortestRouteRecursive(maze map[utils.Coord]string, entrance, exit utils.Coord) int {
	portals := d20FindPortals(maze)
	min, max := utils.ExtentsOfStringMap(maze)
	mazeMin, mazeMax := min.Add(utils.Coord{2, 2}), max.Sub(utils.Coord{2, 2}) // accounts for the border

	useEntrance := utils.Coord3d{
		X: entrance.X,
		Y: entrance.Y,
		Z: 0,
	}
	useExit := utils.Coord3d{
		X: exit.X,
		Y: exit.Y,
		Z: 0,
	}

	tips := []utils.Coord3d{
		useEntrance,
	}

	seenPoints := map[utils.Coord3d]bool{
		useEntrance: true,
	}

	var (
		north = utils.Coord3d{0, -1, 0}
		east  = utils.Coord3d{1, 0, 0}
		south = utils.Coord3d{0, 1, 0}
		west  = utils.Coord3d{-1, 0, 0}
	)
	movementOptions := []utils.Coord3d{north, east, south, west}

	length := 0
	for len(tips) > 0 {
		newTips := []utils.Coord3d{}
		length++
		for _, c := range tips {

			for _, move := range movementOptions {
				next := c.Add(move)

				if next == useExit {
					return length
				}

				if _, seen := seenPoints[next]; seen {
					continue
				}
				if maze[next.TwoD()] == "." {
					seenPoints[next] = true
					newTips = append(newTips, next)
				} else if len(maze[next.TwoD()]) == 2 {

					var thisDepth int
					if d20IsCoordInsideBounds(mazeMin, mazeMax, next.TwoD()) {
						thisDepth = next.Z + 1
					} else {
						if next.Z == 0 {
							// not allowed through outside portal
							continue
						}
						thisDepth = next.Z - 1
					}

					// change next to be the other side of the portal and then find nearest open tile
					next = portals[maze[next.TwoD()]][next.TwoD()].ThreeD(thisDepth)
					for _, move2 := range movementOptions {
						next2 := next.Add(move2)
						if _, seen2 := seenPoints[next2]; seen2 {
							continue
						}
						if maze[next2.TwoD()] == "." {
							seenPoints[next2] = true
							newTips = append(newTips, next2)
						}
					}

				}
			}

		}
		tips = newTips
	}
	return -1
}

func d20IsCoordInsideBounds(min, max, test utils.Coord) bool {

	return test.X > min.X && test.Y > min.Y && test.X < max.X && test.Y < max.Y

}
