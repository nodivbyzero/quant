package quant

// Time is the dimension marker for time quantities.
type Time struct{}

// Time units.
var (
	Nanosecond  = scaleUnit[Time]{factor: 1e-9}
	Microsecond = scaleUnit[Time]{factor: 1e-6}
	Millisecond = scaleUnit[Time]{factor: 1e-3}
	Second      = scaleUnit[Time]{factor: 1}
	Minute      = scaleUnit[Time]{factor: 60}
	Hour        = scaleUnit[Time]{factor: 3600}
	Day         = scaleUnit[Time]{factor: secondsPerDay}
	Week        = scaleUnit[Time]{factor: secondsPerWeek}
	Month       = scaleUnit[Time]{factor: secondsPerMonth}
	Year        = scaleUnit[Time]{factor: secondsPerYear}
	Decade      = scaleUnit[Time]{factor: 10 * secondsPerYear}
	Century     = scaleUnit[Time]{factor: 100 * secondsPerYear}
)
