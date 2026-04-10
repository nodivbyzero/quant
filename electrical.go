package quant

// Voltage is the dimension marker for electrical potential quantities.
type Voltage struct{}

// Current is the dimension marker for electrical current quantities.
type Current struct{}

// Power is the dimension marker for active power quantities.
type Power struct{}

// ApparentPower is the dimension marker for apparent power quantities.
type ApparentPower struct{}

// ReactivePower is the dimension marker for reactive power quantities.
type ReactivePower struct{}

// Voltage units.
var (
	Volt      = scaleUnit[Voltage]{factor: 1}
	Millivolt = scaleUnit[Voltage]{factor: 1e-3}
	Kilovolt  = scaleUnit[Voltage]{factor: 1e3}
)

// Current units.
var (
	Ampere      = scaleUnit[Current]{factor: 1}
	Milliampere = scaleUnit[Current]{factor: 1e-3}
	Kiloampere  = scaleUnit[Current]{factor: 1e3}
)

// Power units.
var (
	Watt                    = scaleUnit[Power]{factor: 1}
	Milliwatt               = scaleUnit[Power]{factor: 1e-3}
	Kilowatt                = scaleUnit[Power]{factor: 1e3}
	Megawatt                = scaleUnit[Power]{factor: 1e6}
	Gigawatt                = scaleUnit[Power]{factor: 1e9}
	MetricHorsepower        = scaleUnit[Power]{factor: 735.49875}
	BTUPerSecond            = scaleUnit[Power]{factor: 1055.05585262}
	FootPoundForcePerSecond = scaleUnit[Power]{factor: 1.3558179483314004}
	Horsepower              = scaleUnit[Power]{factor: 745.6998715822702}
)

// Apparent power units.
var (
	VoltAmpere      = scaleUnit[ApparentPower]{factor: 1}
	MilliVoltAmpere = scaleUnit[ApparentPower]{factor: 1e-3}
	KiloVoltAmpere  = scaleUnit[ApparentPower]{factor: 1e3}
	MegaVoltAmpere  = scaleUnit[ApparentPower]{factor: 1e6}
	GigaVoltAmpere  = scaleUnit[ApparentPower]{factor: 1e9}
)

// Reactive power units.
var (
	VoltAmpereReactive      = scaleUnit[ReactivePower]{factor: 1}
	MilliVoltAmpereReactive = scaleUnit[ReactivePower]{factor: 1e-3}
	KiloVoltAmpereReactive  = scaleUnit[ReactivePower]{factor: 1e3}
	MegaVoltAmpereReactive  = scaleUnit[ReactivePower]{factor: 1e6}
	GigaVoltAmpereReactive  = scaleUnit[ReactivePower]{factor: 1e9}
)
