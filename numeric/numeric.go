package numeric

import (
	"math"
	"math/big"
	"strconv"

	"github.com/jackytck/gowboy/common"
	"github.com/jackytck/gowboy/number"
)

// Sqrt computes the square root of n with precision number of decimal digits.
// The first part gives the integer part, the second gives the integer + decimal part.
// e.g. Sqrt(2, 20) => 1, 14142135623730950488
func Sqrt(n, precision int) (int, *big.Int) {
	d := int(math.Floor(math.Sqrt(float64(n))))
	ds := strconv.Itoa(d)
	precision += len(ds)

	limit := common.Exp(10, precision+1)
	a := big.NewInt(5 * int64(n))
	b := big.NewInt(5)
	ten := big.NewInt(10)
	hundred := big.NewInt(100)
	fortyFive := big.NewInt(45)
	for b.Cmp(limit) < 0 {
		if a.Cmp(b) >= 0 {
			a.Sub(a, b)
			b.Add(b, ten)
		} else {
			a.Mul(a, hundred)
			b.Mul(b, ten)
			b.Sub(b, fortyFive)
		}
	}
	b.Div(b, hundred)
	return d, b
}

// SqrtExpand gives the expansion of continued fraction of square root of n.
// e.g. SqrtExapnd(3) => [1 1 2]
// i.e.
// 1+1/1
// 1+1/(1+1/2)
// 1+1/(1+1/(2+1/1))
// 1+1/(1+1/(2+1/(1+1/2))
// https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Continued_fraction_expansion
func SqrtExpand(n int) []int {
	srt := math.Sqrt(float64(n))
	if number.IsSquareNumber(n) {
		return []int{int(srt)}
	}
	var terms []int
	var m int
	d := 1
	a := int(math.Floor(srt))
	a0 := a
	terms = append(terms, a)
	for a != 2*a0 {
		m = d*a - m
		d = (n - m*m) / d
		a = int(math.Floor(float64((a0 + m)) / float64(d)))
		terms = append(terms, a)
	}
	return terms
}
