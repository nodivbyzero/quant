package quant

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// Quantity stores a physical quantity in the base unit for its dimension.
//
// The dimension type parameter is a phantom type used only for compile-time
// safety.
type Quantity[D any] struct {
	value float64
}

// New constructs a quantity from a value expressed in the provided unit.
func New[D any, U Unit[D]](v float64, u U) Quantity[D] {
	return Quantity[D]{value: u.toBase(v)}
}

// From constructs a quantity from a value expressed in a built-in unit.
func From[D any](v float64, u BuiltinUnit[D]) Quantity[D] {
	return Quantity[D]{value: u.toBase(v)}
}

// To converts the quantity to the provided unit and returns the scalar value.
func (q Quantity[D]) To(u Unit[D]) float64 {
	return u.fromBase(q.value)
}

// Value converts the quantity to the provided unit and returns the scalar value.
func (q Quantity[D]) Value(u Unit[D]) float64 {
	return q.To(u)
}

// MarshalJSON encodes the quantity as its scalar value in the base unit for
// the dimension.
func (q Quantity[D]) MarshalJSON() ([]byte, error) {
	return json.Marshal(q.value)
}

// UnmarshalJSON decodes a quantity from a scalar value in the base unit for
// the dimension.
func (q *Quantity[D]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &q.value)
}

// MarshalText encodes the quantity as a base-unit scalar value.
func (q Quantity[D]) MarshalText() ([]byte, error) {
	return []byte(strconv.FormatFloat(q.value, 'g', 12, 64)), nil
}

// UnmarshalText decodes a base-unit scalar value into the quantity.
func (q *Quantity[D]) UnmarshalText(text []byte) error {
	v, err := strconv.ParseFloat(string(text), 64)
	if err != nil {
		return err
	}

	q.value = v
	return nil
}

// Format formats the quantity in the provided unit.
func (q Quantity[D]) Format(u Unit[D]) string {
	value := strconv.FormatFloat(q.To(u), 'g', 12, 64)
	symbol := unitSymbol(u)
	if symbol == "" {
		return value
	}

	return value + " " + symbol
}

// SQLQuantity wraps a quantity for database/sql interop using base-unit scalar values.
type SQLQuantity[D any] struct {
	Quantity[D]
}

// SQL wraps a quantity in a database-friendly adapter.
func SQL[D any](q Quantity[D]) SQLQuantity[D] {
	return SQLQuantity[D]{Quantity: q}
}

// Value implements database/sql/driver.Valuer using the base-unit scalar value.
func (q SQLQuantity[D]) Value() (driver.Value, error) {
	return q.Quantity.value, nil
}

// Scan implements database/sql.Scanner using a base-unit scalar value.
func (q *SQLQuantity[D]) Scan(src any) error {
	v, err := scanFloat64(src)
	if err != nil {
		return err
	}

	q.Quantity.value = v
	return nil
}

// String formats the quantity in the base unit for its dimension.
func (q Quantity[D]) String() string {
	u := baseUnit[D]()
	if u == nil {
		return strconv.FormatFloat(q.value, 'g', 12, 64)
	}

	return q.Format(u)
}

// Add adds two quantities of the same dimension.
func (q Quantity[D]) Add(other Quantity[D]) Quantity[D] {
	return Quantity[D]{value: q.value + other.value}
}

// Sub subtracts another quantity of the same dimension.
func (q Quantity[D]) Sub(other Quantity[D]) Quantity[D] {
	return Quantity[D]{value: q.value - other.value}
}

// Mul multiplies the quantity by a scalar.
func (q Quantity[D]) Mul(scalar float64) Quantity[D] {
	return Quantity[D]{value: q.value * scalar}
}

// DivScalar divides the quantity by a scalar.
func (q Quantity[D]) DivScalar(scalar float64) Quantity[D] {
	return Quantity[D]{value: q.value / scalar}
}

// Neg returns the additive inverse of the quantity.
func (q Quantity[D]) Neg() Quantity[D] {
	return Quantity[D]{value: -q.value}
}

// Abs returns the absolute value of the quantity.
func (q Quantity[D]) Abs() Quantity[D] {
	return Quantity[D]{value: math.Abs(q.value)}
}

// LessThan reports whether q is less than other.
func (q Quantity[D]) LessThan(other Quantity[D]) bool {
	return q.value < other.value
}

