package quant

// Force is the dimension marker for force quantities.
type Force struct{}

// Force units.
var (
	Newton        = scaleUnit[Force]{factor: 1}
	Kilonewton    = scaleUnit[Force]{factor: 1e3}
	PoundForce    = scaleUnit[Force]{factor: 4.4482216152605}
	KilogramForce = scaleUnit[Force]{factor: standardGravity}
)
