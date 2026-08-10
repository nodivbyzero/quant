package quant

// MagneticFlux is the dimension marker for magnetic flux quantities.
type MagneticFlux struct{}

// MagneticFluxDensity is the dimension marker for magnetic flux density quantities.
type MagneticFluxDensity struct{}

// Magnetic flux units.
var (
	Weber   = scaleUnit[MagneticFlux]{factor: 1}
	Maxwell = scaleUnit[MagneticFlux]{factor: 1e-8}
)

// Magnetic flux density units.
var (
	Tesla      = scaleUnit[MagneticFluxDensity]{factor: 1}
	Millitesla = scaleUnit[MagneticFluxDensity]{factor: 1e-3}
	Gauss      = scaleUnit[MagneticFluxDensity]{factor: 1e-4}
)