// GreaterThan reports whether q is greater than other.
func (q Quantity[D]) GreaterThan(other Quantity[D]) bool {
	return q.value > other.value
}

// EqualWithin reports whether q and other differ by no more than tolerance.
func (q Quantity[D]) EqualWithin(other, tolerance Quantity[D]) bool {
	return math.Abs(q.value-other.value) <= math.Abs(tolerance.value)
}

// IsZero reports whether the quantity's base-unit value is exactly zero.
func (q Quantity[D]) IsZero() bool {
	return q.value == 0
}

// IsPositive reports whether the quantity's base-unit value is positive.
func (q Quantity[D]) IsPositive() bool {
	return q.value > 0
}

// IsNegative reports whether the quantity's base-unit value is negative.
func (q Quantity[D]) IsNegative() bool {
	return q.value < 0
}

// Min returns the smaller of two quantities with the same dimension.
func Min[D any](a, b Quantity[D]) Quantity[D] {
	if a.LessThan(b) {
		return a
	}

	return b
}

// Max returns the larger of two quantities with the same dimension.
func Max[D any](a, b Quantity[D]) Quantity[D] {
	if a.GreaterThan(b) {
		return a
	}

	return b
}

func baseUnitSymbol[D any]() string {
	return unitSymbol(baseUnit[D]())
}

func baseUnit[D any]() Unit[D] {
	switch any(*new(D)).(type) {
	case Mass:
		return any(Kilogram).(Unit[D])
	case Length:
		return any(Meter).(Unit[D])
	case Area:
		return any(SquareMeter).(Unit[D])
	case Acidity:
		return any(PH).(Unit[D])
	case Volume:
		return any(CubicMeter).(Unit[D])
	case VolumeFlowRate:
		return any(CubicMeterPerSecond).(Unit[D])
	case Temperature:
		return any(Kelvin).(Unit[D])
	case Time:
		return any(Second).(Unit[D])
	case Frequency:
		return any(Hertz).(Unit[D])
	case Speed:
		return any(MeterPerSecond).(Unit[D])
	case Torque:
		return any(NewtonMeter).(Unit[D])
	case Pace:
		return any(SecondPerMeter).(Unit[D])
	case Pressure:
		return any(Pascal).(Unit[D])
	case Digital:
		return any(Bit).(Unit[D])
	case Illuminance:
		return any(Lux).(Unit[D])
	case PartsPer:
		return any(Ratio).(Unit[D])
	case Voltage:
		return any(Volt).(Unit[D])
	case Current:
		return any(Ampere).(Unit[D])
	case Power:
		return any(Watt).(Unit[D])
	case ApparentPower:
		return any(VoltAmpere).(Unit[D])
	case AngularAcceleration:
		return any(AngularRadianPerSecondSquared).(Unit[D])
	case AngularVelocity:
		return any(AngularRadianPerSecond).(Unit[D])
	case Capacitance:
		return any(Farad).(Unit[D])
	case CatalyticActivity:
		return any(Katal).(Unit[D])
	case ReactivePower:
		return any(VoltAmpereReactive).(Unit[D])
	case Concentration:
		return any(MolePerCubicMeter).(Unit[D])
	case DataRate:
		return any(BitPerSecond).(Unit[D])
	case Density:
		return any(KilogramPerCubicMeter).(Unit[D])
	case DynamicViscosity:
		return any(PascalSecond).(Unit[D])
	case ElectricalConductivity:
		return any(SiemensPerMeter).(Unit[D])
	case ElectricField:
		return any(VoltPerMeter).(Unit[D])
	case Energy:
		return any(Joule).(Unit[D])
	case AbsorbedDose:
		return any(Gray).(Unit[D])
	case ReactiveEnergy:
		return any(VoltAmpereReactiveHour).(Unit[D])
	case EquivalentDose:
		return any(Sievert).(Unit[D])
	case Angle:
		return any(Radian).(Unit[D])
	case Charge:
		return any(Coulomb).(Unit[D])
	case Force:
		return any(Newton).(Unit[D])
	case Acceleration:
		return any(MeterPerSecondSquared).(Unit[D])
	case Inductance:
		return any(Henry).(Unit[D])
	case KinematicViscosity:
		return any(SquareMeterPerSecond).(Unit[D])
	case LuminousFlux:
		return any(Lumen).(Unit[D])
	case LuminousIntensity:
		return any(Candela).(Unit[D])
	case MagneticFlux:
		return any(Weber).(Unit[D])
	case MagneticFluxDensity:
		return any(Tesla).(Unit[D])
	case Molality:
		return any(MolePerKilogram).(Unit[D])
	case Momentum:
		return any(KilogramMeterPerSecond).(Unit[D])
	case Pieces:
		return any(Piece).(Unit[D])
	case Radioactivity:
		return any(Becquerel).(Unit[D])
	case Resistance:
		return any(Ohm).(Unit[D])
	case SpecificHeatCapacity:
		return any(JoulePerKilogramKelvin).(Unit[D])
	case SurfaceTension:
		return any(NewtonPerMeter).(Unit[D])
	case TemperatureDelta:
		return any(KelvinDelta).(Unit[D])
	case ThermalConductivity:
		return any(WattPerMeterKelvin).(Unit[D])
	default:
		var zero Unit[D]
		return zero
	}
}

