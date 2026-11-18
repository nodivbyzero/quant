package quant

// Torque is the dimension marker for torque quantities.
type Torque struct{}

const poundForceFootInNewtonMeters = 4.4482216152605 * 0.3048

// Torque units.
var (
	NewtonMeter    = scaleUnit[Torque]{factor: 1}
	PoundForceFoot = scaleUnit[Torque]{factor: poundForceFootInNewtonMeters}
	OunceForceInch = scaleUnit[Torque]{factor: poundForceFootInNewtonMeters / 192}
)
