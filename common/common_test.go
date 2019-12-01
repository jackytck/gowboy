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