func unitSymbol(u any) string {
	switch any(u) {
	case any(Microgram):
		return "mcg"
	case any(Milligram):
		return "mg"
	case any(Gram):
		return "g"
	case any(Kilogram):
		return "kg"
	case any(Ounce):
		return "oz"
	case any(Pound):
		return "lb"
	case any(Stone):
		return "st"
	case any(MetricTon):
		return "mt"
	case any(Tonne):
		return "t"
	case any(PH):
		return "pH"
	case any(POH):
		return "pOH"
	case any(Nanometer):
		return "nm"
	case any(Micrometer):
		return "um"
	case any(Millimeter):
		return "mm"
	case any(Centimeter):
		return "cm"
	case any(Meter):
		return "m"
	case any(Inch):
		return "in"
	case any(Yard):
		return "yd"
	case any(USFoot):
		return "ft-us"
	case any(Foot):
		return "ft"
	case any(Fathom):
		return "fathom"
	case any(Kilometer):
		return "km"
	case any(Mile):
		return "mi"
	case any(NauticalMile):
		return "nMi"
	case any(Celsius):
		return "C"
	case any(Fahrenheit):
		return "F"
	case any(Kelvin):
		return "K"
	case any(Rankine):
		return "R"
	case any(KelvinDelta):
		return "K-delta"
	case any(CelsiusDelta):
		return "C-delta"
	case any(FahrenheitDelta):
		return "F-delta"
	case any(RankineDelta):
		return "R-delta"
	case any(Second):
		return "s"
	case any(Minute):
		return "min"
	case any(Hour):
		return "h"
	case any(Day):
		return "d"
	case any(Week):
		return "week"
	case any(Month):
		return "month"
	case any(Year):
		return "year"
	case any(Decade):
		return "decade"
	case any(Century):
		return "century"
	case any(Liter):
		return "l"
	case any(LiterPerSecond):
		return "l/s"
	case any(LiterPerMinute):
		return "l/min"
	case any(KilometerPerHour):
		return "km/h"
	case any(MilePerHour):
		return "mph"
	case any(MeterPerSecond):
		return "m/s"
	case any(NewtonMeter):
		return "Nm"
	case any(PoundForceFoot):
		return "lbf-ft"
	case any(KilogramPerCubicMeter):
		return "kg/m3"
	case any(GramPerCubicCentimeter):
		return "g/cm3"
	case any(PoundPerCubicFoot):
		return "lb/ft3"
	case any(KilogramPerLiter):
		return "kg/l"
	case any(Pascal):
		return "Pa"
	case any(Bar):
		return "bar"
	case any(PSI):
		return "psi"
	case any(Ohm):
		return "Ohm"
	case any(Milliohm):
		return "mOhm"
	case any(Kiloohm):
		return "kOhm"
	case any(Megaohm):
		return "MOhm"
	case any(Farad):
		return "F"
	case any(Microfarad):
		return "uF"
	case any(Nanofarad):
		return "nF"
	case any(Picofarad):
		return "pF"
	case any(Henry):
		return "H"
	case any(Millihenry):
		return "mH"
	case any(Microhenry):
		return "uH"
	case any(Bit):
		return "bit"
	case any(Byte):
		return "byte"
	case any(Petabyte):
		return "PB"
	case any(Kibibyte):
		return "KiB"
	case any(BitPerSecond):
		return "bit/s"
	case any(KilobitPerSecond):
		return "kbps"
	case any(MegabitPerSecond):
		return "Mbps"
	case any(GigabitPerSecond):
		return "Gbps"
	case any(TerabitPerSecond):
		return "Tbps"
	case any(BytePerSecond):
		return "B/s"
	case any(KilobytePerSecond):
		return "kB/s"
	case any(MegabytePerSecond):
		return "MB/s"
	case any(GigabytePerSecond):
		return "GB/s"
	case any(TerabytePerSecond):
		return "TB/s"
	case any(KibibytePerSecond):
		return "KiB/s"
	case any(MebibytePerSecond):
		return "MiB/s"
	case any(GibibytePerSecond):
		return "GiB/s"
	case any(TebibytePerSecond):
		return "TiB/s"
	case any(Watt):
		return "W"
	case any(Kilowatt):
		return "kW"
	case any(VoltPerMeter):
		return "V/m"
	case any(NewtonPerCoulomb):
		return "N/C"
	case any(SiemensPerMeter):
		return "S/m"
	case any(MillisiemensPerMeter):
		return "mS/m"
	case any(MicrosiemensPerMeter):
		return "uS/m"
	case any(SiemensPerCentimeter):
		return "S/cm"
	case any(MillisiemensPerCentimeter):
		return "mS/cm"
	case any(MicrosiemensPerCentimeter):
		return "uS/cm"
	case any(Joule):
		return "J"
	case any(Lumen):
		return "lm"
	case any(Candela):
		return "cd"
	case any(Weber):
		return "Wb"
	case any(Maxwell):
		return "Mx"
	case any(Tesla):
		return "T"
	case any(Millitesla):
		return "mT"
	case any(Gauss):
		return "G"
	case any(Radian):
		return "rad"
	case any(Degree):
		return "deg"
	case any(AngularRadianPerSecond):
		return "rad/s"
	case any(AngularDegreePerSecond):
		return "deg/s"
	case any(AngularRevolutionPerMinute):
		return "rpm"
	case any(AngularRadianPerSecondSquared):
		return "rad/s2"
	case any(AngularDegreePerSecondSquared):
		return "deg/s2"
	case any(Coulomb):
		return "C"
	case any(Newton):
		return "N"
	case any(MeterPerSecondSquared):
		return "m/s2"
	case any(GForce):
		return "g"
	case any(PascalSecond):
		return "Pa*s"
	case any(Poise):
		return "P"
	case any(Centipoise):
		return "cP"
	case any(SquareMeterPerSecond):
		return "m2/s"
	case any(Stokes):
		return "St"
	case any(Centistokes):
		return "cSt"
	case any(MolePerCubicMeter):
		return "mol/m3"
	case any(Molar):
		return "M"
	case any(Millimolar):
		return "mM"
	case any(MolePerKilogram):
		return "mol/kg"
	case any(Molal):
		return "molal"
	case any(NewtonPerMeter):
		return "N/m"
	case any(DynePerCentimeter):
		return "dyn/cm"
	case any(WattPerMeterKelvin):
		return "W/(m*K)"
	case any(JoulePerKilogramKelvin):
		return "J/(kg*K)"
	case any(KilogramMeterPerSecond):
		return "kg*m/s"
	case any(NewtonSecond):
		return "N*s"
	case any(Becquerel):
		return "Bq"
	case any(Curie):
		return "Ci"
	case any(Gray):
		return "Gy"
	case any(Rad):
		return "rad"
	case any(Sievert):
		return "Sv"
	case any(Rem):
		return "rem"
	case any(Katal):
		return "kat"
	case any(Piece):
		return "pcs"
	case any(Ratio):
		return "ratio"
	default:
		return ""
	}
}

func scanFloat64(src any) (float64, error) {
	switch v := src.(type) {
	case nil:
		return 0, nil
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int:
		return float64(v), nil
	case []byte:
		return strconv.ParseFloat(string(v), 64)
	case string:
		return strconv.ParseFloat(v, 64)
	default:
		return 0, fmt.Errorf("quant: unsupported scan type %T", src)
	}
}
