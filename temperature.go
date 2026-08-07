package quant

// Temperature is the dimension marker for temperature quantities.
type Temperature struct{}

// Temperature units.
var (
	Kelvin     = scaleUnit[Temperature]{factor: 1, offset: 0}
	Celsius    = scaleUnit[Temperature]{factor: 1, offset: 273.15}
	Fahrenheit = scaleUnit[Temperature]{factor: 5.0 / 9.0, offset: 459.67}
	Rankine    = scaleUnit[Temperature]{factor: 5.0 / 9.0, offset: 0}
)
