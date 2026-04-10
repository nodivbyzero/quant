package quant

// Mass is the dimension marker for mass quantities.
type Mass struct{}

// Mass units.
var (
	Microgram = scaleUnit[Mass]{factor: 1e-9}
	Milligram = scaleUnit[Mass]{factor: 1e-6}
	Gram      = scaleUnit[Mass]{factor: 1e-3}
	Kilogram  = scaleUnit[Mass]{factor: 1}
	Ounce     = scaleUnit[Mass]{factor: 0.028349523125}
	Pound     = scaleUnit[Mass]{factor: 0.45359237}
	Stone     = scaleUnit[Mass]{factor: 6.35029318}
	MetricTon = scaleUnit[Mass]{factor: 1000}
	Tonne     = scaleUnit[Mass]{factor: 1000}
)
