package quant

// ThermalConductivity is the dimension marker for thermal conductivity quantities.
type ThermalConductivity struct{}

// SpecificHeatCapacity is the dimension marker for specific heat capacity quantities.
type SpecificHeatCapacity struct{}

// Thermal conductivity units.
var (
	WattPerMeterKelvin = scaleUnit[ThermalConductivity]{factor: 1}
)

// Specific heat capacity units.
var (
	JoulePerKilogramKelvin = scaleUnit[SpecificHeatCapacity]{factor: 1}
)
