package numeric_test

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/numeric"
)

func TestSqrt(t *testing.T) {
	type args struct {
		n         int
		precision int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 *big.Int
	}{
		{"Case 1", args{2, 60}, 1, str2int("1414213562373095048801688724209698078569671875376948073176679")},
		{"Case 2", args{3, 60}, 1, str2int("1732050807568877293527446341505872366942805253810380628055806")},
		{"Case 3", args{5, 60}, 2, str2int("2236067977499789696409173668731276235440618359611525724270897")},
		{"Case 4", args{10, 60}, 3, str2int("3162277660168379331998893544432718533719555139325216826857504")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := numeric.Sqrt(tt.args.n, tt.args.precision)
			if got != tt.want {
				t.Errorf("Sqrt() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Sqrt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func ExampleSqrt() {
	fmt.Println(numeric.Sqrt(23418, 20))
	// Output: 153 15302940893828218664881
}

func TestSqrtExapnd(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{2}, []int{1, 2}},
		{"Case 2", args{3}, []int{1, 1, 2}},
		{"Case 3", args{17}, []int{4, 8}},
		{"Case 4", args{64}, []int{8}},
		{"Case 5", args{144}, []int{12}},
		{"Case 6", args{234}, []int{15, 3, 2, 1, 2, 1, 2, 3, 30}},
		{"Case 7", args{3889}, []int{62, 2, 1, 3, 4, 2, 1, 7, 1, 1, 1, 1, 1, 17, 5, 7, 7, 5, 17, 1,
			1, 1, 1, 1, 7, 1, 2, 4, 3, 1, 2, 124}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numeric.SqrtExpand(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SqrtExapnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSqrtExpand() {
	fmt.Println(numeric.SqrtExpand(12))
	est := 3 + 1.0/(2+1.0/(6+1.0/(2+1.0/6)))
	act := math.Sqrt(12)
	fmt.Printf("Actual: %.5f Estimate: %.5f\n", act, est)
	// Output:
	// [3 2 6]
	// Actual: 3.46410 Estimate: 3.46409
}

func str2int(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}
