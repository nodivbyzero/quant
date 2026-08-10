package quant

func quantityFrom[D any, U Unit[D]](v float64, u U) Quantity[D] {
	return New[D](v, u)
}

// Mass constructors.
func Micrograms(v float64) Quantity[Mass] { return quantityFrom[Mass](v, Microgram) }
func Milligrams(v float64) Quantity[Mass] { return quantityFrom[Mass](v, Milligram) }
func Grams(v float64) Quantity[Mass]      { return quantityFrom[Mass](v, Gram) }
func Kilograms(v float64) Quantity[Mass]  { return quantityFrom[Mass](v, Kilogram) }
func Ounces(v float64) Quantity[Mass]     { return quantityFrom[Mass](v, Ounce) }
func Pounds(v float64) Quantity[Mass]     { return quantityFrom[Mass](v, Pound) }
func Stones(v float64) Quantity[Mass]     { return quantityFrom[Mass](v, Stone) }
func MetricTons(v float64) Quantity[Mass] { return quantityFrom[Mass](v, MetricTon) }
func Tonnes(v float64) Quantity[Mass]     { return quantityFrom[Mass](v, Tonne) }

// Length constructors.
func Nanometers(v float64) Quantity[Length]    { return quantityFrom[Length](v, Nanometer) }
func Micrometers(v float64) Quantity[Length]   { return quantityFrom[Length](v, Micrometer) }
func Millimeters(v float64) Quantity[Length]   { return quantityFrom[Length](v, Millimeter) }
func Centimeters(v float64) Quantity[Length]   { return quantityFrom[Length](v, Centimeter) }
func Meters(v float64) Quantity[Length]        { return quantityFrom[Length](v, Meter) }
func Inches(v float64) Quantity[Length]        { return quantityFrom[Length](v, Inch) }
func Yards(v float64) Quantity[Length]         { return quantityFrom[Length](v, Yard) }
func USFeet(v float64) Quantity[Length]        { return quantityFrom[Length](v, USFoot) }
func Feet(v float64) Quantity[Length]          { return quantityFrom[Length](v, Foot) }
func Fathoms(v float64) Quantity[Length]       { return quantityFrom[Length](v, Fathom) }
func Kilometers(v float64) Quantity[Length]    { return quantityFrom[Length](v, Kilometer) }
func Miles(v float64) Quantity[Length]         { return quantityFrom[Length](v, Mile) }
func NauticalMiles(v float64) Quantity[Length] { return quantityFrom[Length](v, NauticalMile) }

// Area constructors.
func SquareMillimeters(v float64) Quantity[Area] { return quantityFrom[Area](v, SquareMillimeter) }
func SquareCentimeters(v float64) Quantity[Area] { return quantityFrom[Area](v, SquareCentimeter) }
func SquareMeters(v float64) Quantity[Area]      { return quantityFrom[Area](v, SquareMeter) }
func Hectares(v float64) Quantity[Area]          { return quantityFrom[Area](v, Hectare) }
func SquareKilometers(v float64) Quantity[Area]  { return quantityFrom[Area](v, SquareKilometer) }
func SquareInches(v float64) Quantity[Area]      { return quantityFrom[Area](v, SquareInch) }
func SquareFeet(v float64) Quantity[Area]        { return quantityFrom[Area](v, SquareFoot) }
func Acres(v float64) Quantity[Area]             { return quantityFrom[Area](v, Acre) }
func SquareMiles(v float64) Quantity[Area]       { return quantityFrom[Area](v, SquareMile) }

// Acidity constructors.
func PHValue(v float64) Quantity[Acidity]  { return quantityFrom[Acidity](v, PH) }
func POHValue(v float64) Quantity[Acidity] { return quantityFrom[Acidity](v, POH) }

