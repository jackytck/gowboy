package divide

import (
	"github.com/jackytck/gowboy/common"
	"github.com/jackytck/gowboy/prime"
)

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
