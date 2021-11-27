package utils

type SpaceCards interface {
	DealIntoNewStack()
	DealWithIncrement(inc int)
	Cut(at int)
	CardAt(idx int) int
	IndexOf(val int) int
	LimitedCSV(n int) string
	Reapply(times int)
}

type simplePack struct {
	cards []int
}

func NewPack(size int) SpaceCards {

	return NewBigPack(size)

}

func NewBigPack(size int) SpaceCards {

	return &bigPack{
		size:      size,
		hopLength: 1,
	}

}

func NewSimplePack(size int) SpaceCards {

	pack := &simplePack{
		cards: make([]int, size),
	}

	for i := 0; i < size; i++ {
		pack.cards[i] = i
	}
	return pack

}

// deal into new stack
// deal with increment 64
// cut 8580

func (pack *simplePack) DealIntoNewStack() {
	// reverse operation
	for left, right := 0, len(pack.cards)-1; left < right; left, right = left+1, right-1 {
		pack.cards[left], pack.cards[right] = pack.cards[right], pack.cards[left]
	}
}

func (pack *simplePack) DealWithIncrement(inc int) {
	length := len(pack.cards)
	newpack := make([]int, length)

	j := 0
	for i := 0; i < length; i++ {
		newpack[j] = pack.cards[i]
		j = (j + inc) % length
	}

	pack.cards = newpack
}

func (pack *simplePack) Cut(at int) {
	length := len(pack.cards)
	newpack := make([]int, length)

	if at < 0 {
		at = length + at
	}

	for i := 0; i < length; i++ {
		newpack[i] = pack.cards[(at+i)%length]
	}

	pack.cards = newpack
}

func (pack *simplePack) CardAt(idx int) int {
	return pack.cards[idx]
}

func (pack *simplePack) IndexOf(val int) int {
	for idx, num := range pack.cards {
		if val == num {
			return idx
		}
	}
	return -1
}

func (pack *simplePack) LimitedCSV(n int) string {
	if n <= 0 {
		return ""
	}
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		vals[i] = pack.CardAt(i)
	}
	return IntsToCSV(vals)
}

func (pack *simplePack) Reapply(times int) {
	panic("not implemented")
}
