package quant

// Pace is the derived dimension for time divided by length.
type Pace = Div[Time, Length]

// Pace units.
var (
	SecondPerMeter     = scaleUnit[Pace]{factor: 1}
	MinutePerKilometer = scaleUnit[Pace]{factor: 60.0 / 1000.0}
	SecondPerFoot      = scaleUnit[Pace]{factor: 1.0 / 0.3048}
	MinutePerMile      = scaleUnit[Pace]{factor: 60.0 / 1609.344}
)
