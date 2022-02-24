package array

import (
	"fmt"
	"testing"
)

func TestSumArray(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSumArry01",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumArray(tt.args.data); got != tt.want {
				t.Errorf("SumArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSequentialSearch(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testSequentialSearch01",
			args: args{data: []int{1, 2, 3, 4, 5}, value: 3},
			want: true,
		},
		{
			name: "testSequentialSearch02",
			args: args{data: []int{1, 2, 3, 4, 5}, value: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SequentialSearch(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("SequentialSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinarySearch(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testBinarySearch01",
			args: args{data: []int{1, 2, 3, 4, 5}, value: 2},
			want: true,
		},
		{
			name: "testBinarySearch02",
			args: args{data: []int{1, 2, 3, 4, 5}, value: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseArray(t *testing.T) {
	type args struct {
		data  []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testReverseArray01",
			args: args{data: []int{1, 2, 3, 4, 5}, start: 1, end: 4},
			want: []int{1, 5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Input: %v, Expect: %v\t", tt.args.data, tt.want)
			ReverseArray(tt.args.data, tt.args.start, tt.args.end)
			fmt.Printf("Output: %v\n", tt.args.data)
		})
	}
}

func TestRotateArray(t *testing.T) {
	type args struct {
		data []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testRotateArray01",
			args: args{data: []int{1, 2, 3, 4, 5}, k: 2},
			want: []int{3, 4, 5, 1, 2},
		},
		{
			name: "testRotateArray02",
			args: args{data: []int{1, 2, 3, 4, 5}, k: 3},
			want: []int{4, 5, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("Input: %v, Expect: %v\t", tt.args.data, tt.want)
			RotateArray(tt.args.data, tt.args.k)
			fmt.Printf("Output: %v\n", tt.args.data)
		})
	}
}

func TestMaxSubArraySum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMaxSubArraySum01",
			args: args{data: []int{1, -2, 3, 4, -4, 6, -14, 8, 2}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArraySum(tt.args.data); got != tt.want {
				t.Errorf("MaxSubArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