// Volume constructors.
func CubicMillimeters(v float64) Quantity[Volume] { return quantityFrom[Volume](v, CubicMillimeter) }
func CubicCentimeters(v float64) Quantity[Volume] { return quantityFrom[Volume](v, CubicCentimeter) }
func Milliliters(v float64) Quantity[Volume]      { return quantityFrom[Volume](v, Milliliter) }
func Liters(v float64) Quantity[Volume]           { return quantityFrom[Volume](v, Liter) }
func Kiloliters(v float64) Quantity[Volume]       { return quantityFrom[Volume](v, Kiloliter) }
func Megaliters(v float64) Quantity[Volume]       { return quantityFrom[Volume](v, Megaliter) }
func Gigaliters(v float64) Quantity[Volume]       { return quantityFrom[Volume](v, Gigaliter) }
func CubicMeters(v float64) Quantity[Volume]      { return quantityFrom[Volume](v, CubicMeter) }
func CubicKilometers(v float64) Quantity[Volume]  { return quantityFrom[Volume](v, CubicKilometer) }
func Teaspoons(v float64) Quantity[Volume]        { return quantityFrom[Volume](v, Teaspoon) }
func Tablespoons(v float64) Quantity[Volume]      { return quantityFrom[Volume](v, Tablespoon) }
func CubicInches(v float64) Quantity[Volume]      { return quantityFrom[Volume](v, CubicInch) }
func FluidOunces(v float64) Quantity[Volume]      { return quantityFrom[Volume](v, FluidOunce) }
func Cups(v float64) Quantity[Volume]             { return quantityFrom[Volume](v, Cup) }
func Pints(v float64) Quantity[Volume]            { return quantityFrom[Volume](v, Pint) }
func Quarts(v float64) Quantity[Volume]           { return quantityFrom[Volume](v, Quart) }
func Gallons(v float64) Quantity[Volume]          { return quantityFrom[Volume](v, Gallon) }
func CubicFeet(v float64) Quantity[Volume]        { return quantityFrom[Volume](v, CubicFoot) }
func CubicYards(v float64) Quantity[Volume]       { return quantityFrom[Volume](v, CubicYard) }

// Volume flow rate constructors.
func CubicMillimetersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicMillimeterPerSecond)
}
func CubicCentimetersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicCentimeterPerSecond)
}
func MillilitersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, MilliliterPerSecond)
}
func CentilitersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CentiliterPerSecond)
}
func DecilitersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, DeciliterPerSecond)
}
func LitersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, LiterPerSecond)
}
func LitersPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, LiterPerMinute)
}
func LitersPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, LiterPerHour)
}
func KilolitersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, KiloliterPerSecond)
}
func KilolitersPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, KiloliterPerMinute)
}
func KilolitersPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, KiloliterPerHour)
}
func CubicMetersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicMeterPerSecond)
}
func CubicMetersPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicMeterPerMinute)
}
func CubicMetersPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicMeterPerHour)
}
func CubicKilometersPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicKilometerPerSecond)
}
func TeaspoonsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, TeaspoonPerSecond)
}
func TablespoonsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, TablespoonPerSecond)
}
func CubicInchesPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicInchPerSecond)
}
func CubicInchesPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicInchPerMinute)
}
func CubicInchesPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicInchPerHour)
}
func FluidOuncesPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, FluidOuncePerSecond)
}
func FluidOuncesPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, FluidOuncePerMinute)
}
func FluidOuncesPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, FluidOuncePerHour)
}
func CupsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CupPerSecond)
}
func PintsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, PintPerSecond)
}
func PintsPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, PintPerMinute)
}
func PintsPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, PintPerHour)
}
func QuartsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, QuartPerSecond)
}
func GallonsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, GallonPerSecond)
}
func GallonsPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, GallonPerMinute)
}
func GallonsPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, GallonPerHour)
}
func CubicFeetPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicFootPerSecond)
}
func CubicFeetPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicFootPerMinute)
}
func CubicFeetPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicFootPerHour)
}
func CubicYardsPerSecond(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicYardPerSecond)
}
func CubicYardsPerMinute(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicYardPerMinute)
}
func CubicYardsPerHour(v float64) Quantity[VolumeFlowRate] {
	return quantityFrom[VolumeFlowRate](v, CubicYardPerHour)
}

// Temperature constructors.
func DegreesCelsius(v float64) Quantity[Temperature] { return quantityFrom[Temperature](v, Celsius) }
func DegreesFahrenheit(v float64) Quantity[Temperature] {
	return quantityFrom[Temperature](v, Fahrenheit)
}
func Kelvins(v float64) Quantity[Temperature]        { return quantityFrom[Temperature](v, Kelvin) }
func DegreesRankine(v float64) Quantity[Temperature] { return quantityFrom[Temperature](v, Rankine) }

