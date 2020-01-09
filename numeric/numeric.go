package numeric

import (
	"math"
	"math/big"
	"strconv"

	"github.com/jackytck/gowboy/common"
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
