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

func TestComb(t *testing.T) {
	type args struct {
		collection interface{}
		k          int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Case 1", args{"abc", 2}, []string{"ab", "ac", "bc"}},
		{"Case 2", args{"jacky", 3}, []string{"jac", "jak", "jay", "jck", "jcy",
			"jky", "ack", "acy", "aky", "cky"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatString(iter.Comb(tt.args.collection, tt.args.k)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Comb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleComb() {
	for v := range iter.Comb([]int{2, 3, 5, 7, 11, 13}, 3) {
		fmt.Println(v)
	}
	for v := range iter.Comb("gowboy", 3) {
		var s string
		for _, c := range v {
			s += string(c.(uint8))
		}
		fmt.Println(s)
	}
	// Output:
	// [2 3 5]
	// [2 3 7]
	// [2 3 11]
	// [2 3 13]
	// [2 5 7]
	// [2 5 11]
	// [2 5 13]
	// [2 7 11]
	// [2 7 13]
	// [2 11 13]
	// [3 5 7]
	// [3 5 11]
	// [3 5 13]
	// [3 7 11]
	// [3 7 13]
	// [3 11 13]
	// [5 7 11]
	// [5 7 13]
	// [5 11 13]
	// [7 11 13]
	// gow
	// gob
	// goo
	// goy
	// gwb
	// gwo
	// gwy
	// gbo
	// gby
	// goy
	// owb
	// owo
	// owy
	// obo
	// oby
	// ooy
	// wbo
	// wby
	// woy
	// boy
}

func TestCartProduct(t *testing.T) {
	type args struct {
		size   int
		repeat int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Case 1", args{0, 3}, [][]int{{0}}},
		{"Case 2", args{2, 0}, [][]int{{0}}},
		{"Case 3", args{1, 5}, [][]int{{0, 0, 0, 0, 0}}},
		{"Case 4", args{2, 3}, [][]int{{0, 0, 0}, {0, 0, 1}, {0, 1, 0},
			{0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}}},
		{"Case 5", args{4, 3}, [][]int{{0, 0, 0}, {0, 0, 1}, {0, 0, 2}, {0, 0, 3},
			{0, 1, 0}, {0, 1, 1}, {0, 1, 2}, {0, 1, 3}, {0, 2, 0}, {0, 2, 1}, {0, 2, 2},
			{0, 2, 3}, {0, 3, 0}, {0, 3, 1}, {0, 3, 2}, {0, 3, 3}, {1, 0, 0}, {1, 0, 1},
			{1, 0, 2}, {1, 0, 3}, {1, 1, 0}, {1, 1, 1}, {1, 1, 2}, {1, 1, 3}, {1, 2, 0},
			{1, 2, 1}, {1, 2, 2}, {1, 2, 3}, {1, 3, 0}, {1, 3, 1}, {1, 3, 2}, {1, 3, 3},
			{2, 0, 0}, {2, 0, 1}, {2, 0, 2}, {2, 0, 3}, {2, 1, 0}, {2, 1, 1}, {2, 1, 2},
			{2, 1, 3}, {2, 2, 0}, {2, 2, 1}, {2, 2, 2}, {2, 2, 3}, {2, 3, 0}, {2, 3, 1},
			{2, 3, 2}, {2, 3, 3}, {3, 0, 0}, {3, 0, 1}, {3, 0, 2}, {3, 0, 3}, {3, 1, 0},
			{3, 1, 1}, {3, 1, 2}, {3, 1, 3}, {3, 2, 0}, {3, 2, 1}, {3, 2, 2}, {3, 2, 3},
			{3, 3, 0}, {3, 3, 1}, {3, 3, 2}, {3, 3, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flatInt(iter.CartProduct(tt.args.size, tt.args.repeat)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CartProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCartProduct() {
	for v := range iter.CartProduct(4, 3) {
		fmt.Println(v)
	}
	// Output:
	// [0 0 0]
	// [0 0 1]
	// [0 0 2]
	// [0 0 3]
	// [0 1 0]
	// [0 1 1]
	// [0 1 2]
	// [0 1 3]
	// [0 2 0]
	// [0 2 1]
	// [0 2 2]
	// [0 2 3]
	// [0 3 0]
	// [0 3 1]
	// [0 3 2]
	// [0 3 3]
	// [1 0 0]
	// [1 0 1]
	// [1 0 2]
	// [1 0 3]
	// [1 1 0]
	// [1 1 1]
	// [1 1 2]
	// [1 1 3]
	// [1 2 0]
	// [1 2 1]
	// [1 2 2]
	// [1 2 3]
	// [1 3 0]
	// [1 3 1]
	// [1 3 2]
	// [1 3 3]
	// [2 0 0]
	// [2 0 1]
	// [2 0 2]
	// [2 0 3]
	// [2 1 0]
	// [2 1 1]
	// [2 1 2]
	// [2 1 3]
	// [2 2 0]
	// [2 2 1]
	// [2 2 2]
	// [2 2 3]
	// [2 3 0]
	// [2 3 1]
	// [2 3 2]
	// [2 3 3]
	// [3 0 0]
	// [3 0 1]
	// [3 0 2]
	// [3 0 3]
	// [3 1 0]
	// [3 1 1]
	// [3 1 2]
	// [3 1 3]
	// [3 2 0]
	// [3 2 1]
	// [3 2 2]
	// [3 2 3]
	// [3 3 0]
	// [3 3 1]
	// [3 3 2]
	// [3 3 3]
}

func flatInt(c <-chan []int) [][]int {
	var ret [][]int
	for s := range c {
		ret = append(ret, s)
	}
	return ret
}

func flatString(c <-chan []interface{}) []string {
	var ret []string
	for slice := range c {
		var s string
		for _, c := range slice {
			s += string(c.(uint8))
		}
		ret = append(ret, s)
	}
	return ret
}
