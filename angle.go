package quant

// Angle is the dimension marker for angular quantities.
type Angle struct{}

// Angle units.
var (
	Radian    = scaleUnit[Angle]{factor: 1}
	Degree    = scaleUnit[Angle]{factor: radiansPerDegree}
	Gradian   = scaleUnit[Angle]{factor: 3.141592653589793 / 200.0}
	ArcMinute = scaleUnit[Angle]{factor: radiansPerDegree / 60.0}
	ArcSecond = scaleUnit[Angle]{factor: radiansPerDegree / 3600.0}
)
