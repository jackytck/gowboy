package iter_test

import (
	"fmt"
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
		want []int
	}{
		{"Case 1", args{[]int{1, 2, 3}}, []int{1, 2, 3, 1, 3, 2, 2, 1, 3, 2, 3, 1, 3, 1, 2, 3, 2, 1}},
		{"Case 2", args{[]int{2, 3, 5, 7}}, []int{2, 3, 5, 7, 2, 3, 7, 5, 2, 5, 3, 7, 2, 5, 7, 3, 2,
			7, 3, 5, 2, 7, 5, 3, 3, 2, 5, 7, 3, 2, 7, 5, 3, 5, 2, 7, 3, 5, 7, 2, 3, 7,
			2, 5, 3, 7, 5, 2, 5, 2, 3, 7, 5, 2, 7, 3, 5, 3, 2, 7, 5, 3, 7, 2, 5, 7, 2,
			3, 5, 7, 3, 2, 7, 2, 3, 5, 7, 2, 5, 3, 7, 3, 2, 5, 7, 3, 5, 2, 7, 5, 2, 3,
			7, 5, 3, 2}},
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

func flatInt(c chan []int) []int {
	var ret []int
	for s := range c {
		ret = append(ret, s...)
	}
	return ret
}
