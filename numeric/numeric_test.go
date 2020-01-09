package numeric_test

import (
	"fmt"
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

func str2int(s string) *big.Int {
	i := new(big.Int)
	i.SetString(s, 10)
	return i
}
