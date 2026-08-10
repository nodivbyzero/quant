package quant

// Density is the dimension marker for mass per volume quantities.
type Density struct{}

// Density units.
var (
	KilogramPerCubicMeter  = scaleUnit[Density]{factor: 1}
	GramPerCubicCentimeter = scaleUnit[Density]{factor: 1000}
	PoundPerCubicFoot      = scaleUnit[Density]{factor: 0.45359237 / 0.028316846592}
	KilogramPerLiter       = scaleUnit[Density]{factor: 1000}
)
