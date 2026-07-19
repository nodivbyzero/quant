package quant

// ElectricalConductivity is the dimension marker for electrical conductivity quantities.
type ElectricalConductivity struct{}

// Electrical conductivity units.
var (
	SiemensPerMeter           = scaleUnit[ElectricalConductivity]{factor: 1}
	MillisiemensPerMeter      = scaleUnit[ElectricalConductivity]{factor: 1e-3}
	MicrosiemensPerMeter      = scaleUnit[ElectricalConductivity]{factor: 1e-6}
	SiemensPerCentimeter      = scaleUnit[ElectricalConductivity]{factor: 100}
	MillisiemensPerCentimeter = scaleUnit[ElectricalConductivity]{factor: 0.1}
	MicrosiemensPerCentimeter = scaleUnit[ElectricalConductivity]{factor: 1e-4}
)
