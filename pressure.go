package quant

// Pressure is the dimension marker for pressure quantities.
type Pressure struct{}

// Pressure units.
var (
	Pascal              = scaleUnit[Pressure]{factor: 1}
	Hectopascal         = scaleUnit[Pressure]{factor: 100}
	Kilopascal          = scaleUnit[Pressure]{factor: 1000}
	Megapascal          = scaleUnit[Pressure]{factor: 1e6}
	Bar                 = scaleUnit[Pressure]{factor: 1e5}
	Torr                = scaleUnit[Pressure]{factor: 101325.0 / 760.0}
	MeterOfWater        = scaleUnit[Pressure]{factor: 9806.65}
	MillimeterOfMercury = scaleUnit[Pressure]{factor: 133.322387415}
	PSI                 = scaleUnit[Pressure]{factor: 6894.757293168}
	KSI                 = scaleUnit[Pressure]{factor: 6894757.293168}
)
