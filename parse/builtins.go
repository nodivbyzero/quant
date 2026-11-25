package parse

import "github.com/nodivbyzero/quant"

// Areas returns an immutable registry for common area units.
func Areas() Registry[quant.Area] {
	return NewRegistry[quant.Area]().
		With(quant.UnitSymbol(quant.SquareMeter), quant.SquareMeter, "square meter", "square meters").
		With(quant.UnitSymbol(quant.SquareMillimeter), quant.SquareMillimeter, "square millimeter", "square millimeters").
		With(quant.UnitSymbol(quant.SquareCentimeter), quant.SquareCentimeter, "square centimeter", "square centimeters").
		With(quant.UnitSymbol(quant.SquareKilometer), quant.SquareKilometer, "square kilometer", "square kilometers").
		With(quant.UnitSymbol(quant.SquareInch), quant.SquareInch, "square inch", "square inches").
		With(quant.UnitSymbol(quant.SquareFoot), quant.SquareFoot, "square foot", "square feet").
		With(quant.UnitSymbol(quant.Acre), quant.Acre, "acres").
		With(quant.UnitSymbol(quant.Hectare), quant.Hectare, "hectare", "hectares")
}

// Volumes returns an immutable registry for common volume units.
func Volumes() Registry[quant.Volume] {
	return NewRegistry[quant.Volume]().
		With(quant.UnitSymbol(quant.CubicMeter), quant.CubicMeter, "cubic meter", "cubic meters").
		With(quant.UnitSymbol(quant.Liter), quant.Liter, "liter", "liters", "litre", "litres").
		With(quant.UnitSymbol(quant.Milliliter), quant.Milliliter, "milliliter", "milliliters").
		With(quant.UnitSymbol(quant.Gallon), quant.Gallon, "gallon", "gallons").
		With(quant.UnitSymbol(quant.CubicFoot), quant.CubicFoot, "cubic foot", "cubic feet").
		With(quant.UnitSymbol(quant.CubicInch), quant.CubicInch, "cubic inch", "cubic inches").
		With(quant.UnitSymbol(quant.Cup), quant.Cup, "cups").
		With(quant.UnitSymbol(quant.Pint), quant.Pint, "pint", "pints").
		With(quant.UnitSymbol(quant.Quart), quant.Quart, "quart", "quarts")
}

// Speeds returns an immutable registry for common speed units.
func Speeds() Registry[quant.Speed] {
	return NewRegistry[quant.Speed]().
		With(quant.UnitSymbol(quant.MeterPerSecond), quant.MeterPerSecond, "meter per second", "meters per second").
		With(quant.UnitSymbol(quant.KilometerPerHour), quant.KilometerPerHour, "kilometer per hour", "kilometers per hour").
		With(quant.UnitSymbol(quant.MilePerHour), quant.MilePerHour, "mile per hour", "miles per hour").
		With(quant.UnitSymbol(quant.FootPerSecond), quant.FootPerSecond, "foot per second", "feet per second").
		With(quant.UnitSymbol(quant.FootPerMinute), quant.FootPerMinute, "foot per minute", "feet per minute").
		With(quant.UnitSymbol(quant.InchPerHour), quant.InchPerHour, "inch per hour", "inches per hour")
}

// Accelerations returns an immutable registry for acceleration units.
func Accelerations() Registry[quant.Acceleration] {
	return NewRegistry[quant.Acceleration]().
		With(quant.UnitSymbol(quant.MeterPerSecondSquared), quant.MeterPerSecondSquared, "meter per second squared", "meters per second squared").
		With(quant.UnitSymbol(quant.GForce), quant.GForce, "g-force", "g forces")
}

// Forces returns an immutable registry for force units.
func Forces() Registry[quant.Force] {
	return NewRegistry[quant.Force]().
		With(quant.UnitSymbol(quant.Newton), quant.Newton, "newton", "newtons").
		With(quant.UnitSymbol(quant.Kilonewton), quant.Kilonewton, "kilonewton", "kilonewtons").
		With(quant.UnitSymbol(quant.PoundForce), quant.PoundForce, "pound-force", "pound-forces").
		With(quant.UnitSymbol(quant.KilogramForce), quant.KilogramForce, "kilogram-force", "kilogram-forces")
}

// Pressures returns an immutable registry for pressure units.
func Pressures() Registry[quant.Pressure] {
	return NewRegistry[quant.Pressure]().
		With(quant.UnitSymbol(quant.Pascal), quant.Pascal, "pascal", "pascals").
		With(quant.UnitSymbol(quant.Hectopascal), quant.Hectopascal, "hectopascal", "hectopascals").
		With(quant.UnitSymbol(quant.Kilopascal), quant.Kilopascal, "kilopascal", "kilopascals").
		With(quant.UnitSymbol(quant.Megapascal), quant.Megapascal, "megapascal", "megapascals").
		With(quant.UnitSymbol(quant.Bar), quant.Bar, "bars").
		With(quant.UnitSymbol(quant.PSI), quant.PSI)
}

// Energies returns an immutable registry for energy units.
func Energies() Registry[quant.Energy] {
	return NewRegistry[quant.Energy]().
		With(quant.UnitSymbol(quant.Joule), quant.Joule, "joule", "joules").
		With(quant.UnitSymbol(quant.Kilojoule), quant.Kilojoule, "kilojoule", "kilojoules").
		With(quant.UnitSymbol(quant.Megajoule), quant.Megajoule, "megajoule", "megajoules").
		With(quant.UnitSymbol(quant.KilowattHour), quant.KilowattHour, "kilowatt hour", "kilowatt hours")
}

