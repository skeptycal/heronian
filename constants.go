package heronian

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
