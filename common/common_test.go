package common_test

import (
	"fmt"
	"testing"

	"github.com/jackytck/gowboy/common"
)

func TestSum(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 2, 3}}, 6},
		{"Case 2", args{[]int{-3, -12, 7}}, -8},
		{"Case 3", args{[]int{2, 3, 5, 7, 11, 13}}, 41},
		{"Case 4", args{[]int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}}, 376},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.Sum(tt.args.a...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSum() {
	fmt.Println(common.Sum([]int{2, 3, 5, 7, 11, 13}...))
	// Output: 41
}

func TestSqrtInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Complete", args{64}, 8},
		{"Complete 2", args{196}, 14},
		{"Non-complete", args{102}, 10},
		{"Non-complete 2", args{178}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.SqrtInt(tt.args.n); got != tt.want {
				t.Errorf("SqrtInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSqrtInt() {
	fmt.Println(common.SqrtInt(123))
	// Output: 11
}

func TestDivmod(t *testing.T) {
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
		{"Case 1", args{0, 1}, 0, 0},
		{"Case 2", args{27, 16}, 1, 11},
		{"Case 3", args{30, 3}, 10, 0},
		{"Case 4", args{35, 3}, 11, 2},
		{"Case 5", args{16, 6}, 2, 4},
		{"Case 6", args{32, 12}, 2, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := common.Divmod(tt.args.a, tt.args.b)
			if got != tt.want {
				t.Errorf("Divmod() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Divmod() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
func ExampleDivmod() {
	fmt.Println(common.Divmod(669, 13))
	// Output: 51 6
}
