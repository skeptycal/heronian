package heronian

import (
	"fmt"
	"math"
	"testing"
)

func Test_side(t *testing.T) {
	for i := 0; i < 10000; i++ {
		t.Run(fmt.Sprintf("side() test #%d", i), func(t *testing.T) {
			got := side()
			if math.IsNaN(got) || got < minSideLength || got > maxSideLength {
				t.Errorf("side() = %v (out of range)", got)
			}
		})
	}
}
