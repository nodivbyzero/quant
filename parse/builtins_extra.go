package parse

import "github.com/nodivbyzero/quant"

// The accessors in this file cover the remaining concrete dimensions with
// compact SI-first registries. Applications can add domain-specific units with
// Registry.With without changing the core package.

func Torques() Registry[quant.Torque] {
	return NewRegistry[quant.Torque]().With(quant.UnitSymbol(quant.NewtonMeter), quant.NewtonMeter, "newton meter", "newton meters")
}

func Densities() Registry[quant.Density] {
	return NewRegistry[quant.Density]().With(quant.UnitSymbol(quant.KilogramPerCubicMeter), quant.KilogramPerCubicMeter, "kilogram per cubic meter", "kilograms per cubic meter")
}

func VolumeFlowRates() Registry[quant.VolumeFlowRate] {
	return NewRegistry[quant.VolumeFlowRate]().With(quant.UnitSymbol(quant.CubicMeterPerSecond), quant.CubicMeterPerSecond, "cubic meter per second", "cubic meters per second")
}

func AngularVelocities() Registry[quant.AngularVelocity] {
	return NewRegistry[quant.AngularVelocity]().With(quant.UnitSymbol(quant.AngularRadianPerSecond), quant.AngularRadianPerSecond, "radian per second", "radians per second")
}

func AngularAccelerations() Registry[quant.AngularAcceleration] {
	return NewRegistry[quant.AngularAcceleration]().With(quant.UnitSymbol(quant.AngularRadianPerSecondSquared), quant.AngularRadianPerSecondSquared, "radian per second squared", "radians per second squared")
}

func Paces() Registry[quant.Pace] {
	return NewRegistry[quant.Pace]().With(quant.UnitSymbol(quant.MinutePerMile), quant.MinutePerMile, "minute per mile", "minutes per mile")
}

func PartsPer() Registry[quant.PartsPer] {
	return NewRegistry[quant.PartsPer]().With(quant.UnitSymbol(quant.Ratio), quant.Ratio).With(quant.UnitSymbol(quant.PPM), quant.PPM, "parts per million")
}

func Pieces() Registry[quant.Pieces] {
	return NewRegistry[quant.Pieces]().With(quant.UnitSymbol(quant.Piece), quant.Piece, "piece", "pieces").With(quant.UnitSymbol(quant.Dozen), quant.Dozen)
}

func Acidities() Registry[quant.Acidity] {
	return NewRegistry[quant.Acidity]().With(quant.UnitSymbol(quant.PH), quant.PH, "ph").With(quant.UnitSymbol(quant.POH), quant.POH, "poh")
}

func Concentrations() Registry[quant.Concentration] {
	return NewRegistry[quant.Concentration]().With(quant.UnitSymbol(quant.MolePerCubicMeter), quant.MolePerCubicMeter, "mole per cubic meter", "moles per cubic meter").With(quant.UnitSymbol(quant.Molar), quant.Molar, "molar")
}

func Molalities() Registry[quant.Molality] {
	return NewRegistry[quant.Molality]().With(quant.UnitSymbol(quant.MolePerKilogram), quant.MolePerKilogram, "mole per kilogram", "moles per kilogram")
}

func DynamicViscosities() Registry[quant.DynamicViscosity] {
	return NewRegistry[quant.DynamicViscosity]().With(quant.UnitSymbol(quant.PascalSecond), quant.PascalSecond, "pascal second", "pascal seconds").With(quant.UnitSymbol(quant.Centipoise), quant.Centipoise, "centipoise")
}

func KinematicViscosities() Registry[quant.KinematicViscosity] {
	return NewRegistry[quant.KinematicViscosity]().With(quant.UnitSymbol(quant.SquareMeterPerSecond), quant.SquareMeterPerSecond, "square meter per second", "square meters per second")
}

func ThermalConductivities() Registry[quant.ThermalConductivity] {
	return NewRegistry[quant.ThermalConductivity]().With(quant.UnitSymbol(quant.WattPerMeterKelvin), quant.WattPerMeterKelvin, "watt per meter kelvin", "watts per meter kelvin")
}

func SpecificHeatCapacities() Registry[quant.SpecificHeatCapacity] {
	return NewRegistry[quant.SpecificHeatCapacity]().With(quant.UnitSymbol(quant.JoulePerKilogramKelvin), quant.JoulePerKilogramKelvin, "joule per kilogram kelvin", "joules per kilogram kelvin")
}

func Radioactivities() Registry[quant.Radioactivity] {
	return NewRegistry[quant.Radioactivity]().With(quant.UnitSymbol(quant.Becquerel), quant.Becquerel, "becquerel", "becquerels")
}

func AbsorbedDoses() Registry[quant.AbsorbedDose] {
	return NewRegistry[quant.AbsorbedDose]().With(quant.UnitSymbol(quant.Gray), quant.Gray, "gray", "grays")
}

func EquivalentDoses() Registry[quant.EquivalentDose] {
	return NewRegistry[quant.EquivalentDose]().With(quant.UnitSymbol(quant.Sievert), quant.Sievert, "sievert", "sieverts")
}

func CatalyticActivities() Registry[quant.CatalyticActivity] {
	return NewRegistry[quant.CatalyticActivity]().With(quant.UnitSymbol(quant.Katal), quant.Katal, "katal", "katals")
}

func Illuminances() Registry[quant.Illuminance] {
	return NewRegistry[quant.Illuminance]().With(quant.UnitSymbol(quant.Lux), quant.Lux, "lux")
}

func LuminousFluxes() Registry[quant.LuminousFlux] {
	return NewRegistry[quant.LuminousFlux]().With(quant.UnitSymbol(quant.Lumen), quant.Lumen, "lumen", "lumens")
}

func LuminousIntensities() Registry[quant.LuminousIntensity] {
	return NewRegistry[quant.LuminousIntensity]().With(quant.UnitSymbol(quant.Candela), quant.Candela, "candela", "candelas")
}

func MagneticFluxes() Registry[quant.MagneticFlux] {
	return NewRegistry[quant.MagneticFlux]().With(quant.UnitSymbol(quant.Weber), quant.Weber, "weber", "webers")
}

func MagneticFluxDensities() Registry[quant.MagneticFluxDensity] {
	return NewRegistry[quant.MagneticFluxDensity]().With(quant.UnitSymbol(quant.Tesla), quant.Tesla, "tesla", "teslas")
}
