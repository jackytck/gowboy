package number_test

import (
	"fmt"
	"math/big"
	"reflect"
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

func TestPentagonNumber(t *testing.T) {
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
		{"Case 3", args{2}, 5},
		{"Case 4", args{3}, 12},
		{"Case 5", args{4}, 22},
		{"Case 6", args{5}, 35},
		{"Case 7", args{6}, 51},
		{"Case 8", args{7}, 70},
		{"Case 9", args{8}, 92},
		{"Case 10", args{9}, 117},
		{"Case 11", args{10}, 145},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.PentagonNumber(tt.args.n); got != tt.want {
				t.Errorf("PentagonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExamplePentagonNumber() {
	fmt.Println(number.PentagonNumber(15))
	// Output: 330
}

func TestIsPentagonNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, false},
		{"Case 2", args{1}, true},
		{"Case 3", args{5}, true},
		{"Case 4", args{52}, false},
		{"Case 5", args{71}, false},
		{"Case 6", args{3290}, true},
		{"Case 7", args{4187}, true},
		{"Case 8", args{4032}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsPentagonNumber(tt.args.n); got != tt.want {
				t.Errorf("IsPentagonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPentagonNumber() {
	fmt.Println(number.IsPentagonNumber(330))
	// Output: true
}

func TestHexagonNumber(t *testing.T) {
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
		{"Case 3", args{2}, 6},
		{"Case 4", args{3}, 15},
		{"Case 5", args{4}, 28},
		{"Case 6", args{5}, 45},
		{"Case 7", args{6}, 66},
		{"Case 8", args{7}, 91},
		{"Case 9", args{8}, 120},
		{"Case 10", args{9}, 153},
		{"Case 11", args{10}, 190},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.HexagonNumber(tt.args.n); got != tt.want {
				t.Errorf("HexagonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHexagonNumber() {
	fmt.Println(number.HexagonNumber(17))
	// Output: 561
}

func TestIsHexagonNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, false},
		{"Case 2", args{1}, true},
		{"Case 3", args{28}, true},
		{"Case 4", args{154}, false},
		{"Case 5", args{189}, false},
		{"Case 6", args{861}, true},
		{"Case 7", args{946}, true},
		{"Case 8", args{947}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsHexagonNumber(tt.args.n); got != tt.want {
				t.Errorf("IsHexagonNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsHexagonNumber() {
	fmt.Println(number.IsHexagonNumber(561))
	// Output: true
}

func TestHeptagonalNumber(t *testing.T) {
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
		{"Case 3", args{2}, 7},
		{"Case 4", args{3}, 18},
		{"Case 5", args{4}, 34},
		{"Case 6", args{5}, 55},
		{"Case 7", args{6}, 81},
		{"Case 8", args{7}, 112},
		{"Case 9", args{8}, 148},
		{"Case 10", args{9}, 189},
		{"Case 11", args{10}, 235},
		{"Case 12", args{27}, 1782},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.HeptagonalNumber(tt.args.n); got != tt.want {
				t.Errorf("HeptagonalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleHeptagonalNumber() {
	fmt.Println(number.HeptagonalNumber(19))
	// Output: 874
}

func TestIsHeptagonalNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, false},
		{"Case 2", args{1}, true},
		{"Case 3", args{55}, true},
		{"Case 4", args{111}, false},
		{"Case 5", args{149}, false},
		{"Case 6", args{235}, true},
		{"Case 7", args{1782}, true},
		{"Case 8", args{1783}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsHeptagonalNumber(tt.args.n); got != tt.want {
				t.Errorf("IsHeptagonalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsHeptagonalNumber() {
	fmt.Println(number.IsHeptagonalNumber(874))
	// Output: true
}

func TestOctagonalNumber(t *testing.T) {
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
		{"Case 3", args{2}, 8},
		{"Case 4", args{3}, 21},
		{"Case 5", args{4}, 40},
		{"Case 6", args{5}, 65},
		{"Case 7", args{6}, 96},
		{"Case 8", args{7}, 133},
		{"Case 9", args{8}, 176},
		{"Case 10", args{9}, 225},
		{"Case 11", args{10}, 280},
		{"Case 12", args{18}, 936},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.OctagonalNumber(tt.args.n); got != tt.want {
				t.Errorf("OctagonalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleOctagonalNumber() {
	fmt.Println(number.OctagonalNumber(23))
	// Output: 1541
}

func TestIsOctagonalNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, false},
		{"Case 2", args{1}, true},
		{"Case 3", args{65}, true},
		{"Case 4", args{67}, false},
		{"Case 5", args{95}, false},
		{"Case 6", args{176}, true},
		{"Case 7", args{936}, true},
		{"Case 8", args{939}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.IsOctagonalNumber(tt.args.n); got != tt.want {
				t.Errorf("IsOctagonalNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsOctagonalNumber() {
	fmt.Println(number.IsOctagonalNumber(1541))
	// Output: true
}

func TestCatalan(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{0}, big.NewInt(1)},
		{"Case 2", args{1}, big.NewInt(1)},
		{"Case 3", args{2}, big.NewInt(2)},
		{"Case 4", args{3}, big.NewInt(5)},
		{"Case 5", args{4}, big.NewInt(14)},
		{"Case 6", args{5}, big.NewInt(42)},
		{"Case 7", args{6}, big.NewInt(132)},
		{"Case 8", args{7}, big.NewInt(429)},
		{"Case 9", args{8}, big.NewInt(1430)},
		{"Case 10", args{9}, big.NewInt(4862)},
		{"Case 11", args{10}, big.NewInt(16796)},
		{"Case 12", args{11}, big.NewInt(58786)},
		{"Case 13", args{12}, big.NewInt(208012)},
		{"Case 14", args{13}, big.NewInt(742900)},
		{"Case 15", args{14}, big.NewInt(2674440)},
		{"Case 16", args{15}, big.NewInt(9694845)},
		{"Case 17", args{16}, big.NewInt(35357670)},
		{"Case 18", args{17}, big.NewInt(129644790)},
		{"Case 19", args{18}, big.NewInt(477638700)},
		{"Case 20", args{19}, big.NewInt(1767263190)},
		{"Case 21", args{20}, big.NewInt(6564120420)},
		{"Case 22", args{21}, big.NewInt(24466267020)},
		{"Case 23", args{22}, big.NewInt(91482563640)},
		{"Case 24", args{23}, big.NewInt(343059613650)},
		{"Case 25", args{24}, big.NewInt(1289904147324)},
		{"Case 26", args{25}, big.NewInt(4861946401452)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := number.Catalan(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Catalan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCatalan() {
	fmt.Println(number.Catalan(26))
	// Output: 18367353072152
}
