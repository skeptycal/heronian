package heronian

import "math/rand"

// side returns a random triangle side length between
// minSideLength and maxSideLength.
func side() float64 {
	return rand.Float64()*(maxSideLength-minSideLength) + minSideLength
}
