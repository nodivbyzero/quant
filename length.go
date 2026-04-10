package quant

// Length is the dimension marker for length quantities.
type Length struct{}

// Length units.
var (
	Nanometer    = scaleUnit[Length]{factor: 1e-9}
	Micrometer   = scaleUnit[Length]{factor: 1e-6}
	Millimeter   = scaleUnit[Length]{factor: 1e-3}
	Centimeter   = scaleUnit[Length]{factor: 1e-2}
	Meter        = scaleUnit[Length]{factor: 1}
	Inch         = scaleUnit[Length]{factor: 0.0254}
	Yard         = scaleUnit[Length]{factor: 0.9144}
	USFoot       = scaleUnit[Length]{factor: usSurveyFootInMeters}
	Foot         = scaleUnit[Length]{factor: 0.3048}
	Fathom       = scaleUnit[Length]{factor: 1.8288}
	Kilometer    = scaleUnit[Length]{factor: 1000}
	Mile         = scaleUnit[Length]{factor: 1609.344}
	NauticalMile = scaleUnit[Length]{factor: 1852}
)
