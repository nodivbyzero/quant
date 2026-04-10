package quant

// Volume is the dimension marker for volume quantities.
type Volume struct{}

// Volume units.
var (
	CubicMillimeter = scaleUnit[Volume]{factor: 1e-9}
	CubicCentimeter = scaleUnit[Volume]{factor: 1e-6}
	Milliliter      = scaleUnit[Volume]{factor: 1e-6}
	Liter           = scaleUnit[Volume]{factor: 1e-3}
	Kiloliter       = scaleUnit[Volume]{factor: 1}
	Megaliter       = scaleUnit[Volume]{factor: 1e3}
	Gigaliter       = scaleUnit[Volume]{factor: 1e6}
	CubicMeter      = scaleUnit[Volume]{factor: 1}
	CubicKilometer  = scaleUnit[Volume]{factor: 1e9}
	Teaspoon        = scaleUnit[Volume]{factor: 4.92892159375e-6}
	Tablespoon      = scaleUnit[Volume]{factor: 1.478676478125e-5}
	CubicInch       = scaleUnit[Volume]{factor: 1.6387064e-5}
	FluidOunce      = scaleUnit[Volume]{factor: 2.95735295625e-5}
	Cup             = scaleUnit[Volume]{factor: 2.365882365e-4}
	Pint            = scaleUnit[Volume]{factor: 4.73176473e-4}
	Quart           = scaleUnit[Volume]{factor: 9.46352946e-4}
	Gallon          = scaleUnit[Volume]{factor: 3.785411784e-3}
	CubicFoot       = scaleUnit[Volume]{factor: 0.028316846592}
	CubicYard       = scaleUnit[Volume]{factor: 0.764554857984}
)
