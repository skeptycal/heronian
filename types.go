package heronian

import (
	"fmt"
	"math"
	"strings"
)

// Shape implements an example of a shape that can be
// used as a template for creating other shapes.
type Shape interface {
	Area() float64
	Perimeter() float64
	Volume() float64
	String() string
}

// shape is an unrealistic n-dimensional shape used as
// a template for creating other shapes.
type shape struct {
	dimensions []float64
}

func NewShape(d ...float64) Shape {
	return &shape{d}
}

func (s *shape) Perimeter() float64 {
	return math.NaN()
}

func (s *shape) Area() float64 {
	return math.NaN()
}

func (s *shape) Volume() float64 {
	return math.NaN()
}

func (s *shape) String() string {
	sb := strings.Builder{}
	defer sb.Reset()

	sb.WriteString("shape{")
	for _, v := range s.dimensions {
		sb.WriteString(fmt.Sprintf("%v,", v))
	}
	sb.WriteString("}")

	return sb.String()
}
