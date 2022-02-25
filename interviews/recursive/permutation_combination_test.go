package recursive

import (
	"fmt"
	"testing"
)

func TestPermutation(t *testing.T) {
	type args struct {
		data   []int
		n      int
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testPermutation01",
			args: args{data: []int{1, 2, 3}, n: 0, length: 2},
		},
		{
			name: "testPermutation02",
			args: args{data: []int{1, 2, 3}, n: 1, length: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Permutation(tt.args.data, tt.args.n, tt.args.length)
		})
	}
}

func Test_printSlice(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testPrintSlice",
			args: args{data: []int{1, 2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printSlice(tt.args.data)
		})
	}
}

func Test_swap(t *testing.T) {
	type args struct {
		data []int
		a    int
		b    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testSwap01",
			args: args{data: []int{1, 2, 3, 4, 5}, a: 1, b: 4},
			want: []int{1, 5, 3, 4, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Input: %v, Expect: %v\t", tt.args.data, tt.want)
			swap(tt.args.data, tt.args.a, tt.args.b)
			fmt.Printf("Output: %v\n", tt.args.data)
		})
	}
}
