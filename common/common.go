package common

import (
	"math"
	"math/big"
	"sort"
)

// Sum returns the sum of its argument.
func Sum(a ...int) int {
	var s int
	for _, v := range a {
		s += v
	}
	return s
}

// SqrtInt computes the square root of n and casts to an int.
func SqrtInt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// Divmod gives the quotient and remainder of integer division.
func Divmod(a, b int) (int, int) {
	return a / b, a % b
}

// Exp computes x to the power of y.
func Exp(x, y int) *big.Int {
	a := big.NewInt(1)
	b := big.NewInt(int64(x))
	for i := 0; i < y; i++ {
		a.Mul(a, b)
	}
	return a
}

// Factorial returns the fractorial of a number.
func Factorial(n int) *big.Int {
	a := big.NewInt(1)
	for i := 2; i <= n; i++ {
		a.Mul(a, big.NewInt(int64(i)))
	}
	return a
}

// ProdInts computes the product of the slice of ints.
func ProdInts(slice []int) int {
	p := 1
	for _, v := range slice {
		p *= v
	}
	return p
}

// MinInt returns the smallest number.
func MinInt(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v < ret {
			ret = v
		}
	}
	return ret
}

// MaxInt returns the largest number.
func MaxInt(a ...int) int {
	var ret int
	for i, v := range a {
		if i == 0 {
			ret = v
		}
		if v > ret {
			ret = v
		}
	}
	return ret
}

// ReverseString reverses a string
func ReverseString(s string) string {
	var r string
	for _, v := range s {
		r = string(v) + r
	}
	return r
}

// CopySliceInt copy a slice of int.
func CopySliceInt(a []int) []int {
	ret := make([]int, len(a))
	copy(ret, a)
	return ret
}

// IncludesInt determines whether a slice includes a certain element.
func IncludesInt(s []int, e int) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// UniqueInt returns the set of unique and sorted int from the given slice of int.
func UniqueInt(s []int) []int {
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
