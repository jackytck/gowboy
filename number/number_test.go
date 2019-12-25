package number_test

import (
	"fmt"
	"testing"

	"github.com/jackytck/gowboy/number"
)

func TestSquareNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{0}, 0},
		{"Case 2", args{1}, 1},
		{"Case 3", args{2}, 4},
		{"Case 4", args{3}, 9},
		{"Case 5", args{4}, 16},
		{"Case 6", args{5}, 25},
		{"Case 7", args{6}, 36},
		{"Case 8", args{7}, 49},
		{"Case 9", args{8}, 64},
		{"Case 10", args{9}, 81},
		{"Case 11", args{10}, 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.SquareNumber(tt.args.n); got != tt.want {
				t.Errorf("SquareNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSquareNumber() {
	fmt.Println(number.SquareNumber(12))
	// Output: 144
}

func TestIsSquareNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, true},
		{"Case 2", args{1}, true},
		{"Case 3", args{4}, true},
		{"Case 4", args{10}, false},
		{"Case 5", args{17}, false},
		{"Case 6", args{1468944}, true},
		{"Case 7", args{1874160}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsSquareNumber(tt.args.n); got != tt.want {
				t.Errorf("IsSquareNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsSquareNumber() {
	fmt.Println(number.IsSquareNumber(144))
	// Output: true
}

func TestTriangleNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{0}, 0},
		{"Case 2", args{1}, 1},
		{"Case 3", args{2}, 3},
		{"Case 4", args{3}, 6},
		{"Case 5", args{4}, 10},
		{"Case 6", args{5}, 15},
		{"Case 7", args{6}, 21},
		{"Case 8", args{7}, 28},
		{"Case 9", args{8}, 36},
		{"Case 10", args{9}, 45},
		{"Case 11", args{10}, 55},
		{"Case 12", args{28}, 406},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.TriangleNumber(tt.args.n); got != tt.want {
				t.Errorf("TriangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleTriangleNumber() {
	fmt.Println(number.TriangleNumber(12))
	// Output: 78
}

func TestIsTriangleNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, true},
		{"Case 2", args{1}, true},
		{"Case 3", args{3}, true},
		{"Case 4", args{46}, false},
		{"Case 5", args{56}, false},
		{"Case 6", args{407}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsTriangleNumber(tt.args.n); got != tt.want {
				t.Errorf("IsTriangleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsTriangleNumber() {
	fmt.Println(number.IsTriangleNumber(78))
	// Output: true
}
