package main

import (
	"math"
)

type password int

func passwordFromDigits(in []int) password {
	p := 0
	l := len(in)
	for i, d := range in {
		p += d * int(math.Pow10(l-i-1))
	}
	return password(p)
}

func (p *password) makeValidMultipleSameDigits() {

	p.add(0)
	for !p.hasMultipleSameDigits() {
		p.add(1)
	}

}

func (p *password) makeValidDoubleDigit() {

	p.add(0)
	for !p.hasDoubleDigit() {
		p.add(1)
	}
}

func (p *password) add(n int) {

	*p = password(int(*p) + n)

	digits := p.toDigits()
	cp := 0
	bumped := false
	for i, d := range digits {
		if bumped {
			digits[i] = cp
		} else if d < cp {
			digits[i] = cp
			bumped = true
		}
		cp = digits[i]
	}

	*p = passwordFromDigits(digits)
}

func (p *password) hasMultipleSameDigits() bool {
	seenDigits := map[int]bool{}
	for _, d := range p.toDigits() {
		if _, seen := seenDigits[d]; seen {
			return true
		}
		seenDigits[d] = true
	}
	return false
}

func (p *password) hasDoubleDigit() bool {
	seenDigits := map[int]int{}
	for _, d := range p.toDigits() {
		cnt, seen := seenDigits[d]
		if !seen {
			seenDigits[d] = 1
		}
		seenDigits[d] = cnt + 1
	}
	for _, cnt := range seenDigits {
		if cnt == 2 {
			return true
		}
	}
	return false
}

func (p *password) toDigits() []int {

	pow := int(math.Log10(float64(*p)))
	digits := make([]int, pow+1)

	intP := int(*p)
	for p := pow; p >= 0; p-- {
		digits[p] = intP % 10
		intP = intP / 10
	}

	return digits
}
