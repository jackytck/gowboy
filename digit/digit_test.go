package digit_test

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/digit"
)

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
