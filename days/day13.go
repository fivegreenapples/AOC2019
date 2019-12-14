package days

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/fivegreenapples/AOC2019/intcode"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day13Part1(in string) string {
	game := intcode.New(utils.CsvToInts(in))
	_, out := game.RunSlice(nil)

	blockTiles := 0
	for i := 0; i < len(out); i += 3 {
		if out[i+2] == 2 {
			blockTiles++
		}
	}

	return strconv.Itoa(blockTiles)
}

func (r *Runner) Day13Part2(in string) string {
	gameProgram := utils.CsvToInts(in)
	gameProgram[0] = 2
	game := intcode.New(gameProgram)

	input := make(chan int)
	output := make(chan int)
	grid := map[utils.Coord]int{}
	score := 0

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		paddleInstr, paddleX, ballX := 0, 0, 0
		for {
			select {
			case a, valid := <-output:
				if !valid {
					if r.verbose {
						renderGrid(grid, score)
					}
					wg.Done()
					return
				}
				b := <-output
				c := <-output

				if a == -1 {
					score = c
				} else {
					if c == 3 {
						paddleX = a
					} else if c == 4 {
						ballX = a
					}
					grid[utils.Coord{a, b}] = c
				}

				if ballX == paddleX {
					paddleInstr = 0
				} else if ballX < paddleX {
					paddleInstr = -1
				} else {
					paddleInstr = 1
				}

			case input <- paddleInstr:
				if r.verbose {
					renderGrid(grid, score)
					time.Sleep(time.Millisecond * 40)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		game.Run(input, output)
		close(output)
		wg.Done()
	}()

	wg.Wait()
	return strconv.Itoa(score)
}

func renderGrid(grid map[utils.Coord]int, score int) {
	min, max := utils.ExtentsOfIntMap(grid)

	for y := min.Y; y <= max.Y; y++ {
		for x := min.X; x <= max.X; x++ {

			tile := grid[utils.Coord{x, y}]
			switch tile {
			case 0:
				fmt.Print(" ")
			case 1:
				fmt.Print("W")
			case 2:
				fmt.Print("@")
			case 3:
				fmt.Print("=")
			case 4:
				fmt.Print("O")
			default:
				fmt.Print(tile)
			}
		}
		fmt.Println()
	}
	fmt.Println("Score:", score)
}
