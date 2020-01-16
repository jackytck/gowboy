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

// SqrtConvergent gives the i-th convergent of the continued fraction of the
// squart root of n.
func SqrtConvergent(n, i int) (*big.Int, *big.Int) {
	if n < 0 || i < 0 || number.IsSquareNumber(n) {
		return big.NewInt(0), big.NewInt(0)
	}
	a := SqrtExpand(n)
	if i == 0 {
		return big.NewInt(int64(a[0])), big.NewInt(1)
	}

	s := len(a) - 1
	p, q := big.NewInt(1), big.NewInt(0)
	p2, q2 := big.NewInt(0), big.NewInt(1)
	for j := 0; j <= i; j++ {
		an := a[(j-1)%s+1]
		if j == 0 {
			an = a[0]
		}

		// p2, p = p, an*p+p2
		nextP := big.NewInt(0).Set(p)
		nextP.Mul(nextP, big.NewInt(int64(an)))
		nextP.Add(nextP, p2)
		p2, p = p, nextP

		// q2, q = q, an*q+q2
		nextQ := big.NewInt(0).Set(q)
		nextQ.Mul(nextQ, big.NewInt(int64(an)))
		nextQ.Add(nextQ, q2)
		q2, q = q, nextQ
	}
	return p, q
}

// PellFundamental gives the fundamental solution (mx, my) of Pell's equaltion:
// x^2 - n * y^2 = 1, where n is a given positive nonsquare integer and integer
// solutions are sought for x and y.
func PellFundamental(n int) (*big.Int, *big.Int) {
	mx, my := big.NewInt(1), big.NewInt(0)
	if number.IsSquareNumber(n) {
		return mx, my
	}
	for i := 1; true; i++ {
		x, y := SqrtConvergent(n, i)
		mx.Set(x)
		my.Set(y)
		x.Mul(x, x)
		y.Mul(y, y)
		y.Mul(y, big.NewInt(int64(-n)))
		if x.Add(x, y); x.Int64() == 1 {
			break
		}
	}
	return mx, my
}

// LagrangePoly gives the Lagrange interpolating polynomial P(x) of
// degree <=(n-1) that passes through the n given points.
func LagrangePoly(xs []float64, ys []float64) func(float64) float64 {
	p := func(x float64) float64 {
		var s float64
		for i, u := range ys {
			m := u
			xi := xs[i]
			for j, v := range xs {
				if i == j {
					continue
				}
				m *= (x - v) / (xi - v)
			}
			s += m
		}
		return s
	}
	return p
}
