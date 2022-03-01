package exercises

import (
	"reflect"
	"testing"
)

func TestAverage(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testAverage01",
			args: args{data: []int{1, 2, 3}},
			want: 2,
		},
		{
			name: "testAverage02",
			args: args{data: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "testAverage03",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.data); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfMatrixElement(t *testing.T) {
	type args struct {
		data [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSumOfMatrixElement01",
			args: args{data: [][]int{{1, 2}, {3, 4}}},
			want: 10,
		},
		{
			name: "testSumOfMatrixElement02",
			args: args{data: [][]int{{-1, -2}, {-3, -4}}},
			want: -10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfMatrixElement(tt.args.data); got != tt.want {
				t.Errorf("SumOfMatrixElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxElement(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testMaxElement01",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: 5,
		},
		{
			name: "testMaxElement02",
			args: args{data: []int{-1, -2, -3, -4, -5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxElement(tt.args.data); got != tt.want {
				t.Errorf("MaxElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSmallestElement(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSmallestElement01",
			args: args{data: []int{1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "testSmallestElement02",
			args: args{data: []int{-1, -2, -3, -4, -5}},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SmallestElement(tt.args.data); got != tt.want {
				t.Errorf("SmallestElement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSecondMax(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSecondMax",
			args: args{data: []int{-1, -2, -3, -4, -5}},
			want: -2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SecondMax(tt.args.data); got != tt.want {
				t.Errorf("SecondMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPermutations(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testPermutations",
			args: args{nums: []int{1, 2, 3}},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Permutations(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Permutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumN(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSumX01",
			args: args{n: 10},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumN(tt.args.n); got != tt.want {
				t.Errorf("SumN() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumN1(tt.args.n); got != tt.want {
				t.Errorf("SumN1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxima(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testMaxima01",
			args: args{data: []int{7, 2, 4, 3, 6, 8}},
			want: []int{7, 4, 8},
		},
		{
			name: "testMaxima02",
			args: args{data: []int{7, 2, 4, 3, 6, 2}},
			want: []int{7, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Maxima(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Maxima() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testMergeIntervals01",
			args: args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name: "testMergeIntervals02",
			args: args{[][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeIntervals(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testReverse",
			args: args{nums: []int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumOfDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testSumOfDigits01",
			args: args{num: 1894},
			want: 22,
		},
		{
			name: "testSumOfDigits01",
			args: args{num: 1111},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfDigits(tt.args.num); got != tt.want {
				t.Errorf("SumOfDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegregate0and1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testSegregate0and101",
			args: args{nums: []int{0, 1, 0, 1, 0}},
			want: []int{0, 0, 0, 1, 1},
		},
		{
			name: "testSegregate0and102",
			args: args{nums: []int{0, 1, 0, 1, 1, 0}},
			want: []int{0, 0, 0, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Segregate0and1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SegregateSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSegregate0and1and2(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testSegregate0and1and201",
			args: args{nums: []int{0, 1, 2, 2, 1, 0}},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "testSegregate0and1and202",
			args: args{nums: []int{2, 1, 0, 1, 1, 0}},
			want: []int{0, 0, 1, 1, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Segregate0and1and2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Segregate0and1and2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testFindDuplicates",
			args: args{nums: []int{1, 1, 1, 2, 2, 3, 0}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMax(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFindMax01",
			args: args{nums: []int{5, 6, 7, 8, 1, 2, 3, 4}},
			want: 8,
		},
		{
			name: "testFindMax01",
			args: args{nums: []int{50, 60, 70, 1, 2, 3, 4}},
			want: 70,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMax(tt.args.nums); got != tt.want {
				t.Errorf("FindMaximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testFindMin01",
			args: args{nums: []int{4, 5, 6, 1, 2, 3}},
			want: 1,
		},
		{
			name: "testFindMin02",
			args: args{nums: []int{-4, -5, -6, -1, -2, -3}},
			want: -6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMin(tt.args.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testRemoveDuplicates",
			args: args{nums: []int{1, 2, 3, 2, 1, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPairs(t *testing.T) {
	type args struct {
		nums []int
		x    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testCheckPairs01",
			args: args{nums: []int{2, 7, 11, 15}, x: 9},
			want: []int{0, 1},
		},
		{
			name: "testCheckPairs02",
			args: args{nums: []int{3, 2, 4}, x: 6},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPairs(tt.args.nums, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
