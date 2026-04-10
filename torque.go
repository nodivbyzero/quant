package quant

// Torque is the dimension marker for torque quantities.
type Torque struct{}

// Torque units.
var (
	NewtonMeter    = scaleUnit[Torque]{factor: 1}
	PoundForceFoot = scaleUnit[Torque]{factor: 1.3558179483314004}
)
