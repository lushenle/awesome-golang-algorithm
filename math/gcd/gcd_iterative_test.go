package gcd

import "testing"

func TestIterative(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testGcd01",
			args: args{a: 10, b: 8},
			want: 2,
		},
		{
			name: "testGcd02",
			args: args{a: 17, b: 13},
			want: 1,
		},
		{
			name: "testGcd03",
			args: args{a: 8, b: 4},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Iterative(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Iterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
