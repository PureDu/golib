package math

import "math"

const (
	FLOAT32EPSILON float32 = 1.192092896e-07
	FLOAT64EPSILON float64 = 2.2204460492503131e-016
)

func IsZeroFloat32(value float32) bool {
	return IsZeroFloat64(float64(value))
}

func IsZeroFloat64(value float64) bool {
	return math.Abs(value) < FLOAT64EPSILON
}

func IsFloat64Equal(l, r float64) bool {
	return math.Abs(l-r) < FLOAT64EPSILON
}

func IsFloat32Equal(l, r float32) bool {
	return IsFloat64Equal(float64(l), float64(r))
}
