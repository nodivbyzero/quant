package quant

// PartsPer is the dimension marker for parts-per quantities.
type PartsPer struct{}

// Parts-per units.
var (
	Ratio = scaleUnit[PartsPer]{factor: 1}
	PPM   = scaleUnit[PartsPer]{factor: 1e-6}
	PPB   = scaleUnit[PartsPer]{factor: 1e-9}
	PPT   = scaleUnit[PartsPer]{factor: 1e-12}
	PPQ   = scaleUnit[PartsPer]{factor: 1e-15}
)
