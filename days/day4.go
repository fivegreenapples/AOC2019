package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2019/password"
	"github.com/fivegreenapples/AOC2019/utils"
)

func (r *Runner) Day4Part1(in string) string {
	passRange := utils.StringsToInts(strings.Split(in, "-"))
	min := password.Password(passRange[0])
	max := password.Password(passRange[1])

	min.MakeValidMultipleSameDigits()

	numPasswords := 0
	for p := min; p < max; {
		numPasswords++
		if r.verbose {
			fmt.Println(numPasswords, p)
		}
		p.Add(1)
		p.MakeValidMultipleSameDigits()
	}

	return strconv.Itoa(numPasswords)
}

func (r *Runner) Day4Part2(in string) string {
	passRange := utils.StringsToInts(strings.Split(in, "-"))
	min := password.Password(passRange[0])
	max := password.Password(passRange[1])

	min.MakeValidDoubleDigit()

	numPasswords := 0
	for p := min; p < max; {
		numPasswords++
		if r.verbose {
			fmt.Println(numPasswords, p)
		}
		p.Add(1)
		p.MakeValidDoubleDigit()
	}

	return strconv.Itoa(numPasswords)
}
