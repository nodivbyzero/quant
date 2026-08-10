package quant

import "math"

// AngularVelocity is the dimension marker for rotational rate quantities.
type AngularVelocity struct{}

// AngularAcceleration is the dimension marker for rotational acceleration quantities.
type AngularAcceleration struct{}

// Angular velocity units.
var (
	AngularRadianPerSecond     = scaleUnit[AngularVelocity]{factor: 1}
	AngularDegreePerSecond     = scaleUnit[AngularVelocity]{factor: math.Pi / 180}
	AngularRevolutionPerMinute = scaleUnit[AngularVelocity]{factor: 2 * math.Pi / 60}
)

// Angular acceleration units.
var (
	AngularRadianPerSecondSquared = scaleUnit[AngularAcceleration]{factor: 1}
	AngularDegreePerSecondSquared = scaleUnit[AngularAcceleration]{factor: math.Pi / 180}
)
