package search

import "testing"

func TestBinary(t *testing.T) {
	type args struct {
		nums      []int
		target    int
		lowIndex  int
		highIndex int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "testBinary01",
			args: args{
				nums:      []int{1, 4, 5, 7, 8, 9, 10, 11},
				target:    8,
				lowIndex:  0,
				highIndex: 7,
			},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Binary(tt.args.nums, tt.args.target, tt.args.lowIndex, tt.args.highIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Binary() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Binary() got = %v, want %v", got, tt.want)
			}
		})
	}
}
