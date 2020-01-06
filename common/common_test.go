package common_test

import (
	"fmt"
	"math/big"
	"reflect"
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

func TestExp(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{2, 2}, big.NewInt(4)},
		{"Case 2", args{2, 10}, big.NewInt(1024)},
		{"Case 3", args{5, 23}, big.NewInt(11920928955078125)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.Exp(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleExp() {
	fmt.Println(common.Exp(13, 23))
	// Output: 41753905413413116367045797
}

func TestFactorial(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{0}, big.NewInt(1)},
		{"Case 2", args{1}, big.NewInt(1)},
		{"Case 3", args{2}, big.NewInt(2)},
		{"Case 4", args{10}, big.NewInt(3628800)},
		{"Case 5", args{15}, big.NewInt(1307674368000)},
		{"Case 6", args{20}, big.NewInt(2432902008176640000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.Factorial(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFactorial() {
	fmt.Println(common.Factorial(23))
	// Output: 25852016738884976640000
}

func TestProdInts(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 2, 3, 4, 5, 6}}, 720},
		{"Case 2", args{[]int{0, 2, 3, 4, 5, 6, 3, 12}}, 0},
		{"Case 3", args{[]int{2, 3, 5, 7, 11, 13}}, 30030},
		{"Case 4", args{[]int{12, 46, 78, 9, 2, 13}}, 10075104},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.ProdInts(tt.args.slice); got != tt.want {
				t.Errorf("ProdInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleProdInts() {
	fmt.Println(common.ProdInts([]int{2, 3, 5, 7}))
	// Output: 210
}

func TestMinInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23}}, 2},
		{"Case 2", args{[]int{2, 3, 4, -1}}, -1},
		{"Case 3", args{[]int{1, 3, 1, 0, 2, 689}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.MinInt(tt.args.a...); got != tt.want {
				t.Errorf("MinInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMinInt() {
	fmt.Println(common.MinInt(2, 3, 5, 7))
	// Output: 2
}

func TestMaxInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{2, 3, 5, 7, 11, 13, 17, 19, 23}}, 23},
		{"Case 2", args{[]int{2, 3, 4, -1}}, 4},
		{"Case 3", args{[]int{1, 3, 1, 0, 2, 689}}, 689},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.MaxInt(tt.args.a...); got != tt.want {
				t.Errorf("MaxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleMaxInt() {
	fmt.Println(common.MaxInt(2, 3, 5, 7))
	// Output: 7
}

func TestReverseString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{"nat is nat"}, "tan si tan"},
		{"Case 2", args{"29374502937845"}, "54873920547392"},
		{"Case 3", args{"tacocat"}, "tacocat"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.ReverseString(tt.args.s); got != tt.want {
				t.Errorf("ReverseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleReverseString() {
	fmt.Println(common.ReverseString("a racecar"))
	// Output: racecar a
}

func TestCopySliceInt(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{[]int{}}, []int{}},
		{"Case 2", args{[]int{2, 3, 4, 1, 8, 7, 1, 5}}, []int{2, 3, 4, 1, 8, 7, 1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := common.CopySliceInt(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CopySliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCopySliceInt() {
	src := []int{1, 2, 3, 4}
	dst := common.CopySliceInt(src)
	src[1] = 5
	fmt.Println(dst)
	// Output: [1 2 3 4]
}
