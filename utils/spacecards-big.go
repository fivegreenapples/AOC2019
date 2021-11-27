package utils

import "fmt"

type bigPack struct {
	size      int
	hopLength int
	startCard int
}

func (pack *bigPack) DealIntoNewStack() {
	pack.startCard -= pack.hopLength
	pack.hopLength = -pack.hopLength
	pack.applyModulo()
}

func (pack *bigPack) DealWithIncrement(inc int) {
	// find inverse by solving
	// ((x * size)+1) % inc == 0
	// inverse is then ((x * size)+1) / inc
	inverse := -1
	for x := 0; x < inc; x++ {
		if ((x*pack.size)+1)%inc == 0 {
			inverse = ((x * pack.size) + 1) / inc
			break
		}
	}
	if inverse == -1 {
		panic("inverse not found")
	}
	pack.hopLength *= inverse
	pack.applyModulo()
}

func (pack *bigPack) Cut(at int) {
	pack.startCard += at * pack.hopLength
	pack.applyModulo()
}

func (pack *bigPack) CardAt(idx int) int {
	if idx < 0 {
		return 0
	}
	val := pack.startCard
	for i := 0; i < idx; i++ {
		val += pack.hopLength
	}
	if val < 0 {
		return ((val % pack.size) + pack.size) % pack.size
	}
	return val % pack.size
}

func (pack *bigPack) IndexOf(val int) int {
	i := 0
	testval := pack.startCard
	if testval < 0 {
		testval = ((testval % pack.size) + pack.size) % pack.size
	} else {
		testval = testval % pack.size
	}

	for testval != val {
		i++
		testval += pack.hopLength
		if testval < 0 {
			testval = ((testval % pack.size) + pack.size) % pack.size
		} else {
			testval = testval % pack.size
		}
	}

	return i
}

func (pack *bigPack) LimitedCSV(n int) string {
	if n <= 0 {
		return ""
	}
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = pack.CardAt(i)
	}
	return IntsToCSV(vals)
}

func (pack *bigPack) Reapply(times int) {
	fmt.Println("pack.startCard", pack.startCard, "pack.hopLength", pack.hopLength)

	119315717514047
	101741582076661
	// curent(9, w9)
	// one   (9 + 9*9, w9*9)
	// two   (9 + 9*9 + (9 + 9*9)*9*9)
	// three ()

	pack.startCard += times * pack.startCard
	pack.applyModulo()

	// fmt.Println(times)
	// currenthopLength := pack.hopLength
	// fmt.Println(currenthopLength)
	// n := 0
	// for {
	// 	n++
	// 	pack.hopLength *= pack.hopLength
	// 	pack.applyModulo()
	// 	// fmt.Println(n, pack.hopLength)
	// 	if pack.hopLength == currenthopLength {
	// 		break
	// 	}
	// }
	// remainder := times % n
	// fmt.Println(remainder)
	// for r := 0; r < remainder; r++ {
	// 	pack.hopLength *= pack.hopLength
	// 	pack.applyModulo()
	// }
}

func (pack *bigPack) applyModulo() {
	if pack.startCard < 0 {
		pack.startCard = (pack.startCard % pack.size) + pack.size
	} else {
		pack.startCard = pack.startCard % pack.size
	}
	if pack.hopLength < 0 {
		pack.hopLength = (pack.hopLength % pack.size) + pack.size
	} else {
		pack.hopLength = pack.hopLength % pack.size
	}
}
