package common

import (
	"math"
	"math/big"
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