// Temperature delta constructors.
func KelvinDeltas(v float64) Quantity[TemperatureDelta] {
	return quantityFrom[TemperatureDelta](v, KelvinDelta)
}
func CelsiusDeltas(v float64) Quantity[TemperatureDelta] {
	return quantityFrom[TemperatureDelta](v, CelsiusDelta)
}
func FahrenheitDeltas(v float64) Quantity[TemperatureDelta] {
	return quantityFrom[TemperatureDelta](v, FahrenheitDelta)
}
func RankineDeltas(v float64) Quantity[TemperatureDelta] {
	return quantityFrom[TemperatureDelta](v, RankineDelta)
}

// Time constructors.
func Nanoseconds(v float64) Quantity[Time]  { return quantityFrom[Time](v, Nanosecond) }
func Microseconds(v float64) Quantity[Time] { return quantityFrom[Time](v, Microsecond) }
func Milliseconds(v float64) Quantity[Time] { return quantityFrom[Time](v, Millisecond) }
func Seconds(v float64) Quantity[Time]      { return quantityFrom[Time](v, Second) }
func Minutes(v float64) Quantity[Time]      { return quantityFrom[Time](v, Minute) }
func Hours(v float64) Quantity[Time]        { return quantityFrom[Time](v, Hour) }
func Days(v float64) Quantity[Time]         { return quantityFrom[Time](v, Day) }
func Weeks(v float64) Quantity[Time]        { return quantityFrom[Time](v, Week) }
func Months(v float64) Quantity[Time]       { return quantityFrom[Time](v, Month) }
func Years(v float64) Quantity[Time]        { return quantityFrom[Time](v, Year) }
func Decades(v float64) Quantity[Time]      { return quantityFrom[Time](v, Decade) }
func Centuries(v float64) Quantity[Time]    { return quantityFrom[Time](v, Century) }

// Frequency constructors.
func Hertzes(v float64) Quantity[Frequency]      { return quantityFrom[Frequency](v, Hertz) }
func Millihertzes(v float64) Quantity[Frequency] { return quantityFrom[Frequency](v, Millihertz) }
func Kilohertzes(v float64) Quantity[Frequency]  { return quantityFrom[Frequency](v, Kilohertz) }
func Megahertzes(v float64) Quantity[Frequency]  { return quantityFrom[Frequency](v, Megahertz) }
func Gigahertzes(v float64) Quantity[Frequency]  { return quantityFrom[Frequency](v, Gigahertz) }
func Terahertzes(v float64) Quantity[Frequency]  { return quantityFrom[Frequency](v, Terahertz) }
func RevolutionsPerMinute(v float64) Quantity[Frequency] {
	return quantityFrom[Frequency](v, RevolutionPerMinute)
}
func DegreesPerSecond(v float64) Quantity[Frequency] {
	return quantityFrom[Frequency](v, DegreePerSecond)
}
func RadiansPerSecond(v float64) Quantity[Frequency] {
	return quantityFrom[Frequency](v, RadianPerSecond)
}

// Speed constructors.
func MetersPerSecond(v float64) Quantity[Speed]    { return quantityFrom[Speed](v, MeterPerSecond) }
func KilometersPerHour(v float64) Quantity[Speed]  { return quantityFrom[Speed](v, KilometerPerHour) }
func MilesPerHour(v float64) Quantity[Speed]       { return quantityFrom[Speed](v, MilePerHour) }
func MeterPerHours(v float64) Quantity[Speed]      { return quantityFrom[Speed](v, MeterPerHour) }
func Knots(v float64) Quantity[Speed]              { return quantityFrom[Speed](v, Knot) }
func FeetPerSecond(v float64) Quantity[Speed]      { return quantityFrom[Speed](v, FootPerSecond) }
func InchesPerHour(v float64) Quantity[Speed]      { return quantityFrom[Speed](v, InchPerHour) }
func MillimetersPerHour(v float64) Quantity[Speed] { return quantityFrom[Speed](v, MillimeterPerHour) }

// Torque constructors.
func NewtonMeters(v float64) Quantity[Torque]   { return quantityFrom[Torque](v, NewtonMeter) }
func PoundForceFeet(v float64) Quantity[Torque] { return quantityFrom[Torque](v, PoundForceFoot) }

// Pace constructors.
func SecondsPerMeter(v float64) Quantity[Pace]     { return quantityFrom[Pace](v, SecondPerMeter) }
func MinutesPerKilometer(v float64) Quantity[Pace] { return quantityFrom[Pace](v, MinutePerKilometer) }
func SecondsPerFoot(v float64) Quantity[Pace]      { return quantityFrom[Pace](v, SecondPerFoot) }
func MinutesPerMile(v float64) Quantity[Pace]      { return quantityFrom[Pace](v, MinutePerMile) }

