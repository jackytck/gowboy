package digit_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/digit"
)

func TestSliceInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{0}, []int{0}},
		{"Case 2", args{123}, []int{1, 2, 3}},
		{"Case 3", args{2357111317}, []int{2, 3, 5, 7, 1, 1, 1, 3, 1, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.SliceInt(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSliceInt() {
	fmt.Println(digit.SliceInt(31415926535))
	// Output: [3 1 4 1 5 9 2 6 5 3 5]
}

func TestSliceIntBig(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{big.NewInt(0)}, []int{0}},
		{"Case 2", args{big.NewInt(1234567890)}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}},
		{"Case 3", args{big.NewInt(1123581321345589144)}, []int{1, 1, 2, 3, 5, 8, 1, 3, 2, 1, 3, 4,
			5, 5, 8, 9, 1, 4, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.SliceIntBig(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceIntBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSliceIntBig() {
	fmt.Println(digit.SliceIntBig(big.NewInt(3141592653589793238)))
	// Output: [3 1 4 1 5 9 2 6 5 3 5 8 9 7 9 3 2 3 8]
}

func TestJoinInts(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1:", args{[]int{1, 2, 3, 4, 5, 6, 7}}, 1234567},
		{"Case 2:", args{[]int{0, 2, 3, 0, 5, 0, 7}}, 230507},
		{"Case 3:", args{[]int{2, 3, 5, 7, 1}}, 23571},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.JoinInts(tt.args.slice); got != tt.want {
				t.Errorf("JoinInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleJoinInts() {
	fmt.Println(digit.JoinInts([]int{2, 3, 5, 7, 1, 1, 1, 3, 1, 7}))
	// Output: 2357111317
}

func TestJoinIntsBig(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{[]int{1, 2, 3, 4, 5, 6, 7}}, big.NewInt(1234567)},
		{"Case 2", args{[]int{0, 2, 3, 0, 5, 0, 7}}, big.NewInt(230507)},
		{"Case 3", args{[]int{2, 3, 5, 7, 1}}, big.NewInt(23571)},
		{"Case 4", args{[]int{2, 2, 4, 6, 3, 8, 7, 2, 3, 6, 9, 2, 8, 7, 8, 5, 7, 1}}, big.NewInt(224638723692878571)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.JoinIntsBig(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JoinIntsBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleJoinIntsBig() {
	fmt.Println(digit.JoinIntsBig([]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9, 3, 2, 3, 8}))
	// Output: 3141592653589793238
}

func TestSumBig(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{big.NewInt(123456789)}, big.NewInt(45)},
		{"Case 2", args{big.NewInt(3984759237523475945)}, big.NewInt(101)},
		{"Case 3", args{big.NewInt(9223372036854775807)}, big.NewInt(88)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.SumBig(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSumBig() {
	fmt.Println(digit.SumBig(big.NewInt(3141592653589793238)))
	// Output: 93
}
