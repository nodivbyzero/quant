package quant

import "math"

// Unit converts values between a unit and the base unit for a dimension.
//
// Implementations should be small value types so conversions stay allocation-free
// and inline well.
type Unit[D any] interface {
	toBase(float64) float64
	fromBase(float64) float64
}

// BuiltinUnit is the concrete unit type used by quant's built-in units.
type BuiltinUnit[D any] struct {
	factor float64
	offset float64
}

func (u BuiltinUnit[D]) toBase(v float64) float64 {
	return (v + u.offset) * u.factor
}

func (u BuiltinUnit[D]) fromBase(v float64) float64 {
	return v/u.factor - u.offset
}

type scaleUnit[D any] = BuiltinUnit[D]

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