// Pressure constructors.
func Pascals(v float64) Quantity[Pressure]       { return quantityFrom[Pressure](v, Pascal) }
func Hectopascals(v float64) Quantity[Pressure]  { return quantityFrom[Pressure](v, Hectopascal) }
func Kilopascals(v float64) Quantity[Pressure]   { return quantityFrom[Pressure](v, Kilopascal) }
func Megapascals(v float64) Quantity[Pressure]   { return quantityFrom[Pressure](v, Megapascal) }
func Bars(v float64) Quantity[Pressure]          { return quantityFrom[Pressure](v, Bar) }
func Torrs(v float64) Quantity[Pressure]         { return quantityFrom[Pressure](v, Torr) }
func MetersOfWater(v float64) Quantity[Pressure] { return quantityFrom[Pressure](v, MeterOfWater) }
func MillimetersOfMercury(v float64) Quantity[Pressure] {
	return quantityFrom[Pressure](v, MillimeterOfMercury)
}
func PSIs(v float64) Quantity[Pressure] { return quantityFrom[Pressure](v, PSI) }
func KSIs(v float64) Quantity[Pressure] { return quantityFrom[Pressure](v, KSI) }

// Digital constructors.
func Bits(v float64) Quantity[Digital]      { return quantityFrom[Digital](v, Bit) }
func Kilobits(v float64) Quantity[Digital]  { return quantityFrom[Digital](v, Kilobit) }
func Megabits(v float64) Quantity[Digital]  { return quantityFrom[Digital](v, Megabit) }
func Gigabits(v float64) Quantity[Digital]  { return quantityFrom[Digital](v, Gigabit) }
func Terabits(v float64) Quantity[Digital]  { return quantityFrom[Digital](v, Terabit) }
func Bytes(v float64) Quantity[Digital]     { return quantityFrom[Digital](v, Byte) }
func Kilobytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Kilobyte) }
func Megabytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Megabyte) }
func Gigabytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Gigabyte) }
func Terabytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Terabyte) }
func Petabytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Petabyte) }
func Kibibytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Kibibyte) }
func Mebibytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Mebibyte) }
func Gibibytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Gibibyte) }
func Tebibytes(v float64) Quantity[Digital] { return quantityFrom[Digital](v, Tebibyte) }

// Illuminance constructors.
func Luxes(v float64) Quantity[Illuminance]       { return quantityFrom[Illuminance](v, Lux) }
func FootCandles(v float64) Quantity[Illuminance] { return quantityFrom[Illuminance](v, FootCandle) }

// Parts-per constructors.
func PartsPerMillion(v float64) Quantity[PartsPer]     { return quantityFrom[PartsPer](v, PPM) }
func PartsPerBillion(v float64) Quantity[PartsPer]     { return quantityFrom[PartsPer](v, PPB) }
func PartsPerTrillion(v float64) Quantity[PartsPer]    { return quantityFrom[PartsPer](v, PPT) }
func PartsPerQuadrillion(v float64) Quantity[PartsPer] { return quantityFrom[PartsPer](v, PPQ) }

