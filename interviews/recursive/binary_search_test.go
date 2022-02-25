package recursive

import "testing"

func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		data  []int
		low   int
		high  int
		value int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testBinarySearch",
			args: args{data: []int{1, 2, 3, 4, 5, 6, 7, 8}, low: 1, high: 7, value: 3},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchRecursive(tt.args.data, tt.args.low, tt.args.high, tt.args.value); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
