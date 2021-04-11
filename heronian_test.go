package heronian

import (
	"fmt"
	"math"
	"testing"
)

func TestTriangle(t *testing.T) {

	addons := triangleTestSet{
		MakeTriangleTest(1, 1, 1),
		MakeTriangleTest(5, 5, 6),
		MakeTriangleTest(3, 4, 5),
		MakeTriangleTest(51, 52, 53),
	}

	for testNum, tt := range MakeTriangleTestSet(5, addons) {

		tr := New(tt.sides.a, tt.sides.b, tt.sides.c)

		funcs := []funcstruct{
			{"Perimeter", Triangle.Perimeter, tr.Perimeter},
			{"HeronArea", Triangle.HeronArea, tr.HeronArea},
			{"SemiPerimeter", Triangle.SemiPerimeter, tr.SemiPerimeter},
			{"hero1", Triangle.hero1, tr.Area},
			{"hero2", Triangle.hero2, tr.Area},
			{"hero3", Triangle.hero3, tr.Area},
		}
		for _, fs := range funcs {
			t.Run(tt.name, func(t *testing.T) {
				got := fs.got(*tr)
				want := fs.want()
				if !math.IsNaN(got) && !math.IsNaN(want) { // there really shouldn't be NaN values ... but eh.
					if !FuzzyEquals(got, want, 0) {
						t.Errorf("#%3d: %s(%s, p: %0.1f, a: %0.1f) = %v, want %v", testNum, fs.name, tr, tr.Perimeter(), tr.Area(), got, want)
					}
				}
			})
		}

	}
}

type funcstruct struct {
	name string
	got  func(Triangle) float64
	want func() float64
}

type triangleSides struct {
	a float64
	b float64
	c float64
}

type triangleTestsStruct struct {
	name      string
	sides     triangleSides
	perimeter float64
	area      float64
}

type triangleTestSet []triangleTestsStruct

var premadeTriangleTests = triangleTestSet{
	MakeTriangleTest(3, 4, 5),
	MakeTriangleTest(5, 5, 5),
	MakeTriangleTest(10, 13, 13),
	MakeTriangleTest(17, 10, 21),
	MakeTriangleTest(5, 12, 13),
}

func MakeTriangleTest(a, b, c float64) triangleTestsStruct {
	if a < minSideLength || a > maxSideLength {
		a = side()
	}
	if b < minSideLength || b > maxSideLength {
		b = side()
	}
	if c < minSideLength || c > maxSideLength {
		c = side()
	}

	tr := New(a, b, c)

	name := fmt.Sprintf("%0.0f,%0.0f,%0.0f", tr.a(), tr.b(), tr.c())
	return triangleTestsStruct{name, triangleSides{tr.a(), tr.b(), tr.c()}, tr.Perimeter(), tr.Area()}
}

func MakeTriangleTestSet(n int, addons []triangleTestsStruct) (retval []triangleTestsStruct) {
	if minSampleCount > n || n > maxSampleCount {
		n = defaultSampleCount
	}

	// add tests passed as parameters
	retval = append(retval, addons...)

	// add premade tests
	retval = append(retval, premadeTriangleTests...)

	// add n random tests
	for i := 1; i < n; i++ {
		retval = append(retval, MakeTriangleTest(0, 0, 0))
	}
	return
}

func Test_FuzzyEqual(t *testing.T) {
	const fuzzyTestRelativeErrorAllowed float64 = 0.000001 // 1 ppm for tests

	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// using fuzzyTestRelativeErrorAllowed (1ppm) for testing
		{"", args{1, 2}, false},
		{"", args{100, 200}, false},
		{"", args{199, 200}, false},
		{"", args{1999999, 2000000}, true},
		{"", args{999998, 1000000}, false},  // 2 ppm
		{"", args{999999, 1000000}, false},  // 1 ppm
		{"", args{999999.9, 1000000}, true}, // 0.1 ppm
		{"", args{999999.5, 1000000}, true}, // 0.5 ppm
	}
	for _, tt := range tests {
		a := tt.args.a
		b := tt.args.b
		if a > b {
			a, b = b, a
		}
		aErrorAllowed := a * fuzzyTestRelativeErrorAllowed
		absError := b - a

		t.Run(tt.name, func(t *testing.T) {
			if got := FuzzyEquals(tt.args.a, tt.args.b, fuzzyTestRelativeErrorAllowed); got != tt.want {
				t.Errorf("fuzzyEqual(%0.0f, %0.0f) (allowed: %f, absError: %f) got %v, want %v", tt.args.a, tt.args.b, aErrorAllowed, absError, got, tt.want)
			}
		})
	}

	// test for custom relativeErrorAllowed
	t.Run("custom relativeErrorAllowed", func(t *testing.T) {
		a := 1.9
		b := 2.0
		customRelativeErrorAllowed := 0.1 // 10% error allowed
		got := FuzzyEquals(a, b, customRelativeErrorAllowed)
		want := true
		if got != want {
			t.Errorf("fuzzyEqual(%0.1f, %0.1f) (allowed: %f, absError: %f) got %v, want %v", a, b, a*customRelativeErrorAllowed, b-a, got, want)
		}
	})

	// test for custom relativeErrorAllowed out of range triggering default
	// also, b < a ... so should trigger value swap of a and b
	t.Run("custom relativeErrorAllowed", func(t *testing.T) {
		a := 2.0
		b := 1.9
		customRelativeErrorAllowed := -1.0 // 10% error allowed
		got := FuzzyEquals(a, b, customRelativeErrorAllowed)
		want := false
		if got != want {
			t.Errorf("fuzzyEqual(%0.1f, %0.1f) (allowed: %f, absError: %f) got %v, want %v", a, b, a*DefaultRelativeErrorAllowed, a-b, got, want)
		}
	})
}
