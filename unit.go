package quant

import "math"

// Dimension is the set of valid dimension markers. Dimensions are struct
// types, which prevents accidental types such as int from being used as a
// Quantity dimension while still allowing user-defined dimensions.
type Dimension interface {
	~struct{}
}

// Unit converts values between a unit and the base unit for a dimension.
// Implement ToBase and FromBase to define a custom unit outside this package.
type Unit[D Dimension] interface {
	ToBase(float64) float64
	FromBase(float64) float64
}

// BuiltinUnit is the concrete unit type used by quant's built-in units.
type BuiltinUnit[D Dimension] struct {
	factor float64
	offset float64
}

func (u BuiltinUnit[D]) ToBase(v float64) float64 {
	return (v + u.offset) * u.factor
}

func (u BuiltinUnit[D]) FromBase(v float64) float64 {
	return v/u.factor - u.offset
}

type scaleUnit[D Dimension] = BuiltinUnit[D]

const (
	usSurveyFootInMeters = 1200.0 / 3937.0
	secondsPerDay        = 24 * 60 * 60
	secondsPerWeek       = 7 * secondsPerDay
	secondsPerYear       = 365.25 * secondsPerDay
	secondsPerMonth      = secondsPerYear / 12
	standardGravity      = 9.80665
)

var (
	metersPerSquareMile = 1609.344 * 1609.344
	radiansPerDegree    = math.Pi / 180
)