// Electrical constructors.
func Volts(v float64) Quantity[Voltage]        { return quantityFrom[Voltage](v, Volt) }
func Millivolts(v float64) Quantity[Voltage]   { return quantityFrom[Voltage](v, Millivolt) }
func Kilovolts(v float64) Quantity[Voltage]    { return quantityFrom[Voltage](v, Kilovolt) }
func Amperes(v float64) Quantity[Current]      { return quantityFrom[Current](v, Ampere) }
func Milliamperes(v float64) Quantity[Current] { return quantityFrom[Current](v, Milliampere) }
func Kiloamperes(v float64) Quantity[Current]  { return quantityFrom[Current](v, Kiloampere) }
func Watts(v float64) Quantity[Power]          { return quantityFrom[Power](v, Watt) }
func Milliwatts(v float64) Quantity[Power]     { return quantityFrom[Power](v, Milliwatt) }
func Kilowatts(v float64) Quantity[Power]      { return quantityFrom[Power](v, Kilowatt) }
func Megawatts(v float64) Quantity[Power]      { return quantityFrom[Power](v, Megawatt) }
func Gigawatts(v float64) Quantity[Power]      { return quantityFrom[Power](v, Gigawatt) }
func MetricHorsepowerValues(v float64) Quantity[Power] {
	return quantityFrom[Power](v, MetricHorsepower)
}
func BTUsPerSecond(v float64) Quantity[Power] { return quantityFrom[Power](v, BTUPerSecond) }
func FootPoundForcePerSeconds(v float64) Quantity[Power] {
	return quantityFrom[Power](v, FootPoundForcePerSecond)
}
func HorsepowerValues(v float64) Quantity[Power] { return quantityFrom[Power](v, Horsepower) }
func VoltAmperes(v float64) Quantity[ApparentPower] {
	return quantityFrom[ApparentPower](v, VoltAmpere)
}
func MilliVoltAmperes(v float64) Quantity[ApparentPower] {
	return quantityFrom[ApparentPower](v, MilliVoltAmpere)
}
func KiloVoltAmperes(v float64) Quantity[ApparentPower] {
	return quantityFrom[ApparentPower](v, KiloVoltAmpere)
}
func MegaVoltAmperes(v float64) Quantity[ApparentPower] {
	return quantityFrom[ApparentPower](v, MegaVoltAmpere)
}
func GigaVoltAmperes(v float64) Quantity[ApparentPower] {
	return quantityFrom[ApparentPower](v, GigaVoltAmpere)
}
func VoltAmpereReactives(v float64) Quantity[ReactivePower] {
	return quantityFrom[ReactivePower](v, VoltAmpereReactive)
}
func MilliVoltAmpereReactives(v float64) Quantity[ReactivePower] {
	return quantityFrom[ReactivePower](v, MilliVoltAmpereReactive)
}
func KiloVoltAmpereReactives(v float64) Quantity[ReactivePower] {
	return quantityFrom[ReactivePower](v, KiloVoltAmpereReactive)
}
func MegaVoltAmpereReactives(v float64) Quantity[ReactivePower] {
	return quantityFrom[ReactivePower](v, MegaVoltAmpereReactive)
}
func GigaVoltAmpereReactives(v float64) Quantity[ReactivePower] {
	return quantityFrom[ReactivePower](v, GigaVoltAmpereReactive)
}

// Resistance constructors.
func Ohms(v float64) Quantity[Resistance]      { return quantityFrom[Resistance](v, Ohm) }
func Milliohms(v float64) Quantity[Resistance] { return quantityFrom[Resistance](v, Milliohm) }
func Kiloohms(v float64) Quantity[Resistance]  { return quantityFrom[Resistance](v, Kiloohm) }
func Megaohms(v float64) Quantity[Resistance]  { return quantityFrom[Resistance](v, Megaohm) }

// Capacitance constructors.
func Farads(v float64) Quantity[Capacitance]      { return quantityFrom[Capacitance](v, Farad) }
func Microfarads(v float64) Quantity[Capacitance] { return quantityFrom[Capacitance](v, Microfarad) }
func Nanofarads(v float64) Quantity[Capacitance]  { return quantityFrom[Capacitance](v, Nanofarad) }
func Picofarads(v float64) Quantity[Capacitance]  { return quantityFrom[Capacitance](v, Picofarad) }

// Inductance constructors.
func Henrys(v float64) Quantity[Inductance]      { return quantityFrom[Inductance](v, Henry) }
func Millihenrys(v float64) Quantity[Inductance] { return quantityFrom[Inductance](v, Millihenry) }
func Microhenrys(v float64) Quantity[Inductance] { return quantityFrom[Inductance](v, Microhenry) }

// Electric field constructors.
func VoltsPerMeter(v float64) Quantity[ElectricField] {
	return quantityFrom[ElectricField](v, VoltPerMeter)
}
func NewtonsPerCoulomb(v float64) Quantity[ElectricField] {
	return quantityFrom[ElectricField](v, NewtonPerCoulomb)
}

// Electrical conductivity constructors.
func SiemensPerMeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, SiemensPerMeter)
}
func MillisiemensPerMeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, MillisiemensPerMeter)
}
func MicrosiemensPerMeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, MicrosiemensPerMeter)
}
func SiemensPerCentimeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, SiemensPerCentimeter)
}
func MillisiemensPerCentimeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, MillisiemensPerCentimeter)
}
func MicrosiemensPerCentimeters(v float64) Quantity[ElectricalConductivity] {
	return quantityFrom[ElectricalConductivity](v, MicrosiemensPerCentimeter)
}

