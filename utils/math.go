package utils

import "math"

import "fmt"

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func AbsInt(val int) int {
	return int(math.Abs(float64(val)))
}

func PrimeFactors(val int) []int {

	if val <= 0 {
		panic(fmt.Errorf("can't factorise %d as less than or equal to zero", val))
	}

	if val == 1 {
		return []int{1}
	}

	if val > primes[len(primes)-1] {
		panic(fmt.Errorf("can't factorise %d as greater than largest known prime (%d)", val, primes[len(primes)-1]))
	}

	primeFactors := []int{}
	pIdx := 0
	for {

		if val%primes[pIdx] != 0 {
			pIdx++
			continue
		}

		primeFactors = append(primeFactors, primes[pIdx])
		val = val / primes[pIdx]

		if val == 1 {
			break
		}
	}

	return primeFactors
}

func LargestCommonFactor(a, b int) int {

	pfA := PrimeFactors(a)
	pfB := PrimeFactors(b)

	aIdx := 0
	bIdx := 0

	cf := 1

	for aIdx < len(pfA) && bIdx < len(pfB) {
		if pfA[aIdx] == pfB[bIdx] {
			cf *= pfA[aIdx]
			aIdx++
			bIdx++
		} else if pfA[aIdx] < pfB[bIdx] {
			aIdx++
		} else {
			bIdx++
		}
	}

	return cf

}
