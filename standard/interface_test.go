package standard

import (
	"testing"
)

var r = Rect{width: 10, height: 5}
var c = Circle{Radius: 10}

func TestRect_Area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "test01",
			fields: fields{width: 10, height: 5},
			want:   50,
		},
		{
			name:   "test02",
			fields: fields{width: 8, height: 5},
			want:   40,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRect_Perimeter(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "test01",
			fields: fields{width: 10, height: 5},
			want:   30,
		},
		{
			name:   "test02",
			fields: fields{width: 8, height: 5},
			want:   26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rect{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "test01",
			fields: fields{radius: 10},
			want:   314.1592653589793,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				Radius: tt.fields.radius,
			}
			if got := c.Area(); got != tt.want {
				t.Errorf("Area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_Perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "test01",
			fields: fields{radius: 10},
			want:   62.83185307179586,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Circle{
				Radius: tt.fields.radius,
			}
			if got := c.Perimeter(); got != tt.want {
				t.Errorf("Perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalArea(t *testing.T) {
	type args struct {
		shapes []Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test01",
			args: args{shapes: []Shape{r, c}},
			want: 364.1592653589793,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalArea(tt.args.shapes...); got != tt.want {
				t.Errorf("TotalArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalPerimeter(t *testing.T) {
	type args struct {
		shapes []Shape
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test01",
			args: args{shapes: []Shape{r, c}},
			want: 92.83185307179586,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalPerimeter(tt.args.shapes...); got != tt.want {
				t.Errorf("TotalPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
