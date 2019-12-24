package digit

import (
	"math/big"
	"strconv"

	"github.com/jackytck/gowboy/common"
)

// SliceInt returns the individual digits as a slice of int.
func SliceInt(n int) []int {
	if n == 0 {
		return []int{0}
	}
	var d []int
	for n > 0 {
		d = append([]int{n % 10}, d...)
		n /= 10
	}
	return d
}

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

// ReverseSliceInts reverses a slice of ints, and return the new slice.
func ReverseSliceInts(a []int) []int {
	s := len(a)
	r := make([]int, s)
	for i, v := range a {
		r[s-1-i] = v
	}
	return r
}

// JoinInts joins slice of single digit ints and return it as an int.
func JoinInts(slice []int) int {
	var sum int
	len := len(slice) - 1
	for i, v := range slice {
		p := 1
		for j := 0; j < len-i; j++ {
			p *= 10
		}
		sum += p * v
	}
	return sum
}

// JoinIntsBig joins slice of single digit ints into a big.Int.
func JoinIntsBig(slice []int) *big.Int {
	sum := big.NewInt(0)
	ten := big.NewInt(10)
	p := common.Exp(10, len(slice)-1)
	for _, v := range slice {
		z := big.NewInt(1)
		z.Mul(p, big.NewInt(int64(v)))
		sum.Add(sum, z)
		p.Div(p, ten)
	}
	return sum
}

// Sum sums the digit^p of a given number.
func Sum(n, p int) *big.Int {
	s := big.NewInt(0)
	for _, c := range strconv.Itoa(n) {
		i, _ := strconv.Atoi(string(c))
		s.Add(s, common.Exp(i, p))
	}
	return s
}

// SumBig sums the digits of a big number.
func SumBig(n *big.Int) *big.Int {
	s := big.NewInt(0)
	for _, d := range SliceIntBig(n) {
		z := big.NewInt(int64(d))
		s.Add(s, z)
	}
	return s
}

// GetIth returns the i-th digit of a number n.
func GetIth(n, i int) int {
	d := SliceInt(n)
	if i < 0 || i >= len(d) {
		return -1
	}
	return d[i]
}

// ReverseInt reverses a non-negative int.
func ReverseInt(n int) int {
	return JoinInts(ReverseSliceInts(SliceInt(n)))
}

// ReverseIntBig reverses a given big.Int.
func ReverseIntBig(n *big.Int) *big.Int {
	return JoinIntsBig(ReverseSliceInts(SliceIntBig(n)))
}

// IsPalindromeInt tells if a number is a palindrome,
// i.e. reads the same backward as forward.
func IsPalindromeInt(n int) bool {
	return IsPalindromeString(strconv.Itoa(n))
}

// IsPalindromeIntBig tells if a big.Int number is a palindrome.
func IsPalindromeIntBig(n *big.Int) bool {
	return IsPalindromeString(n.String())
}

// IsPalindromeString tells if a given string is a palindrome.
func IsPalindromeString(s string) bool {
	return s == common.ReverseString(s)
}
