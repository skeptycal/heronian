package heronian

// DefaultRelativeErrorAllowed defines the maximum relative error allowed
// between floating point numbers. It may be used with the FuzzyEquals()
// function to determine if two numbers are equal within a given
// margin of error.
//
// This value is set to 1 ppm relative error allowed by default. It is
// fairly uncommon to have more than six significant figures in a reliable
// measurement.
//
// However, if your situation requires and allows more precision, this
// value was set as a variable instead of a constant so you may adjust
// it  as needed.
var DefaultRelativeErrorAllowed float64 = 0.000001

// FuzzyEquals is a helper function used to compare two float64 values
// and make sure they are equal within allowable tolerances
//
// The two float64 values are a and b. If a > b, the values are swapped
// before testing so that 'a' is always the smaller of the two and
// is the value that absolute error allowed is based on.
//
// RelativeErrorAllowed may be passed in if desired. If it is <= 0,
// then DefaultRelativeErrorAllowed (1ppm) is used.
func FuzzyEquals(a, b, relativeErrorAllowed float64) bool {

	if relativeErrorAllowed <= 0 {
		relativeErrorAllowed = DefaultRelativeErrorAllowed
	}

	// make sure a <= b
	if a > b {
		a, b = b, a
	}

	// absolute error of the smaller value
	absoluteErrorAllowed := a * relativeErrorAllowed

	return b-a < absoluteErrorAllowed
}
