package divide_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/divide"
)

func TestDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{1}, []int{1}},
		{"Case 2", args{2}, []int{1, 2}},
		{"Case 3", args{3}, []int{1, 3}},
		{"Case 4", args{4}, []int{1, 2, 4}},
		{"Case 5", args{5}, []int{1, 5}},
		{"Case 6", args{66}, []int{1, 2, 3, 6, 11, 22, 33, 66}},
		{"Case 7", args{71}, []int{1, 71}},
		{"Case 8", args{72}, []int{1, 2, 3, 4, 6, 8, 9, 12, 18, 24, 36, 72}},
		{"Case 9", args{180}, []int{1, 2, 3, 4, 5, 6, 9, 10, 12, 15, 18, 20, 30, 36, 45, 60, 90, 180}},
		{"Case 10", args{576}, []int{1, 2, 3, 4, 6, 8, 9, 12, 16, 18, 24, 32, 36, 48, 64, 72, 96, 144, 192, 288, 576}},
		{"Case 11", args{972}, []int{1, 2, 3, 4, 6, 9, 12, 18, 27, 36, 54, 81, 108, 162, 243, 324, 486, 972}},
		{"Case 12", args{3455}, []int{1, 5, 691, 3455}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.Divisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleDivisors() {
	fmt.Println(divide.Divisors(2135))
	// Output: [1 5 7 35 61 305 427 2135]
}

func TestProperDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{1}, []int{1}},
		{"Case 2", args{2}, []int{1}},
		{"Case 3", args{3}, []int{1}},
		{"Case 4", args{4}, []int{1, 2}},
		{"Case 5", args{5}, []int{1}},
		{"Case 6", args{66}, []int{1, 2, 3, 6, 11, 22, 33}},
		{"Case 7", args{71}, []int{1}},
		{"Case 8", args{72}, []int{1, 2, 3, 4, 6, 8, 9, 12, 18, 24, 36}},
		{"Case 9", args{180}, []int{1, 2, 3, 4, 5, 6, 9, 10, 12, 15, 18, 20, 30, 36, 45, 60, 90}},
		{"Case 10", args{576}, []int{1, 2, 3, 4, 6, 8, 9, 12, 16, 18, 24, 32, 36, 48, 64, 72, 96, 144, 192, 288}},
		{"Case 11", args{972}, []int{1, 2, 3, 4, 6, 9, 12, 18, 27, 36, 54, 81, 108, 162, 243, 324, 486}},
		{"Case 12", args{3455}, []int{1, 5, 691}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.ProperDivisors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleProperDivisors() {
	fmt.Println(divide.ProperDivisors(669))
	// Output: [1 3 223]
}

func TestSumDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{1}, 1},
		{"Case 2", args{2}, 3},
		{"Case 3", args{3}, 4},
		{"Case 4", args{4}, 7},
		{"Case 5", args{5}, 6},
		{"Case 6", args{66}, 144},
		{"Case 7", args{71}, 72},
		{"Case 8", args{72}, 195},
		{"Case 9", args{180}, 546},
		{"Case 10", args{576}, 1651},
		{"Case 11", args{972}, 2548},
		{"Case 12", args{3455}, 4152},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.SumDivisors(tt.args.n); got != tt.want {
				t.Errorf("SumDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumProperDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{1}, 1},
		{"Case 2", args{2}, 1},
		{"Case 3", args{3}, 1},
		{"Case 4", args{4}, 3},
		{"Case 5", args{5}, 1},
		{"Case 6", args{66}, 78},
		{"Case 7", args{71}, 1},
		{"Case 8", args{72}, 123},
		{"Case 9", args{180}, 366},
		{"Case 10", args{576}, 1075},
		{"Case 11", args{972}, 1576},
		{"Case 12", args{3455}, 697},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.SumProperDivisors(tt.args.n); got != tt.want {
				t.Errorf("SumDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSumProperDivisors() {
	fmt.Println(divide.SumProperDivisors(175))
	// Output: 73
}

func TestNumDivisors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{1}, 1},
		{"Case 2", args{2}, 2},
		{"Case 3", args{3}, 2},
		{"Case 4", args{4}, 3},
		{"Case 5", args{5}, 2},
		{"Case 6", args{6}, 4},
		{"Case 7", args{7}, 2},
		{"Case 8", args{8}, 4},
		{"Case 9", args{9}, 3},
		{"Case 10", args{10}, 4},
		{"Case 11", args{11}, 2},
		{"Case 12", args{12}, 6},
		{"Case 13", args{13}, 2},
		{"Case 14", args{14}, 4},
		{"Case 15", args{15}, 4},
		{"Case 16", args{16}, 5},
		{"Case 17", args{3455}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.NumDivisors(tt.args.n); got != tt.want {
				t.Errorf("NumDivisors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNumDivisors() {
	fmt.Println(divide.NumDivisors(2318))
	// Output: 8
}

func TestGCD(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{0, 0}, -1},
		{"Case 2", args{8, 12}, 4},
		{"Case 3", args{54, 24}, 6},
		{"Case 4", args{18, 48}, 6},
		{"Case 5", args{42, 56}, 14},
		{"Case 6", args{42, -56}, 14},
		{"Case 7", args{-35, 280}, 35},
		{"Case 8", args{252, 105}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.GCD(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("GCD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGCD() {
	fmt.Println(divide.GCD(234, 819))
	// Output: 117
}

func TestSimplifyFraction(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{"Case 1", args{2, 4}, 1, 2},
		{"Case 2", args{8, 12}, 2, 3},
		{"Case 3", args{10, 35}, 2, 7},
		{"Case 4", args{24, 108}, 2, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := divide.SimplifyFraction(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("SimplifyFraction() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SimplifyFraction() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func ExampleSimplifyFraction() {
	fmt.Println(divide.SimplifyFraction(163926, 679122))
	// Output: 7 29
}

func TestIsAbundant(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{2}, false},
		{"Case 2", args{12}, true},
		{"Case 3", args{18}, true},
		{"Case 4", args{20}, true},
		{"Case 5", args{22}, false},
		{"Case 6", args{48}, true},
		{"Case 7", args{95}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divide.IsAbundant(tt.args.n); got != tt.want {
				t.Errorf("IsAbundant() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsAbundant() {
	fmt.Println(divide.IsAbundant(105))
	fmt.Println(divide.IsAbundant(112))
	// Output:
	// false
	// true
}
