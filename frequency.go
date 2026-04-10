package quant

import "math"

// Frequency is the dimension marker for frequency-like quantities.
type Frequency struct{}

// Frequency units.
var (
	Hertz               = scaleUnit[Frequency]{factor: 1}
	Millihertz          = scaleUnit[Frequency]{factor: 1e-3}
	Kilohertz           = scaleUnit[Frequency]{factor: 1e3}
	Megahertz           = scaleUnit[Frequency]{factor: 1e6}
	Gigahertz           = scaleUnit[Frequency]{factor: 1e9}
	Terahertz           = scaleUnit[Frequency]{factor: 1e12}
	RevolutionPerMinute = scaleUnit[Frequency]{factor: 1.0 / 60.0}
	DegreePerSecond     = scaleUnit[Frequency]{factor: 1.0 / 360.0}
	RadianPerSecond     = scaleUnit[Frequency]{factor: 1.0 / (2.0 * math.Pi)}
)