// Energy constructors.
func WattSeconds(v float64) Quantity[Energy]    { return quantityFrom[Energy](v, WattSecond) }
func WattMinutes(v float64) Quantity[Energy]    { return quantityFrom[Energy](v, WattMinute) }
func MilliwattHours(v float64) Quantity[Energy] { return quantityFrom[Energy](v, MilliwattHour) }
func WattHours(v float64) Quantity[Energy]      { return quantityFrom[Energy](v, WattHour) }
func KilowattHours(v float64) Quantity[Energy]  { return quantityFrom[Energy](v, KilowattHour) }
func MegawattHours(v float64) Quantity[Energy]  { return quantityFrom[Energy](v, MegawattHour) }
func GigawattHours(v float64) Quantity[Energy]  { return quantityFrom[Energy](v, GigawattHour) }
func Joules(v float64) Quantity[Energy]         { return quantityFrom[Energy](v, Joule) }
func Kilojoules(v float64) Quantity[Energy]     { return quantityFrom[Energy](v, Kilojoule) }
func Megajoules(v float64) Quantity[Energy]     { return quantityFrom[Energy](v, Megajoule) }
func Gigajoules(v float64) Quantity[Energy]     { return quantityFrom[Energy](v, Gigajoule) }
func VoltAmpereReactiveHours(v float64) Quantity[ReactiveEnergy] {
	return quantityFrom[ReactiveEnergy](v, VoltAmpereReactiveHour)
}
func MilliVoltAmpereReactiveHours(v float64) Quantity[ReactiveEnergy] {
	return quantityFrom[ReactiveEnergy](v, MilliVoltAmpereReactiveHour)
}
func KiloVoltAmpereReactiveHours(v float64) Quantity[ReactiveEnergy] {
	return quantityFrom[ReactiveEnergy](v, KiloVoltAmpereReactiveHour)
}
func MegaVoltAmpereReactiveHours(v float64) Quantity[ReactiveEnergy] {
	return quantityFrom[ReactiveEnergy](v, MegaVoltAmpereReactiveHour)
}
func GigaVoltAmpereReactiveHours(v float64) Quantity[ReactiveEnergy] {
	return quantityFrom[ReactiveEnergy](v, GigaVoltAmpereReactiveHour)
}

// Angle constructors.
func Degrees(v float64) Quantity[Angle]    { return quantityFrom[Angle](v, Degree) }
func Radians(v float64) Quantity[Angle]    { return quantityFrom[Angle](v, Radian) }
func Gradians(v float64) Quantity[Angle]   { return quantityFrom[Angle](v, Gradian) }
func ArcMinutes(v float64) Quantity[Angle] { return quantityFrom[Angle](v, ArcMinute) }
func ArcSeconds(v float64) Quantity[Angle] { return quantityFrom[Angle](v, ArcSecond) }

// Charge constructors.
func Coulombs(v float64) Quantity[Charge]      { return quantityFrom[Charge](v, Coulomb) }
func Millicoulombs(v float64) Quantity[Charge] { return quantityFrom[Charge](v, Millicoulomb) }
func Microcoulombs(v float64) Quantity[Charge] { return quantityFrom[Charge](v, Microcoulomb) }
func Nanocoulombs(v float64) Quantity[Charge]  { return quantityFrom[Charge](v, Nanocoulomb) }
func Picocoulombs(v float64) Quantity[Charge]  { return quantityFrom[Charge](v, Picocoulomb) }

// Force constructors.
func Newtons(v float64) Quantity[Force]        { return quantityFrom[Force](v, Newton) }
func Kilonewtons(v float64) Quantity[Force]    { return quantityFrom[Force](v, Kilonewton) }
func PoundsForce(v float64) Quantity[Force]    { return quantityFrom[Force](v, PoundForce) }
func KilogramsForce(v float64) Quantity[Force] { return quantityFrom[Force](v, KilogramForce) }

// Acceleration constructors.
func MetersPerSecondSquared(v float64) Quantity[Acceleration] {
	return quantityFrom[Acceleration](v, MeterPerSecondSquared)
}
func GForces(v float64) Quantity[Acceleration] { return quantityFrom[Acceleration](v, GForce) }
func StandardGravities(v float64) Quantity[Acceleration] {
	return quantityFrom[Acceleration](v, StandardGravity)
}

