package quant

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
	case ReactivePower:
		return any(VoltAmpereReactive).(Unit[D])
	case Energy:
		return any(Joule).(Unit[D])
	case ReactiveEnergy:
		return any(VoltAmpereReactiveHour).(Unit[D])
	case Angle:
		return any(Radian).(Unit[D])
	case Charge:
		return any(Coulomb).(Unit[D])
	case Force:
		return any(Newton).(Unit[D])
	case Acceleration:
		return any(MeterPerSecondSquared).(Unit[D])
	case Pieces:
		return any(Piece).(Unit[D])
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
	case any(Pascal):
		return "Pa"
	case any(Bar):
		return "bar"
	case any(PSI):
		return "psi"
	case any(Bit):
		return "bit"
	case any(Byte):
		return "byte"
	case any(Petabyte):
		return "PB"
	case any(Kibibyte):
		return "KiB"
	case any(Watt):
		return "W"
	case any(Kilowatt):
		return "kW"
	case any(Joule):
		return "J"
	case any(Radian):
		return "rad"
	case any(Degree):
		return "deg"
	case any(Coulomb):
		return "C"
	case any(Newton):
		return "N"
	case any(MeterPerSecondSquared):
		return "m/s2"
	case any(GForce):
		return "g"
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
