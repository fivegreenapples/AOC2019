package days

import (
	"strconv"

	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day22Part1(in string) string {

	pack := utils.NewPack(10007)
	d22ProcessShuffle(pack, in)

	return strconv.Itoa(int(pack.IndexOf(2019)))
}

func (r *Runner) Day22Part2(in string) string {
	pack := utils.NewPack(119315717514047)
	d22ProcessShuffle(pack, in)
	pack.Reapply(101741582076661 - 1)
	return strconv.Itoa(int(pack.CardAt(2020)))
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
			pack.DealWithIncrement(int64(utils.MustAtoi(technique[5])))
		case "cut ":
			pack.Cut(int64(utils.MustAtoi(technique[5])))
		}
	}

}
