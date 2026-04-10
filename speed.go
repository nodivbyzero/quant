package quant

// Speed is the derived dimension for length divided by time.
type Speed = Div[Length, Time]

// Speed units.
var (
	MeterPerSecond    = scaleUnit[Speed]{factor: 1}
	KilometerPerHour  = scaleUnit[Speed]{factor: 1000.0 / 3600.0}
	MilePerHour       = scaleUnit[Speed]{factor: 1609.344 / 3600.0}
	MeterPerHour      = scaleUnit[Speed]{factor: 1.0 / 3600.0}
	Knot              = scaleUnit[Speed]{factor: 1852.0 / 3600.0}
	FootPerSecond     = scaleUnit[Speed]{factor: 0.3048}
	InchPerHour       = scaleUnit[Speed]{factor: 0.0254 / 3600.0}
	MillimeterPerHour = scaleUnit[Speed]{factor: 1e-3 / 3600.0}
)
