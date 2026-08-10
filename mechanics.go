package quant

// SurfaceTension is the dimension marker for force per length quantities.
type SurfaceTension struct{}

// Momentum is the dimension marker for momentum and impulse quantities.
type Momentum struct{}

// Surface tension units.
var (
	NewtonPerMeter    = scaleUnit[SurfaceTension]{factor: 1}
	DynePerCentimeter = scaleUnit[SurfaceTension]{factor: 0.001}
)

// Momentum units.
var (
	KilogramMeterPerSecond = scaleUnit[Momentum]{factor: 1}
	NewtonSecond           = scaleUnit[Momentum]{factor: 1}
)
