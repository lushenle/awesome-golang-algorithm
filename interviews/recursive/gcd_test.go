package recursive

import "testing"

func TestGcd(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testGcd01",
			args: args{m: 4, n: 5},
			want: 1,
		},
		{
			name: "testGcd02",
			args: args{m: 4, n: 6},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gcd(tt.args.m, tt.args.n); got != tt.want {
				t.Errorf("Gcd() = %v, want %v", got, tt.want)
			}
		})
	}
}
