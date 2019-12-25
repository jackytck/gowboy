package number

import "math"

// SquareNumber returns the n-th square number.
// e.g. 1, 4, 9, 16, 25, ...
func SquareNumber(n int) int {
	return n * n
}

// IsSquareNumber determines if a number is a square number.
func IsSquareNumber(n int) bool {
	t := math.Sqrt(float64(n))
	return t == math.Floor(t)
}

// TriangleNumber returns the n-th triangle number.
// e.g. 1, 3, 6, 10, 15, ...
func TriangleNumber(n int) int {
	return n * (n + 1) / 2
}

// IsTriangleNumber determines if a number is a triangle number.
func IsTriangleNumber(n int) bool {
	return IsSquareNumber(8*n + 1)
}