// Angular velocity constructors.
func AngularRadiansPerSecond(v float64) Quantity[AngularVelocity] {
	return quantityFrom[AngularVelocity](v, AngularRadianPerSecond)
}
func AngularDegreesPerSecond(v float64) Quantity[AngularVelocity] {
	return quantityFrom[AngularVelocity](v, AngularDegreePerSecond)
}
func AngularRevolutionsPerMinute(v float64) Quantity[AngularVelocity] {
	return quantityFrom[AngularVelocity](v, AngularRevolutionPerMinute)
}

// Angular acceleration constructors.
func AngularRadiansPerSecondSquared(v float64) Quantity[AngularAcceleration] {
	return quantityFrom[AngularAcceleration](v, AngularRadianPerSecondSquared)
}
func AngularDegreesPerSecondSquared(v float64) Quantity[AngularAcceleration] {
	return quantityFrom[AngularAcceleration](v, AngularDegreePerSecondSquared)
}

// Density constructors.
func KilogramsPerCubicMeter(v float64) Quantity[Density] {
	return quantityFrom[Density](v, KilogramPerCubicMeter)
}
func GramsPerCubicCentimeter(v float64) Quantity[Density] {
	return quantityFrom[Density](v, GramPerCubicCentimeter)
}
func PoundsPerCubicFoot(v float64) Quantity[Density] {
	return quantityFrom[Density](v, PoundPerCubicFoot)
}
func KilogramsPerLiter(v float64) Quantity[Density] {
	return quantityFrom[Density](v, KilogramPerLiter)
}

// Data rate constructors.
func BitsPerSecond(v float64) Quantity[DataRate] { return quantityFrom[DataRate](v, BitPerSecond) }
func KilobitsPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, KilobitPerSecond)
}
func MegabitsPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, MegabitPerSecond)
}
func GigabitsPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, GigabitPerSecond)
}
func TerabitsPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, TerabitPerSecond)
}
func BytesPerSecond(v float64) Quantity[DataRate] { return quantityFrom[DataRate](v, BytePerSecond) }
func KilobytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, KilobytePerSecond)
}
func MegabytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, MegabytePerSecond)
}
func GigabytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, GigabytePerSecond)
}
func TerabytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, TerabytePerSecond)
}
func KibibytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, KibibytePerSecond)
}
func MebibytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, MebibytePerSecond)
}
func GibibytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, GibibytePerSecond)
}
func TebibytesPerSecond(v float64) Quantity[DataRate] {
	return quantityFrom[DataRate](v, TebibytePerSecond)
}

// Luminous flux constructors.
func Lumens(v float64) Quantity[LuminousFlux] { return quantityFrom[LuminousFlux](v, Lumen) }

// Luminous intensity constructors.
func Candelas(v float64) Quantity[LuminousIntensity] {
	return quantityFrom[LuminousIntensity](v, Candela)
}

// Magnetic flux constructors.
func Webers(v float64) Quantity[MagneticFlux]   { return quantityFrom[MagneticFlux](v, Weber) }
func Maxwells(v float64) Quantity[MagneticFlux] { return quantityFrom[MagneticFlux](v, Maxwell) }

// Magnetic flux density constructors.
func Teslas(v float64) Quantity[MagneticFluxDensity] {
	return quantityFrom[MagneticFluxDensity](v, Tesla)
}
func Milliteslas(v float64) Quantity[MagneticFluxDensity] {
	return quantityFrom[MagneticFluxDensity](v, Millitesla)
}
func Gausses(v float64) Quantity[MagneticFluxDensity] {
	return quantityFrom[MagneticFluxDensity](v, Gauss)
}

// Dynamic viscosity constructors.
func PascalSeconds(v float64) Quantity[DynamicViscosity] {
	return quantityFrom[DynamicViscosity](v, PascalSecond)
}
func Poises(v float64) Quantity[DynamicViscosity] {
	return quantityFrom[DynamicViscosity](v, Poise)
}
func Centipoises(v float64) Quantity[DynamicViscosity] {
	return quantityFrom[DynamicViscosity](v, Centipoise)
}

// Kinematic viscosity constructors.
func SquareMetersPerSecond(v float64) Quantity[KinematicViscosity] {
	return quantityFrom[KinematicViscosity](v, SquareMeterPerSecond)
}
func StokesValues(v float64) Quantity[KinematicViscosity] {
	return quantityFrom[KinematicViscosity](v, Stokes)
}
func CentistokesValues(v float64) Quantity[KinematicViscosity] {
	return quantityFrom[KinematicViscosity](v, Centistokes)
}

