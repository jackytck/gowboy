package prime_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jackytck/gowboy/prime"
)

func TestFactors(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{"Composit 1", args{36}, map[int]int{2: 2, 3: 2}},
		{"Composit 2", args{369}, map[int]int{3: 2, 41: 1}},
		{"Prime", args{97}, map[int]int{97: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prime.Factors(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Factors() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleFactors() {
	facs := prime.Factors(98937249)
	for k, v := range facs {
		fmt.Println(k, v)
	}
	// Output:
	// 3 1
	// 32979083 1
}

func TestIsPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{1}, false},
		{"Case 2", args{2}, true},
		{"Case 3", args{3}, true},
		{"Case 4", args{4}, false},
		{"Case 5", args{1234}, false},
		{"Case 6", args{1235}, false},
		{"Case 7", args{15486803}, true},
		{"Case 8", args{15486873}, false},
		{"Case 9", args{179426447}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prime.IsPrime(tt.args.n); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleIsPrime() {
	fmt.Println(prime.IsPrime(9837249))
	// Output: false
}

func TestSieve(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Boundary 1", args{2}, []int{}},
		{"Normal 1", args{3}, []int{2}},
		{"Normal 2", args{10}, []int{2, 3, 5, 7}},
		{"Normal 3", args{100}, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59,
			61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := prime.Sieve(tt.args.n)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sieve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSieve() {
	fmt.Println(prime.Sieve(20))
	// Output: [2 3 5 7 11 13 17 19]
}

func TestRange(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"Invalid", args{2, 0}, nil},
		{"Normal 1", args{85, 100}, []int{89, 97}},
		{"Normal 2", args{700, 770}, []int{701, 709, 719, 727, 733, 739, 743, 751, 757, 761, 769}},
		{"Normal 3", args{100, 1000}, []int{101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157,
			163, 167, 173, 179, 181, 191, 193, 197, 199, 211, 223, 227, 229, 233, 239,
			241, 251, 257, 263, 269, 271, 277, 281, 283, 293, 307, 311, 313, 317, 331,
			337, 347, 349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419, 421,
			431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509,
			521, 523, 541, 547, 557, 563, 569, 571, 577, 587, 593, 599, 601, 607, 613,
			617, 619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683, 691, 701, 709,
			719, 727, 733, 739, 743, 751, 757, 761, 769, 773, 787, 797, 809, 811, 821,
			823, 827, 829, 839, 853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919,
			929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prime.Range(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleRange() {
	fmt.Println(prime.Range(300, 400))
	// Output: [307 311 313 317 331 337 347 349 353 359 367 373 379 383 389 397]
}
