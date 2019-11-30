package prime

import (
	"github.com/jackytck/gowboy/common"
)

// Factors gives the prime factors of an int.
func Factors(n int) map[int]int {
	factors := make(map[int]int)
	d := 2
	for n > 1 {
		for n%d == 0 {
			n /= d
			factors[d]++
		}
		d++
	}
	return factors
}

// IsPrime determines if a given number is prime.
// https://en.wikipedia.org/wiki/Primality_test
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	b := common.SqrtInt(n)
	for i := 5; i <= b; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// Sieve gives all the primes below n via the method of Eratosthenes.
func Sieve(n int) []int {
	var primes []int
	comp := make([]bool, n+2)
	for i := 2; i <= common.SqrtInt(n); i++ {
		for j := i * i; j < n; j += i {
			comp[j] = true
		}
	}
	for i, v := range comp {
		if i > 1 && !v && i < n {
			primes = append(primes, i)
		}
	}
	return primes
}

// Range gives the primes in range [a, b).
func Range(a, b int) []int {
	var primes []int
	if b <= a {
		return primes
	}
	ps := Sieve(common.SqrtInt(b) + 1)
	size := b - a + 1
	comp := make([]bool, size)
	for _, p := range ps {
		for j := a/p*p - a; j < size; j += p {
			if j < 0 {
				continue
			}
			comp[j] = true
		}
	}
	for i, v := range comp {
		if !v && i < size {
			primes = append(primes, a+i)
		}
	}
	return primes
}
