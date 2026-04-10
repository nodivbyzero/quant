package quant

// Area is the dimension marker for area quantities.
type Area struct{}

// Area units.
var (
	SquareMillimeter = scaleUnit[Area]{factor: 1e-6}
	SquareCentimeter = scaleUnit[Area]{factor: 1e-4}
	SquareMeter      = scaleUnit[Area]{factor: 1}
	Hectare          = scaleUnit[Area]{factor: 10000}
	SquareKilometer  = scaleUnit[Area]{factor: 1e6}
	SquareInch       = scaleUnit[Area]{factor: 0.00064516}
	SquareFoot       = scaleUnit[Area]{factor: 0.09290304}
	Acre             = scaleUnit[Area]{factor: 4046.8564224}
	SquareMile       = scaleUnit[Area]{factor: metersPerSquareMile}
)
