package utils

import (
	"math/big"
)

type SpaceCards interface {
	DealIntoNewStack()
	DealWithIncrement(inc int64)
	Cut(at int64)
	CardAt(idx int64) int64
	IndexOf(val int64) int64
	LimitedCSV(n int64) string
	Reapply(times int)
}

func NewPack(size int64) SpaceCards {
	return &bigPack{
		size:      size,
		hopLength: 1,
	}
}

type bigPack struct {
	size      int64
	hopLength int64
	startCard int64
}

func (pack *bigPack) DealIntoNewStack() {
	pack.startCard -= pack.hopLength
	pack.hopLength = -pack.hopLength
	pack.applyModulo()
}

func (pack *bigPack) DealWithIncrement(inc int64) {
	// find inverse by solving
	// ((x * size)+1) % inc == 0
	// inverse is then ((x * size)+1) / inc
	var inverse int64 = -1
	var x int64
	for x = 0; x < inc; x++ {
		if ((x*pack.size)+1)%inc == 0 {
			inverse = ((x * pack.size) + 1) / inc
			break
		}
	}
	if inverse == -1 {
		panic("inverse not found")
	}

	bigInverse := big.NewInt(inverse)
	bigInverse.Mul(bigInverse, big.NewInt(pack.hopLength))
	bigInverse.Mod(bigInverse, big.NewInt(pack.size))

	pack.hopLength = bigInverse.Int64()
	pack.applyModulo()
}

func (pack *bigPack) Cut(at int64) {
	val := big.NewInt(at)
	val.Mul(val, big.NewInt(pack.hopLength))
	val.Add(val, big.NewInt(pack.startCard))
	val.Mod(val, big.NewInt(pack.size))
	pack.startCard = val.Int64()
	pack.applyModulo()
}

func (pack *bigPack) CardAt(idx int64) int64 {
	if idx < 0 {
		return 0
	}
	val := big.NewInt(idx)
	val.Mul(val, big.NewInt(pack.hopLength))
	val.Add(val, big.NewInt(pack.startCard))
	val.Mod(val, big.NewInt(pack.size))

	return val.Int64() % pack.size
}

func (pack *bigPack) IndexOf(val int64) int64 {
	var i int64 = 0
	var testval int64 = pack.startCard
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

func (pack *bigPack) LimitedCSV(n int64) string {
	if n <= 0 {
		return ""
	}
	vals := make([]int, n)
	var i int64
	for i = 0; i < n; i++ {
		vals[i] = int(pack.CardAt(i))
	}
	return IntsToCSV(vals)
}

func (pack *bigPack) Reapply(times int) {

	type packSettings struct {
		start int64
		hop   int64
	}

	remainingTimesToReapply := times
	settingsByTimesApplied := map[int]packSettings{}
	lastNumberOfTimesApplied := 1
	totalNumberOfTimesReapplied := 0

	settingsByTimesApplied[1] = packSettings{
		start: pack.startCard,
		hop:   pack.hopLength,
	}

	for remainingTimesToReapply > 0 {
		var thisSettingsToApply packSettings
		var thisTimesToApply int
		if lastNumberOfTimesApplied <= remainingTimesToReapply {
			thisTimesToApply = lastNumberOfTimesApplied
			thisSettingsToApply = settingsByTimesApplied[thisTimesToApply]
		} else {
			lastNumberOfTimesApplied = lastNumberOfTimesApplied / 2
			continue
		}

		bigStartDelta := big.NewInt(thisSettingsToApply.start)
		bigStartDelta.Mul(bigStartDelta, big.NewInt(pack.hopLength))
		bigStartDelta.Add(bigStartDelta, big.NewInt(pack.startCard))
		bigStartDelta.Mod(bigStartDelta, big.NewInt(pack.size))
		pack.startCard = bigStartDelta.Int64()

		bigHop := big.NewInt(thisSettingsToApply.hop)
		bigHop.Mul(bigHop, big.NewInt(pack.hopLength))
		bigHop.Mod(bigHop, big.NewInt(pack.size))
		pack.hopLength = bigHop.Int64()

		pack.applyModulo()

		remainingTimesToReapply -= thisTimesToApply
		totalNumberOfTimesReapplied += thisTimesToApply

		if _, found := settingsByTimesApplied[thisTimesToApply*2]; !found {
			settingsByTimesApplied[thisTimesToApply*2] = packSettings{
				start: pack.startCard,
				hop:   pack.hopLength,
			}

			lastNumberOfTimesApplied = thisTimesToApply * 2
		}

	}

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
