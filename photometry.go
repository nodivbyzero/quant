package quant

// LuminousFlux is the dimension marker for luminous flux quantities.
type LuminousFlux struct{}

// LuminousIntensity is the dimension marker for luminous intensity quantities.
type LuminousIntensity struct{}

// Luminous flux units.
var (
	Lumen = scaleUnit[LuminousFlux]{factor: 1}
)

// Luminous intensity units.
var (
	Candela = scaleUnit[LuminousIntensity]{factor: 1}
)