// Powers returns an immutable registry for active power units.
func Powers() Registry[quant.Power] {
	return NewRegistry[quant.Power]().
		With(quant.UnitSymbol(quant.Watt), quant.Watt, "watt", "watts").
		With(quant.UnitSymbol(quant.Kilowatt), quant.Kilowatt, "kilowatt", "kilowatts").
		With(quant.UnitSymbol(quant.Megawatt), quant.Megawatt, "megawatt", "megawatts")
}

// Frequencies returns an immutable registry for frequency units.
func Frequencies() Registry[quant.Frequency] {
	return NewRegistry[quant.Frequency]().
		With(quant.UnitSymbol(quant.Hertz), quant.Hertz, "hertz").
		With(quant.UnitSymbol(quant.Kilohertz), quant.Kilohertz, "kilohertz").
		With(quant.UnitSymbol(quant.Megahertz), quant.Megahertz, "megahertz").
		With(quant.UnitSymbol(quant.RevolutionPerMinute), quant.RevolutionPerMinute, "revolutions per minute")
}

// Angles returns an immutable registry for angle units.
func Angles() Registry[quant.Angle] {
	return NewRegistry[quant.Angle]().
		With(quant.UnitSymbol(quant.Radian), quant.Radian, "radian", "radians").
		With(quant.UnitSymbol(quant.Degree), quant.Degree, "degree", "degrees", "°")
}

// Digitals returns an immutable registry for digital information units.
func Digitals() Registry[quant.Digital] {
	return NewRegistry[quant.Digital]().
		With(quant.UnitSymbol(quant.Bit), quant.Bit, "bits").
		With(quant.UnitSymbol(quant.Byte), quant.Byte, "bytes").
		With(quant.UnitSymbol(quant.Kibibyte), quant.Kibibyte, "kibibyte", "kibibytes").
		With(quant.UnitSymbol(quant.Kilobyte), quant.Kilobyte, "kilobyte", "kilobytes")
}

// DataRates returns an immutable registry for common data-rate units.
func DataRates() Registry[quant.DataRate] {
	return NewRegistry[quant.DataRate]().
		With(quant.UnitSymbol(quant.BitPerSecond), quant.BitPerSecond, "bit per second", "bits per second").
		With(quant.UnitSymbol(quant.KilobitPerSecond), quant.KilobitPerSecond, "kilobit per second").
		With(quant.UnitSymbol(quant.MegabitPerSecond), quant.MegabitPerSecond, "megabit per second").
		With(quant.UnitSymbol(quant.BytePerSecond), quant.BytePerSecond, "byte per second", "bytes per second")
}

// Voltages returns an immutable registry for voltage units.
func Voltages() Registry[quant.Voltage] {
	return NewRegistry[quant.Voltage]().
		With(quant.UnitSymbol(quant.Volt), quant.Volt, "volt", "volts").
		With(quant.UnitSymbol(quant.Kilovolt), quant.Kilovolt, "kilovolt", "kilovolts")
}

// Currents returns an immutable registry for current units.
func Currents() Registry[quant.Current] {
	return NewRegistry[quant.Current]().
		With(quant.UnitSymbol(quant.Ampere), quant.Ampere, "ampere", "amperes", "amp", "amps").
		With(quant.UnitSymbol(quant.Kiloampere), quant.Kiloampere, "kiloampere", "kiloamperes")
}

// Resistances returns an immutable registry for resistance units.
func Resistances() Registry[quant.Resistance] {
	return NewRegistry[quant.Resistance]().
		With(quant.UnitSymbol(quant.Ohm), quant.Ohm, "ohm", "ohms").
		With(quant.UnitSymbol(quant.Kiloohm), quant.Kiloohm, "kiloohm", "kiloohms").
		With(quant.UnitSymbol(quant.Megaohm), quant.Megaohm, "megaohm", "megaohms")
}

// Capacitances returns an immutable registry for capacitance units.
func Capacitances() Registry[quant.Capacitance] {
	return NewRegistry[quant.Capacitance]().
		With(quant.UnitSymbol(quant.Farad), quant.Farad, "farad", "farads").
		With(quant.UnitSymbol(quant.Microfarad), quant.Microfarad, "microfarad", "microfarads").
		With(quant.UnitSymbol(quant.Nanofarad), quant.Nanofarad, "nanofarad", "nanofarads")
}

// Charges returns an immutable registry for electric charge units.
func Charges() Registry[quant.Charge] {
	return NewRegistry[quant.Charge]().
		With(quant.UnitSymbol(quant.Coulomb), quant.Coulomb, "coulomb", "coulombs")
}

// TemperatureDeltas returns an immutable registry for temperature differences.
func TemperatureDeltas() Registry[quant.TemperatureDelta] {
	return NewRegistry[quant.TemperatureDelta]().
		With(quant.UnitSymbol(quant.KelvinDelta), quant.KelvinDelta, "kelvin delta", "kelvin deltas").
		With(quant.UnitSymbol(quant.CelsiusDelta), quant.CelsiusDelta, "celsius delta", "celsius deltas", "°C-delta").
		With(quant.UnitSymbol(quant.FahrenheitDelta), quant.FahrenheitDelta, "fahrenheit delta", "fahrenheit deltas", "°F-delta")
}

// ParseTemperatureDelta parses a temperature difference using the built-in
// temperature-delta registry.
func ParseTemperatureDelta(input string) (quant.Quantity[quant.TemperatureDelta], error) {
	return TemperatureDeltas().Parse(input)
}
