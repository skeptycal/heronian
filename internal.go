package heronian

import "math/rand"

// side returns a random triangle side length between
// minSideLength and maxSideLength.
func side() float64 {
	return rand.Float64()*(maxSideLength-minSideLength) + minSideLength
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
