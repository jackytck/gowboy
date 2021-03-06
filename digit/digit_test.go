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

func TestReverseSliceInts(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Case 1", args{[]int{1, 2, 3, 4, 5, 6, 7}}, []int{7, 6, 5, 4, 3, 2, 1}},
		{"Case 2", args{[]int{2, 3, 5, 7, 11, 13, 17}}, []int{17, 13, 11, 7, 5, 3, 2}},
		{"Case 3", args{[]int{0, 72, 5, -12, 7, 4, 17}}, []int{17, 4, 7, -12, 5, 72, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.ReverseSliceInts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseSliceInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleReverseSliceInts() {
	fmt.Println(digit.ReverseSliceInts([]int{4, 1, 3, 1, 0, 2, 5}))
	// Output: [5 2 0 1 3 1 4]
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

func TestSum(t *testing.T) {
	type args struct {
		n int
		p int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{12, 1}, big.NewInt(3)},
		{"Case 2", args{123, 2}, big.NewInt(14)},
		{"Case 3", args{23571113, 3}, big.NewInt(533)},
		{"Case 4", args{1123581321, 11}, big.NewInt(8639121111)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.Sum(tt.args.n, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSum() {
	fmt.Println(digit.Sum(12345, 2))
	// Output: 55
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

func TestGetIth(t *testing.T) {
	type args struct {
		n int
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{123456, -123}, -1},
		{"Case 1", args{123456, 0}, 1},
		{"Case 1", args{2357111317, 7}, 3},
		{"Case 1", args{27141816, 8}, -1},
		{"Case 1", args{3141516, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.GetIth(tt.args.n, tt.args.i); got != tt.want {
				t.Errorf("DigitsIth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleGetIth() {
	fmt.Println(digit.GetIth(23418715, 5))
	// Output: 7
}

func TestReverseInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Case 1", args{0}, 0},
		{"Case 2", args{1}, 1},
		{"Case 3", args{2357111317}, 7131117532},
		{"Case 4", args{112358132134}, 431231853211},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.ReverseInt(tt.args.n); got != tt.want {
				t.Errorf("ReverseInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleReverseInt() {
	fmt.Println(digit.ReverseInt(234958))
	// Output: 859432
}

func TestReverseIntBig(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"Case 1", args{big.NewInt(0)}, big.NewInt(0)},
		{"Case 2", args{big.NewInt(1234567890987654321)}, big.NewInt(1234567890987654321)},
		{"Case 3", args{big.NewInt(1123581321345589144)}, big.NewInt(4419855431231853211)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.ReverseIntBig(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReverseIntBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleReverseIntBig() {
	fmt.Println(digit.ReverseIntBig(big.NewInt(384590274355982345)))
	// Output: 543289553472095483
}

func TestIsPalindromeInt(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0}, true},
		{"Case 2", args{12321}, true},
		{"Case 3", args{1231}, false},
		{"Case 4", args{112353211}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsPalindromeInt(tt.args.n); got != tt.want {
				t.Errorf("IsPalindromeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPalindromeInt() {
	fmt.Println(digit.IsPalindromeInt(13331))
	// Output: true
}

func TestIsPalindromeIntBig(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{big.NewInt(0)}, true},
		{"Case 2", args{big.NewInt(12321)}, true},
		{"Case 3", args{big.NewInt(1231)}, false},
		{"Case 4", args{big.NewInt(112353211)}, true},
		{"Case 5", args{big.NewInt(314159265562951413)}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsPalindromeIntBig(tt.args.n); got != tt.want {
				t.Errorf("IsPalindromeIntBig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPalindromeIntBig() {
	fmt.Println(digit.IsPalindromeIntBig(big.NewInt(12171330203317121)))
	// Output: true
}

func TestIsPalindromeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{"nat"}, false},
		{"Case 2", args{"tacocat"}, true},
		{"Case 3", args{"madam"}, true},
		{"Case 4", args{"racecar"}, true},
		{"Case 5", args{"10801"}, true},
		{"Case 6", args{"airbus"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsPalindromeString(tt.args.s); got != tt.want {
				t.Errorf("IsPalindromeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPalindromeString() {
	fmt.Println(digit.IsPalindromeString("11933316181512171330203317121518161333911"))
	// Output: true
}

func TestIsPermuted(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{0, 0}, true},
		{"Case 2", args{123, 213}, true},
		{"Case 3", args{1243, 213}, false},
		{"Case 4", args{11235, 53121}, true},
		{"Case 5", args{112035, 53121}, false},
		{"Case 6", args{112035, 503121}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsPermuted(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsPermuted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPermuted() {
	fmt.Println(digit.IsPermuted(312, 123))
	// Output: true
}

func TestIsPandigital(t *testing.T) {
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
		{"Case 3", args{312}, true},
		{"Case 4", args{3142}, true},
		{"Case 5", args{53142}, true},
		{"Case 6", args{536142}, true},
		{"Case 7", args{5376147}, false},
		{"Case 8", args{537861947}, false},
		{"Case 9", args{537861942}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsPandigital(tt.args.n); got != tt.want {
				t.Errorf("IsPandigital() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPandigital() {
	fmt.Println(digit.IsPandigital(243781569))
	fmt.Println(digit.IsPandigital(35218946))
	// Output:
	// true
	// false
}

func TestIsBouncy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{-12}, false},
		{"Case 2", args{-121}, false},
		{"Case 3", args{101}, true},
		{"Case 4", args{123}, false},
		{"Case 5", args{2341871}, true},
		{"Case 5", args{8432}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digit.IsBouncy(tt.args.n); got != tt.want {
				t.Errorf("IsBouncy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsBouncy() {
	fmt.Println(digit.IsBouncy(221))
	fmt.Println(digit.IsBouncy(121))
	// Output:
	// false
	// true
}
