package quant

// Concentration is the dimension marker for amount-of-substance concentration quantities.
type Concentration struct{}

// Molality is the dimension marker for amount-of-substance per mass quantities.
type Molality struct{}

// CatalyticActivity is the dimension marker for catalytic activity quantities.
type CatalyticActivity struct{}

// Concentration units.
var (
	MolePerCubicMeter = scaleUnit[Concentration]{factor: 1}
	Molar             = scaleUnit[Concentration]{factor: 1000}
	Millimolar        = scaleUnit[Concentration]{factor: 1}
)

// Molality units.
var (
	MolePerKilogram = scaleUnit[Molality]{factor: 1}
	Molal           = scaleUnit[Molality]{factor: 1}
)

// Catalytic activity units.
var (
	Katal = scaleUnit[CatalyticActivity]{factor: 1}
)
