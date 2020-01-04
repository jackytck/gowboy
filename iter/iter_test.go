package iter_test

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/iter"
)

func TestPerms(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Case 1", args{[]int{1, 2, 3}}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3},
			{2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"Case 2", args{[]int{2, 3, 5, 7}}, [][]int{{2, 3, 5, 7}, {2, 3, 7, 5},
			{2, 5, 3, 7}, {2, 5, 7, 3}, {2, 7, 3, 5}, {2, 7, 5, 3}, {3, 2, 5, 7},
			{3, 2, 7, 5}, {3, 5, 2, 7}, {3, 5, 7, 2}, {3, 7, 2, 5}, {3, 7, 5, 2},
			{5, 2, 3, 7}, {5, 2, 7, 3}, {5, 3, 2, 7}, {5, 3, 7, 2}, {5, 7, 2, 3},
			{5, 7, 3, 2}, {7, 2, 3, 5}, {7, 2, 5, 3}, {7, 3, 2, 5}, {7, 3, 5, 2},
			{7, 5, 2, 3}, {7, 5, 3, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatInt(iter.Perms(tt.args.slice)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExamplePerms() {
	for v := range iter.Perms([]int{3, 12, 19}) {
		fmt.Println(v)
	}
	// Output:
	// [3 12 19]
	// [3 19 12]
	// [12 3 19]
	// [12 19 3]
	// [19 3 12]
	// [19 12 3]
}

func TestNCR(t *testing.T) {
	type args struct {
		n int
		r int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{12, 5}, big.NewInt(792)},
		{"Case 2", args{37, 13}, big.NewInt(3562467300)},
		{"Case 3", args{49, 25}, big.NewInt(63205303218876)},
		{"Case 4", args{52, 16}, big.NewInt(10363194502115)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := iter.NCR(tt.args.n, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NCR() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleNCR() {
	fmt.Println(iter.NCR(345, 123))
	// Output: 1779250012189601570865983811876820662273364126974614125481776575877025962657174021802908075728000
}

func TestCombIndex(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Case 1", args{5, 2}, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {1, 2},
			{1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{"Case 2", args{7, 3}, [][]int{{0, 1, 2}, {0, 1, 3}, {0, 1, 4}, {0, 1, 5},
			{0, 1, 6}, {0, 2, 3}, {0, 2, 4}, {0, 2, 5}, {0, 2, 6}, {0, 3, 4},
			{0, 3, 5}, {0, 3, 6}, {0, 4, 5}, {0, 4, 6}, {0, 5, 6}, {1, 2, 3},
			{1, 2, 4}, {1, 2, 5}, {1, 2, 6}, {1, 3, 4}, {1, 3, 5}, {1, 3, 6},
			{1, 4, 5}, {1, 4, 6}, {1, 5, 6}, {2, 3, 4}, {2, 3, 5}, {2, 3, 6},
			{2, 4, 5}, {2, 4, 6}, {2, 5, 6}, {3, 4, 5}, {3, 4, 6}, {3, 5, 6},
			{4, 5, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatInt(iter.CombIndex(tt.args.n, tt.args.k)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CombIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCombIndex() {
	for v := range iter.CombIndex(5, 3) {
		fmt.Println(v)
	}
	// Output:
	// [0 1 2]
	// [0 1 3]
	// [0 1 4]
	// [0 2 3]
	// [0 2 4]
	// [0 3 4]
	// [1 2 3]
	// [1 2 4]
	// [1 3 4]
	// [2 3 4]
}

func flatInt(c <-chan []int) [][]int {
	var ret [][]int
	for s := range c {
		ret = append(ret, s)
	}
	return ret
}
