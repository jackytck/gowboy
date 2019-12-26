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

// PentagonNumber returns the n-th pentagonal number.
// e.g. 1, 5, 12, 22, 35, ...
func PentagonNumber(n int) int {
	return n * (3*n - 1) / 2
}

// IsPentagonNumber determines if a number is a pentagonal number.
func IsPentagonNumber(n int) bool {
	t := (math.Sqrt(24*float64(n)+1) + 1) / 6
	return t == math.Floor(t)
}

// HexagonNumber returns the n-th hexagonal number.
// e.g. 1, 6, 15, 28, 45, ...
func HexagonNumber(n int) int {
	return n * (2*n - 1)
}

// IsHexagonNumber determines if a number is a hexagonal number.
func IsHexagonNumber(n int) bool {
	t := (math.Sqrt(8*float64(n)+1) + 1) / 4
	return t == math.Floor(t)
}

// HeptagonalNumber returns the n-th heptagonal number.
// e.g. 1, 7, 18, 34, 55, ...
func HeptagonalNumber(n int) int {
	return n * (5*n - 3) / 2
}

// IsHeptagonalNumber determines if a number is a heptagonal number.
func IsHeptagonalNumber(n int) bool {
	t := (math.Sqrt(40*float64(n)+9) + 3) / 10
	return t == math.Floor(t)
}

// OctagonalNumber returns the n-th octagonal number.
// e.g. 1, 8, 21, 40, 65, ...
func OctagonalNumber(n int) int {
	return n * (3*n - 2)
}

// IsOctagonalNumber determines if a number is a octagonal number.
func IsOctagonalNumber(n int) bool {
	t := (math.Sqrt(3*float64(n)+1) + 1) / 3
	return t == math.Floor(t)
}
