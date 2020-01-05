package iter

import (
	"math/big"
	"reflect"
)

// Perms return a channel of each permutation of a slice.
func Perms(slice []int) <-chan []int {
	c := make(chan []int)
	go func() {
		if len(slice) == 1 {
			c <- slice
		} else {
			for i, v := range slice {
				b := append([]int(nil), slice[:i]...)
				b = append(b, slice[i+1:]...)
				for p := range Perms(b) {
					c <- append([]int{v}, p...)
				}
			}
		}
		close(c)
	}()
	return c
}

// NCR computes the number nCr.
func NCR(n, r int) *big.Int {
	z := big.NewInt(0)
	return z.Binomial(int64(n), int64(r))
}

// CombIndex gives the combinations of selecting [0, k-1] from [0, n-1] in
// ascending order.
func CombIndex(n, k int) <-chan []int {
	c := make(chan []int)
	var comb []int
	for i := 0; i < k; i++ {
		comb = append(comb, i)
	}

	next := func() bool {
		i := k - 1
		comb[i]++
		for i > 0 && comb[i] >= n-k+1+i {
			i--
			comb[i]++
		}
		if comb[0] > n-k {
			return false
		}
		for i = i + 1; i < k; i++ {
			comb[i] = comb[i-1] + 1
		}
		return true
	}

	cop := func(a []int) []int {
		b := make([]int, len(a))
		copy(b, a)
		return b
	}

	go func() {
		defer close(c)
		c <- cop(comb)
		for next() {
			c <- cop(comb)
		}
	}()

	return c
}

// Comb gives the combinations of selecting k items from the slice of collection.
func Comb(collection interface{}, k int) <-chan []interface{} {
	col := reflect.ValueOf(collection)
	c := make(chan []interface{})
	go func() {
		defer close(c)
		for indices := range CombIndex(col.Len(), k) {
			var chosen []interface{}
			for _, i := range indices {
				chosen = append(chosen, col.Index(i).Interface())
			}
			c <- chosen
		}
	}()
	return c
}
