package heronian

import (
	"fmt"
	"math/rand"
	"testing"
)

const (
	defaultSampleCount int     = 1000              // default number of samples to generate
	maxSampleCount     int     = 10000             // maximum number of samples to generate
	minSampleCount     int     = 5                 // minimum number of samples to generate
	maxSideLength      float64 = 20.0              // maximum length of any side of the random test triangles
	minSideLength      float64 = 1.0               // minimum length of any side of the random test triangles
	MaxUint            uint    = ^uint(0)          // maximum value of a uint variable
	MinUint            uint    = 0                 // minimum value of a uint variable
	MaxInt             int     = int(MaxUint >> 1) // maximum value of an int variable
	MinInt             int     = -MaxInt - 1       // minimum value of an int variable
)

type funcstruct struct {
	name string
	got  func(Triangle) float64
	want func() float64
}

type triangleFromSides struct {
	a float64
	b float64
	c float64
}

type triangleTestsStruct struct {
	name      string
	sides     triangleFromSides
	perimeter float64
	area      float64
}

type triangleTestSet []triangleTestsStruct

var premadeTriangleTests = triangleTestSet{
	// TODO: Add test cases.
	{"3,4,5", triangleFromSides{3, 4, 5}, 12, 6},
	{"10,13,13", triangleFromSides{10, 13, 13}, 3, 60},
	{"17,10,21", triangleFromSides{17, 10, 21}, 48, 84},
	{"5,12,13", triangleFromSides{5, 12, 13}, 30, 30},
	{"3,4,5", triangleFromSides{3, 4, 5}, 12, 6},
	{"3,4,5", triangleFromSides{3, 4, 5}, 12, 6},
	{"3,4,5", triangleFromSides{3, 4, 5}, 12, 6},
}

func side() float64 {
	return rand.Float64()*(maxSideLength-minSideLength) + minSideLength
}

func MakeTriangleTestSet(n int, addons []triangleTestsStruct) (retval []triangleTestsStruct) {
	if minSampleCount > n || n > maxSampleCount {
		n = defaultSampleCount
	}

	retval = append(retval, addons...)

	retval = append(retval, premadeTriangleTests...)

	for i := 1; i < n; i++ {
		a := side()
		b := side()
		c := side()
		tr := New(a, b, c)

		name := fmt.Sprintf("%0.0f,%0.0f,%0.0f", tr.a, tr.b, tr.c)
		ts := triangleTestsStruct{name, triangleFromSides{tr.a, tr.b, tr.c}, 0, 0}
		retval = append(retval, ts)
	}
	return
}

func TestTriangle_Perimeter(t *testing.T) {

	/// add any premade triangles to test here
	addons := triangleTestSet{
		{"1,1,1", triangleFromSides{1, 1, 1}, 12, 6},
	}

	for _, tt := range MakeTriangleTestSet(20, addons) {
		tr := Triangle{
			a: tt.sides.a,
			b: tt.sides.b,
			c: tt.sides.c,
		}
		funcs := []funcstruct{
			{"Perimeter", Triangle.Perimeter, tr.SemiPerimeter},
			{"Area", Triangle.Area, tr.SemiPerimeter},
			{"hero1", Triangle.hero1, tr.hero1},
			{"hero2", Triangle.hero2, tr.hero2},
			{"hero3", Triangle.hero3, tr.hero3},
		}
		for _, fs := range funcs {
			t.Run(tt.name, func(t *testing.T) {
				got := fs.got(tr)
				want := fs.want()
				if got != want {
					t.Errorf("Triangle.%s(%s) = %v, want %v", fs.name, tr, got, want)
				}
			})
		}

	}
}