// Concentration constructors.
func MolesPerCubicMeter(v float64) Quantity[Concentration] {
	return quantityFrom[Concentration](v, MolePerCubicMeter)
}
func Molars(v float64) Quantity[Concentration] { return quantityFrom[Concentration](v, Molar) }
func Millimolars(v float64) Quantity[Concentration] {
	return quantityFrom[Concentration](v, Millimolar)
}

// Molality constructors.
func MolesPerKilogram(v float64) Quantity[Molality] {
	return quantityFrom[Molality](v, MolePerKilogram)
}
func Molals(v float64) Quantity[Molality] { return quantityFrom[Molality](v, Molal) }

// Surface tension constructors.
func NewtonsPerMeter(v float64) Quantity[SurfaceTension] {
	return quantityFrom[SurfaceTension](v, NewtonPerMeter)
}
func DynesPerCentimeter(v float64) Quantity[SurfaceTension] {
	return quantityFrom[SurfaceTension](v, DynePerCentimeter)
}

// Thermal conductivity constructors.
func WattsPerMeterKelvin(v float64) Quantity[ThermalConductivity] {
	return quantityFrom[ThermalConductivity](v, WattPerMeterKelvin)
}

// Specific heat capacity constructors.
func JoulesPerKilogramKelvin(v float64) Quantity[SpecificHeatCapacity] {
	return quantityFrom[SpecificHeatCapacity](v, JoulePerKilogramKelvin)
}

// Momentum constructors.
func KilogramMetersPerSecond(v float64) Quantity[Momentum] {
	return quantityFrom[Momentum](v, KilogramMeterPerSecond)
}
func NewtonSeconds(v float64) Quantity[Momentum] {
	return quantityFrom[Momentum](v, NewtonSecond)
}

// Radioactivity constructors.
func Becquerels(v float64) Quantity[Radioactivity] { return quantityFrom[Radioactivity](v, Becquerel) }
func Curies(v float64) Quantity[Radioactivity]     { return quantityFrom[Radioactivity](v, Curie) }

// Absorbed dose constructors.
func Grays(v float64) Quantity[AbsorbedDose] { return quantityFrom[AbsorbedDose](v, Gray) }
func Rads(v float64) Quantity[AbsorbedDose]  { return quantityFrom[AbsorbedDose](v, Rad) }

// Equivalent dose constructors.
func Sieverts(v float64) Quantity[EquivalentDose] { return quantityFrom[EquivalentDose](v, Sievert) }
func Rems(v float64) Quantity[EquivalentDose]     { return quantityFrom[EquivalentDose](v, Rem) }

// Catalytic activity constructors.
func Katals(v float64) Quantity[CatalyticActivity] {
	return quantityFrom[CatalyticActivity](v, Katal)
}

// Pieces constructors.
func PiecesCount(v float64) Quantity[Pieces]  { return quantityFrom[Pieces](v, Piece) }
func BakersDozens(v float64) Quantity[Pieces] { return quantityFrom[Pieces](v, BakersDozen) }
func Couples(v float64) Quantity[Pieces]      { return quantityFrom[Pieces](v, Couple) }
func DozenDozens(v float64) Quantity[Pieces]  { return quantityFrom[Pieces](v, DozenDozen) }
func Dozens(v float64) Quantity[Pieces]       { return quantityFrom[Pieces](v, Dozen) }
func GreatGrosses(v float64) Quantity[Pieces] { return quantityFrom[Pieces](v, GreatGross) }
func Grosses(v float64) Quantity[Pieces]      { return quantityFrom[Pieces](v, Gross) }
func HalfDozens(v float64) Quantity[Pieces]   { return quantityFrom[Pieces](v, HalfDozen) }
func LongHundreds(v float64) Quantity[Pieces] { return quantityFrom[Pieces](v, LongHundred) }
func Reams(v float64) Quantity[Pieces]        { return quantityFrom[Pieces](v, Ream) }
func Scores(v float64) Quantity[Pieces]       { return quantityFrom[Pieces](v, Score) }
func SmallGrosses(v float64) Quantity[Pieces] { return quantityFrom[Pieces](v, SmallGross) }
func Trios(v float64) Quantity[Pieces]        { return quantityFrom[Pieces](v, Trio) }
