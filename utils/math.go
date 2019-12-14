package utils

import "math"

import "fmt"

var primes = []int{}

func makePrimes() {
	if len(primes) > 0 {
		return
	}
	var x, y, n int
	const N = 20000000
	nsqrt := math.Sqrt(N)

	is_prime := [N]bool{}

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= N && (n%12 == 1 || n%12 == 5) {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) + y*y
			if n <= N && n%12 == 7 {
				is_prime[n] = !is_prime[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= N && n%12 == 11 {
				is_prime[n] = !is_prime[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if is_prime[n] {
			for y = n * n; y < N; y += n * n {
				is_prime[y] = false
			}
		}
	}

	is_prime[2] = true
	is_prime[3] = true

	primes = make([]int, 0, 1270606)
	for x = 0; x < len(is_prime)-1; x++ {
		if is_prime[x] {
			primes = append(primes, x)
		}
	}

}

func AbsInt(val int) int {
	return int(math.Abs(float64(val)))
}

func PrimeFactors(val int) []int {
	makePrimes()
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
