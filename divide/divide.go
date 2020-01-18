package divide

import (
	"github.com/jackytck/gowboy/common"
	"github.com/jackytck/gowboy/prime"
)

// Divisors returns all the divisors of n.
func Divisors(n int) []int {
	if n == 1 {
		return []int{1}
	}
	d := ProperDivisors(n)
	d = append(d, n)
	return d
}

// ProperDivisors returns all the proper divisors of n.
func ProperDivisors(n int) []int {
	d := []int{1}
	var others []int
	for i := 2; i <= common.SqrtInt(n); i++ {
		if n%i == 0 {
			d = append(d, i)
			if other := n / i; other != i {
				others = append([]int{other}, others...)
			}
		}
	}
	d = append(d, others...)
	return d
}

// SumDivisors sums all the divisors of n.
func SumDivisors(n int) int {
	return common.Sum(Divisors(n)...)
}

// SumProperDivisors sums all the proper divisors of n.
func SumProperDivisors(n int) int {
	return common.Sum(ProperDivisors(n)...)
}

// NumDivisors gives the number of divisors of n.
func NumDivisors(n int) int {
	c := 1
	p := prime.Factors(n)
	for _, v := range p {
		c *= v + 1
	}
	return c
}

// GCD computes the gcd of x and y.
func GCD(x, y int) int {
	if x == 0 && y == 0 {
		return -1
	}
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// SimplifyFraction simplifies a/b and return the (numerator, denominator).
func SimplifyFraction(a, b int) (int, int) {
	gcd := GCD(a, b)
	return a / gcd, b / gcd
}

// IsAbundant determines if a number is abundant.
func IsAbundant(n int) bool {
	return SumProperDivisors(n) > n
}
