package quant

// Temperature is the dimension marker for temperature quantities.
type Temperature struct{}

// Temperature units.
var (
	Kelvin     = affineUnit[Temperature]{factor: 1, offset: 0}
	Celsius    = affineUnit[Temperature]{factor: 1, offset: 273.15}
	Fahrenheit = affineUnit[Temperature]{factor: 5.0 / 9.0, offset: 459.67}
	Rankine    = affineUnit[Temperature]{factor: 5.0 / 9.0, offset: 0}
)
