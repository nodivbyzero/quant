package quant

// Charge is the dimension marker for electric charge quantities.
type Charge struct{}

// Charge units.
var (
	Coulomb      = scaleUnit[Charge]{factor: 1}
	Millicoulomb = scaleUnit[Charge]{factor: 1e-3}
	Microcoulomb = scaleUnit[Charge]{factor: 1e-6}
	Nanocoulomb  = scaleUnit[Charge]{factor: 1e-9}
	Picocoulomb  = scaleUnit[Charge]{factor: 1e-12}
)
