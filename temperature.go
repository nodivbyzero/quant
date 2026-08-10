package quant

// Temperature is the dimension marker for temperature quantities.
type Temperature struct{}

// TemperatureDelta is the dimension marker for temperature difference quantities.
type TemperatureDelta struct{}

// Temperature units.
var (
	Kelvin     = scaleUnit[Temperature]{factor: 1, offset: 0}
	Celsius    = scaleUnit[Temperature]{factor: 1, offset: 273.15}
	Fahrenheit = scaleUnit[Temperature]{factor: 5.0 / 9.0, offset: 459.67}
	Rankine    = scaleUnit[Temperature]{factor: 5.0 / 9.0, offset: 0}
)

// Temperature delta units.
var (
	KelvinDelta     = scaleUnit[TemperatureDelta]{factor: 1}
	CelsiusDelta    = scaleUnit[TemperatureDelta]{factor: 1}
	FahrenheitDelta = scaleUnit[TemperatureDelta]{factor: 5.0 / 9.0}
	RankineDelta    = scaleUnit[TemperatureDelta]{factor: 5.0 / 9.0}
)

// AddDelta adds a temperature difference to an absolute temperature.
func (q Quantity[Temperature]) AddDelta(delta Quantity[TemperatureDelta]) Quantity[Temperature] {
	return Quantity[Temperature]{value: q.value + delta.value}
}

// SubDelta subtracts a temperature difference from an absolute temperature.
func (q Quantity[Temperature]) SubDelta(delta Quantity[TemperatureDelta]) Quantity[Temperature] {
	return Quantity[Temperature]{value: q.value - delta.value}
}
