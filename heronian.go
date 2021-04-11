// Package heronian is a playground for testing Heronian (or Hero) triangles.
package heronian

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func sortDec(s sort.Float64Slice) sort.Float64Slice {
	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	return s
}

func New(a, b, c float64) *Triangle {
	s := []float64{a, b, c}

	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	sort.Float64s(s)

	return &Triangle{s}
}

type Triangle struct {
	sides sort.Float64Slice // sides of the triangle
}

func (t *Triangle) a() float64 {
	return t.sides[0]
}

func (t *Triangle) b() float64 {
	return t.sides[1]
}

func (t *Triangle) c() float64 {
	return t.sides[2]
}

func (t Triangle) IsHero() bool {
	return false
}

func (t Triangle) Area() float64 {
	return t.Heron()
}

func (t Triangle) Perimeter() float64 {
	return t.a() + t.b() + t.c()
}

func (t Triangle) SemiPerimeter() float64 {
	return t.Perimeter() / 2
}

// Heron is the most efficient method tested for calculating
// the area of a triangle given three sides.
func (t Triangle) Heron() float64 {
	return t.hero1()
}

// hero1 is an alternative implementation.
func (t Triangle) hero1() float64 {
	a := t.a()
	b := t.b()
	c := t.c()
	s := t.SemiPerimeter()
	n := s * (s - a) * (s - b) * (s - c)
	return math.Sqrt(n)
}

// hero2 is an alternative implementation.
func (t Triangle) hero2() float64 {
	a := t.a()
	b := t.b()
	c := t.c()
	return math.Sqrt((a+b+c)*(-a+b+c)*(a-b+c)*(a+b-c)) / 4.0
}

// hero3 is an alternative implementation used in the
// stable() implementation.
func (t Triangle) hero3() float64 {
	tr := t.stable()
	a := tr.a() // should be largest
	b := tr.b() // should be middle
	c := tr.c() // should be smallest
	return (a + (b + c)) * (c - (a - b)) * (c + (a - b)) * (a + (b - c))
}

// stable is an alternative implementation. It is stable when
// using floating point arithmetic for triangles containing
// very small angles.
//
// The stable alternative involves arranging the lengths of
// the sides so that a ≥ b ≥ c and computing.
func (t Triangle) stable() Triangle {
	a := t.a() // should be largest
	b := t.b() // should be middle
	c := t.c() // should be smallest

	s := sort.Float64Slice([]float64{a, b, c})
	// sort.SearchFloat64s(a []float64, x float64)
	sort.Sort(sort.Reverse(sort.Float64Slice(s)))

	return Triangle{s}
}

func (t Triangle) String() string {
	return fmt.Sprintf("Δ [%0.0f,%0.0f,%0.0f]", t.a(), t.b(), t.c())
}

func (t Triangle) Sides() string {
	return fmt.Sprintf("%0.0f,%0.0f,%0.0f", t.a(), t.b(), t.c())
}
