package quant

// Acceleration is the derived dimension for speed divided by time.
type Acceleration = Div[Speed, Time]

// Acceleration units.
var (
	MeterPerSecondSquared = scaleUnit[Acceleration]{factor: 1}
	GForce                = scaleUnit[Acceleration]{factor: standardGravity}
	StandardGravity       = scaleUnit[Acceleration]{factor: standardGravity}
)
