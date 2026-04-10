package quant

// Illuminance is the dimension marker for illuminance quantities.
type Illuminance struct{}

// Illuminance units.
var (
	Lux        = scaleUnit[Illuminance]{factor: 1}
	FootCandle = scaleUnit[Illuminance]{factor: 10.763910416709722}
)
