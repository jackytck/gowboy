package digit

import (
	"math/big"
)

// SliceIntBig returns the individual digits of a big.Int as a slice of int.
func SliceIntBig(n *big.Int) []int {
	x := big.NewInt(0)
	if n.Cmp(x) == 0 {
		return []int{0}
	}
	var d []int
	x.Set(n)
	zero := big.NewInt(0)
	ten := big.NewInt(10)
	for x.Cmp(zero) > 0 {
		m := big.NewInt(1)
		x.DivMod(x, ten, m)
		d = append([]int{int(m.Int64())}, d...)
	}
	return d
}
