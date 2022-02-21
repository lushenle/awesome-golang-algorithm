package standard

import "math"

// Shape which contain two methods Area() and Perimeter()
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rect and Circle implements Shape interface as they implement Area() and Perimeter() methods

// Rect length and width of the rectangle
type Rect struct {
	width  float64
	height float64
}

// Circle radius of the circle
type Circle struct {
	Radius float64
}

// Area the implementation of rectangle area calculation method
func (r Rect) Area() float64 {
	return r.width * r.height
}

// Perimeter the implementation of rectangle perimeter calculation method
func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Area the implementation of circle area calculation method
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

// Perimeter the implementation of circle perimeter calculation method
func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}

// TotalArea() and TotalPerimeter() are two functions which expect list of object of type Shape

// TotalArea calculation all shape areas
func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}

	return area
}

// TotalPerimeter calculation all shape perimeter
func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, p := range shapes {
		peri += p.Perimeter()
	}

	return peri
}
