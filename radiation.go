package quant

// Radioactivity is the dimension marker for radioactive decay rate quantities.
type Radioactivity struct{}

// AbsorbedDose is the dimension marker for absorbed radiation dose quantities.
type AbsorbedDose struct{}

// EquivalentDose is the dimension marker for equivalent radiation dose quantities.
type EquivalentDose struct{}

// Radioactivity units.
var (
	Becquerel = scaleUnit[Radioactivity]{factor: 1}
	Curie     = scaleUnit[Radioactivity]{factor: 3.7e10}
)

// Absorbed dose units.
var (
	Gray = scaleUnit[AbsorbedDose]{factor: 1}
	Rad  = scaleUnit[AbsorbedDose]{factor: 0.01}
)

// Equivalent dose units.
var (
	Sievert = scaleUnit[EquivalentDose]{factor: 1}
	Rem     = scaleUnit[EquivalentDose]{factor: 0.01}
)
