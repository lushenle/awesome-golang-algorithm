package recursive

import "testing"

func TestTowersOfHanoi(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testTowersOfHanoi",
			args: args{num: 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TowersOfHanoi(tt.args.num)
		})
	}
}

func TestTOHUtil(t *testing.T) {
	type args struct {
		num  int
		from string
		to   string
		temp string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testTOHUtil",
			args: args{num: 4, from: "A", to: "B", temp: "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TOHUtil(tt.args.num, tt.args.from, tt.args.to, tt.args.temp)
		})
	}
}
