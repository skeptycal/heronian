package heronian

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFloatIsInt(t *testing.T) {
	type test struct {
		name string
		v    float64
		want bool
	}
	tests := []test{
		{"1", 1, true},
	}

	max := defaultSampleCount
	for i := 0; i < max; i++ {
		n := float64(rand.Intn(12345) + 123)
		intFlag := Flip()
		if !intFlag {
			n = n/100 + 0.12345
		}
		t := test{
			name: fmt.Sprintf("%f", n),
			v:    n,
			want: intFlag,
		}
		tests = append(tests, t)
	}
	count := 0

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FloatIsInt(tt.v)
			if got {
				count++
			}
			pct := (count * 100 / (i + 1)) - 1

			if got != tt.want {
				t.Errorf("%3d(%v%% Ints): FloatIsInt(%v) = %v, want %v", i, pct, tt.v, got, tt.want)
			}
		})
	}
}

func TestFlip(t *testing.T) {
	// TODO - this "test" isn't actually testing the functionality.
	// It just runs the function many times and verifies that the
	// percentage of true and false values are about 50/50

	n := defaultSampleCount * 1000
	trueCount := 0
	for i := 0; i < n; i++ {
		if Flip() {
			trueCount++
		}
	}
	result := trueCount * 100 / n

	name := fmt.Sprintf("percentTrue(%d)", result)

	t.Run(name, func(t *testing.T) {
		if (result+5)/10 != 5 { // this tests for 50% ± 5%
			t.Errorf("Flip() run %d times: got %v, want 50±5%%", n, result)
		}
	})
}
