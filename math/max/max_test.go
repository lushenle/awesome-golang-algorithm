package max

import "testing"

func TestInt(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{

			name: "testMax01",
			args: args{nums: []int{1, 2, 3, 4, 5, 6}},
			want: 6,
		},
		{
			name: "testMax02",
			args: args{nums: []int{127, 128}},
			want: 128,
		},
		{
			name: "testMax03",
			args: args{nums: []int{-9223372036854770001, -9223372036854770000, 256, 333, 777}},
			want: 777,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int(tt.args.nums...); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
