package heronian

import (
	"fmt"
	"math"
)

func New(a, b, c float64) *Triangle {
	return &Triangle{a, b, c}
}

type Triangle struct {
	a, b, c float64 // sides of the triangle
}

func (t Triangle) IsHero() bool {
	return false
}

func (t Triangle) Area() float64 {
	return t.Heron()
}

func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

func (t Triangle) SemiPerimeter() float64 {
	return t.Perimeter() / 2.0
}

// Heron is the most efficient method tested for calculating
// the area of a triangle given three sides.
func (t Triangle) Heron() float64 {
	return t.hero1()
}

// hero1 is an alternative implementation.
func (t Triangle) hero1() float64 {
	s := t.SemiPerimeter()
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}

// hero2 is an alternative implementation.
func (t Triangle) hero2() float64 {
	return math.Sqrt((t.a+t.b+t.c)*(-t.a+t.b+t.c)*(t.a-t.b+t.c)*(t.a+t.b-t.c)) / 4.0
}

// hero3 is an alternative implementation used in the
// stable() implementation.
func (t Triangle) hero3() float64 {
	tr := t.stable()
	a := tr.a // should be largest
	b := tr.b // should be middle
	c := tr.c // should be smallest
	return (a + (b + c)) * (c - (a - b)) * (c + (a - b)) * (a + (b - c))
}

// stable is an alternative implementation. It is stable when
// using floating point arithmetic for triangles containing
// very small angles.
//
// The stable alternative involves arranging the lengths of
// the sides so that a ≥ b ≥ c and computing.
func (t Triangle) stable() Triangle {
	a := t.a // should be largest
	b := t.b // should be middle
	c := t.c // should be smallest

	if b < c {
		b, c = c, b
	}

	if a < c {
		a, c = c, a
	}

	if a < b {
		a, b = b, a
	}

	return Triangle{a, b, c}
}

func (t Triangle) String() string {
	return fmt.Sprintf("Δ [%0.0f,%0.0f,%0.0f]", t.a, t.b, t.c)
}

func (t Triangle) Sides() string {
	return fmt.Sprintf("%0.0f,%0.0f,%0.0f", t.a, t.b, t.c)
}
