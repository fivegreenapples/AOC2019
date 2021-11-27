package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day22Part1(in string) string {

	pack := utils.NewPack(10007)
	d22ProcessShuffle(pack, in)

	return strconv.Itoa(pack.IndexOf(2019))
}

func (r *Runner) Day22Part2(in string) string {
	pack := utils.NewPack(119315717514047)
	d22ProcessShuffle(pack, in)
	pack.Reapply(101741582076661)
	return strconv.Itoa(pack.CardAt(2020))
}

func d22ProcessShuffle(pack utils.SpaceCards, in string) {
	// deal into new stack
	// deal with increment 64
	// cut 8580
	shuffle := utils.StringsFromRegex(in, `^((deal into new stack)|(deal with increment )|(cut ))([-0-9]+)?$`)

	for _, technique := range shuffle {
		switch technique[1] {
		case "deal into new stack":
			pack.DealIntoNewStack()
		case "deal with increment ":
			pack.DealWithIncrement(utils.MustAtoi(technique[5]))
		case "cut ":
			pack.Cut(utils.MustAtoi(technique[5]))
		}
	}

}

func d22ProcessShuffleToGetPositionAt(positionAt int, packLength int, in string, repeats int) int {
	// deal into new stack
	// deal with increment 64
	// cut 8580
	lines := utils.Lines(in)
	// reverse shuffle instructions as we want to apply them in reverse
	lines = utils.StringSliceReverse(lines)
	reversedShuffle := strings.Join(lines, "\n")

	shuffle := utils.StringsFromRegex(reversedShuffle, `^((deal into new stack)|(deal with increment )|(cut ))([-0-9]+)?$`)

	// fmt.Println("")
	// fmt.Println("STARTING")
	currentPosition := positionAt
	seenPos := map[int]int{}
	iter := 0
	for r := 0; r < repeats; r++ {
		for instrIdx, technique := range shuffle {
			// oldPosition := currentPosition
			fmt.Println("current pos", currentPosition)
			// fmt.Println("processing", technique[1], technique[5])
			// if iter == 729990 ||
			// 	iter == 1844616 ||
			// 	iter == 8058119 ||
			// 	iter == 12116191 ||
			// 	iter == 5000431 ||
			// 	iter == 21297233 ||
			// 	iter == 22661467 ||
			// 	iter == 23317160 ||
			// 	iter == 21372876 ||
			// 	iter == 31554680 {
			// 	fmt.Println(iter, technique[0])
			// }
			if idx, seen := seenPos[currentPosition]; seen && idx == instrIdx {
				// fmt.Printf("Seen %d before at %d iterations. (now %d iterations)\n", currentPosition, when, iter)
				fmt.Printf("Seen %d before. (now %d iterations) (at idx %d with instr %s\n", currentPosition, iter, instrIdx, technique[0])
			} else {
				seenPos[currentPosition] = instrIdx
			}
			iter++
			switch technique[1] {
			case "deal into new stack":
				currentPosition = packLength - 1 - currentPosition
				// testOldPosition := packLength - 1 - currentPosition
				// if testOldPosition != oldPosition {
				// 	panic(fmt.Errorf("mismatch at deal into new stack. old %d, new %d,  test old %d ", oldPosition, currentPosition, testOldPosition))
				// }
			case "deal with increment ":
				var inc int
				if currentPosition == 0 {
					continue
				} else {
					// 0 1 2 3 4 5 6 7 8 9  inc 3 becomes
					// 0 7 4 1 8 5 2 9 6 3
					// OR 0 3 6 9 2 5 8 1 4 7 inc 3 becomes
					//    0 1 2 3 4 5 6 7 8 9
					//
					// 0 1 2 3 4 5 6 7 8 9  inc 7 becomes
					// 0 3 6 9 2 5 8 1 4 7
					// OR 0 7 4 1 8 5 2 9 6 3 inc 7 becomes
					//    0 1 2 3 4 5 6 7 8 9

					// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
					// 0     1     2     3     4     5     6     7     8     9
					// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9
					// 0             1             2             3             4             5             6             7             8             9
					// (pos * inc) % packLength = newpos
					// newpos + (N*packLength) / inc == pos
					inc = utils.MustAtoi(technique[5])
					mod := currentPosition % inc
					if mod == 0 {
						currentPosition = currentPosition / inc
					} else {
						packRemainder := packLength % inc
						packNumIncrements := packLength / inc
						test := (packNumIncrements + 1) * inc
						cycle := 0
						extras := 0
						for {
							cycle++
							thisMod := test % packLength
							if thisMod == mod {
								break
							}

							test += (packNumIncrements * inc)
							if thisMod < packRemainder {
								test += inc
								extras++
							}
						}

						// now cycle gives the number of loops round for this position
						currentPosition = (1 + currentPosition/inc) + (cycle * packNumIncrements) + extras

					}
				}
				// testOldPosition := (currentPosition * inc) % packLength
				// if testOldPosition != oldPosition {
				// 	panic(fmt.Errorf("mismatch at deal with increment:%d. old %d, new %d,  test old %d ", inc, oldPosition, currentPosition, testOldPosition))
				// }
			case "cut ":
				// cut 3
				// 0 1 2 3 4 5 6 7 8 9 becomes
				// 3 4 5 6 7 8 9 0 1 2

				// 7 8 9 0 1 2 3 4 5 6 becomes
				// 0 1 2 3 4 5 6 7 8 9
				cut := utils.MustAtoi(technique[5])
				if cut < 0 {
					cut = packLength + cut
				}

				currentPosition = (currentPosition + cut) % packLength

				// testOldPosition := (currentPosition + packLength - cut) % packLength
				// if testOldPosition != oldPosition {
				// 	panic(fmt.Errorf("mismatch at cut:%d. old %d, new %d, test old %d ", cut, oldPosition, currentPosition, testOldPosition))
				// }
			}
		}
	}

	// fmt.Println("returning", currentPosition)
	// fmt.Println()
	return currentPosition
}
