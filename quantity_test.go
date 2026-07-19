package quant_test

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/nodivbyzero/quant"
)

func TestMassConversion(t *testing.T) {
	assertConversion[quant.Mass](t, 10, quant.Pound, 4.5359237, quant.Kilogram)
	assertConversion[quant.Mass](t, 1000, quant.Microgram, 1, quant.Milligram)
	assertConversion[quant.Mass](t, 1000, quant.Gram, 1, quant.Kilogram)
	assertConversion[quant.Mass](t, 1, quant.Stone, 14, quant.Pound)
	assertConversion[quant.Mass](t, 1, quant.MetricTon, 1000, quant.Kilogram)
	assertConversion[quant.Mass](t, 1, quant.Tonne, 1000, quant.Kilogram)
}

func TestLengthConversion(t *testing.T) {
	assertConversion[quant.Length](t, 1, quant.Mile, 1.609344, quant.Kilometer)
	assertConversion[quant.Length](t, 1, quant.NauticalMile, 1852, quant.Meter)
	assertConversion[quant.Length](t, 1, quant.Fathom, 6, quant.Foot)
	assertConversion[quant.Length](t, 1_000_000, quant.Nanometer, 1, quant.Millimeter)
}

func TestTimeConversion(t *testing.T) {
	assertConversion[quant.Time](t, 2, quant.Hour, 120.0, quant.Minute)
	assertConversion[quant.Time](t, 1, quant.Year, 12, quant.Month)
	assertConversion[quant.Time](t, 1, quant.Millisecond, 1000, quant.Microsecond)
	assertConversion[quant.Time](t, 1, quant.Decade, 10, quant.Year)
	assertConversion[quant.Time](t, 1, quant.Century, 10, quant.Decade)
}

func TestTemperatureConversion(t *testing.T) {
	assertConversion[quant.Temperature](t, 25, quant.Celsius, 298.15, quant.Kelvin)
	assertConversion[quant.Temperature](t, 32, quant.Fahrenheit, 273.15, quant.Kelvin)
	assertConversion[quant.Temperature](t, 491.67, quant.Rankine, 273.15, quant.Kelvin)
}

func TestAreaConversion(t *testing.T) {
	assertConversion[quant.Area](t, 1, quant.Acre, 4046.8564224, quant.SquareMeter)
	assertConversion[quant.Area](t, 1, quant.SquareCentimeter, 100, quant.SquareMillimeter)
	assertConversion[quant.Area](t, 1, quant.Hectare, 10000, quant.SquareMeter)
	assertConversion[quant.Area](t, 1, quant.SquareKilometer, 100, quant.Hectare)
	assertConversion[quant.Area](t, 1, quant.SquareFoot, 144, quant.SquareInch)
	assertConversion[quant.Area](t, 1, quant.SquareMile, 640, quant.Acre)
}

func TestAcidityConversion(t *testing.T) {
	assertConversion[quant.Acidity](t, 3, quant.PH, 11, quant.POH)
	assertConversion[quant.Acidity](t, 11, quant.POH, 3, quant.PH)
}

func TestVolumeConversion(t *testing.T) {
	assertConversion[quant.Volume](t, 1, quant.Gallon, 3.785411784, quant.Liter)
	assertConversion[quant.Volume](t, 1, quant.Megaliter, 1000, quant.CubicMeter)
	assertConversion[quant.Volume](t, 1, quant.Gigaliter, 1000, quant.Megaliter)
}

func TestVolumeFlowRateConversion(t *testing.T) {
	assertConversion[quant.VolumeFlowRate](t, 60, quant.LiterPerMinute, 1, quant.LiterPerSecond)
	assertConversion[quant.VolumeFlowRate](t, 1, quant.CubicFootPerMinute, 60, quant.CubicFootPerHour)
	assertConversion[quant.VolumeFlowRate](t, 1, quant.GallonPerHour, 1.0/60.0, quant.GallonPerMinute)
}

func TestFrequencyConversion(t *testing.T) {
	assertConversion[quant.Frequency](t, 60, quant.RevolutionPerMinute, 1, quant.Hertz)
	assertConversion[quant.Frequency](t, 360, quant.DegreePerSecond, 1, quant.Hertz)
}

func TestSpeedConversion(t *testing.T) {
	assertConversion[quant.Speed](t, 36, quant.KilometerPerHour, 10, quant.MeterPerSecond)
	assertConversion[quant.Speed](t, 1, quant.MilePerHour, 1.609344, quant.KilometerPerHour)
	assertConversion[quant.Speed](t, 1, quant.InchPerHour, 25.4, quant.MillimeterPerHour)
}

func TestTorqueConversion(t *testing.T) {
	assertConversion[quant.Torque](t, 1, quant.PoundForceFoot, 1.3558179483314004, quant.NewtonMeter)
}

func TestPaceConversion(t *testing.T) {
	assertConversion[quant.Pace](t, 5, quant.MinutePerKilometer, 0.3, quant.SecondPerMeter)
	assertConversion[quant.Pace](t, 1, quant.MinutePerMile, 60.0/1609.344, quant.SecondPerMeter)
}

func TestPressureConversion(t *testing.T) {
	assertConversion[quant.Pressure](t, 1, quant.Bar, 100000, quant.Pascal)
	assertConversion[quant.Pressure](t, 1, quant.MeterOfWater, 9806.65, quant.Pascal)
	assertConversion[quant.Pressure](t, 1, quant.MillimeterOfMercury, 133.322387415, quant.Pascal)
}

func TestDigitalConversion(t *testing.T) {
	assertConversion[quant.Digital](t, 1, quant.Byte, 8, quant.Bit)
	assertConversion[quant.Digital](t, 1, quant.Megabyte, 8, quant.Megabit)
	assertConversion[quant.Digital](t, 1, quant.Petabyte, 1000, quant.Terabyte)
	assertConversion[quant.Digital](t, 1, quant.Kibibyte, 1024, quant.Byte)
}

func TestIlluminanceConversion(t *testing.T) {
	assertConversion[quant.Illuminance](t, 1, quant.FootCandle, 10.763910416709722, quant.Lux)
}

func TestPartsPerConversion(t *testing.T) {
	assertConversion[quant.PartsPer](t, 1, quant.PPM, 1000, quant.PPB)
}

func TestElectricalConversions(t *testing.T) {
	assertConversion[quant.Voltage](t, 1, quant.Kilovolt, 1000, quant.Volt)
	assertConversion[quant.Current](t, 1, quant.Kiloampere, 1000, quant.Ampere)
	assertConversion[quant.Power](t, 1, quant.Megawatt, 1000, quant.Kilowatt)
	assertConversion[quant.Power](t, 1, quant.Horsepower, 745.6998715822702, quant.Watt)
	assertConversion[quant.Power](t, 1, quant.MetricHorsepower, 735.49875, quant.Watt)
	assertConversion[quant.Power](t, 1, quant.BTUPerSecond, 1055.05585262, quant.Watt)
	assertConversion[quant.Power](t, 1, quant.FootPoundForcePerSecond, 1.3558179483314004, quant.Watt)
	assertConversion[quant.ApparentPower](t, 1, quant.MegaVoltAmpere, 1000, quant.KiloVoltAmpere)
	assertConversion[quant.ReactivePower](t, 1, quant.MegaVoltAmpereReactive, 1000, quant.KiloVoltAmpereReactive)
}

func TestElectricalConductivityConversions(t *testing.T) {
	assertConversion[quant.ElectricalConductivity](t, 1, quant.SiemensPerCentimeter, 100, quant.SiemensPerMeter)
	assertConversion[quant.ElectricalConductivity](t, 1, quant.MillisiemensPerCentimeter, 1000, quant.MicrosiemensPerCentimeter)
	assertConversion[quant.ElectricalConductivity](t, 1000, quant.MicrosiemensPerMeter, 1, quant.MillisiemensPerMeter)
}

func TestEnergyConversions(t *testing.T) {
	assertConversion[quant.Energy](t, 1, quant.KilowattHour, 3.6e6, quant.Joule)
	assertConversion[quant.Energy](t, 1, quant.WattSecond, 1, quant.Joule)
	assertConversion[quant.Energy](t, 1, quant.WattMinute, 60, quant.Joule)
	assertConversion[quant.Energy](t, 1, quant.Gigajoule, 1000, quant.Megajoule)
	assertConversion[quant.ReactiveEnergy](t, 1, quant.MegaVoltAmpereReactiveHour, 1000, quant.KiloVoltAmpereReactiveHour)
}

func TestAngleConversion(t *testing.T) {
	assertConversion[quant.Angle](t, 180, quant.Degree, math.Pi, quant.Radian)
}

func TestChargeConversion(t *testing.T) {
	assertConversion[quant.Charge](t, 1, quant.Coulomb, 1000, quant.Millicoulomb)
	assertConversion[quant.Charge](t, 1, quant.Microcoulomb, 1000, quant.Nanocoulomb)
}

func TestForceConversion(t *testing.T) {
	assertConversion[quant.Force](t, 1, quant.Kilonewton, 1000, quant.Newton)
	assertConversion[quant.Force](t, 1, quant.KilogramForce, 9.80665, quant.Newton)
}

func TestAccelerationConversion(t *testing.T) {
	assertConversion[quant.Acceleration](t, 1, quant.GForce, 9.80665, quant.MeterPerSecondSquared)
	assertConversion[quant.Acceleration](t, 1, quant.StandardGravity, 1, quant.GForce)
}

func TestPiecesConversion(t *testing.T) {
	assertConversion[quant.Pieces](t, 1, quant.Dozen, 12, quant.Piece)
	assertConversion[quant.Pieces](t, 1, quant.GreatGross, 12, quant.Gross)
	assertConversion[quant.Pieces](t, 1, quant.Ream, 500, quant.Piece)
}

func TestArithmetic(t *testing.T) {
	a := quant.New[quant.Length](2, quant.Kilometer)
	b := quant.New[quant.Length](500, quant.Meter)

	sum := a.Add(b)
	diff := a.Sub(b)

	if got, want := sum.To(quant.Meter), 2500.0; !almostEqual(got, want) {
		t.Fatalf("sum mismatch: got %.10f want %.10f", got, want)
	}

	if got, want := diff.To(quant.Kilometer), 1.5; !almostEqual(got, want) {
		t.Fatalf("difference mismatch: got %.10f want %.10f", got, want)
	}
}

func TestDerivedDivision(t *testing.T) {
	distance := quant.Kilometers(5)
	duration := quant.Minutes(30)

	speed := distance.Div(duration)

	if got, want := speed.To(quant.KilometerPerHour), 10.0; !almostEqual(got, want) {
		t.Fatalf("derived speed mismatch: got %.10f want %.10f", got, want)
	}
}

func TestAmountBuilder(t *testing.T) {
	q := quant.Amount[quant.Mass](10).From(quant.Pound)

	if got, want := q.Value(quant.Kilogram), 4.5359237; !almostEqual(got, want) {
		t.Fatalf("builder conversion mismatch: got %.10f want %.10f", got, want)
	}
}

func TestConvenienceConstructors(t *testing.T) {
	if got := quant.Pounds(10).To(quant.Kilogram); !almostEqual(got, 4.5359237) {
		t.Fatalf("Pounds constructor mismatch: got %.10f", got)
	}

	if got := quant.Kilometers(5).To(quant.Meter); !almostEqual(got, 5000) {
		t.Fatalf("Kilometers constructor mismatch: got %.10f", got)
	}

	if got := quant.DegreesCelsius(100).To(quant.Kelvin); !almostEqual(got, 373.15) {
		t.Fatalf("DegreesCelsius constructor mismatch: got %.10f", got)
	}

	if got := quant.LitersPerMinute(120).To(quant.LiterPerSecond); !almostEqual(got, 2) {
		t.Fatalf("LitersPerMinute constructor mismatch: got %.10f", got)
	}

	if got := quant.Kibibytes(1).To(quant.Byte); !almostEqual(got, 1024) {
		t.Fatalf("Kibibytes constructor mismatch: got %.10f", got)
	}

	if got := quant.GForces(1).To(quant.MeterPerSecondSquared); !almostEqual(got, 9.80665) {
		t.Fatalf("GForces constructor mismatch: got %.10f", got)
	}
}

func TestStringFormattingUsesBaseUnit(t *testing.T) {
	if got := quant.Pounds(10).String(); got != "4.5359237 kg" {
		t.Fatalf("mass String mismatch: got %q", got)
	}

	if got := quant.Kilometers(5).String(); got != "5000 m" {
		t.Fatalf("length String mismatch: got %q", got)
	}

	if got := quant.DegreesCelsius(25).String(); got != "298.15 K" {
		t.Fatalf("temperature String mismatch: got %q", got)
	}
}

func TestExplicitFormattingHelpers(t *testing.T) {
	if got := quant.Pounds(10).Format(quant.Pound); got != "10 lb" {
		t.Fatalf("mass Format mismatch: got %q", got)
	}

	if got := quant.Kilometers(5).Format(quant.Mile); got != "3.10685596119 mi" {
		t.Fatalf("length Format mismatch: got %q", got)
	}

	if got := quant.DegreesCelsius(25).Format(quant.Celsius); got != "25 C" {
		t.Fatalf("temperature Format mismatch: got %q", got)
	}

	if got := quant.LitersPerMinute(120).Format(quant.LiterPerSecond); got != "2 l/s" {
		t.Fatalf("flow Format mismatch: got %q", got)
	}
}

func TestJSONMarshalingUsesBaseUnit(t *testing.T) {
	data, err := json.Marshal(quant.Pounds(10))
	if err != nil {
		t.Fatalf("MarshalJSON error: %v", err)
	}

	if got := string(data); got != "4.535923700000001" {
		t.Fatalf("MarshalJSON mismatch: got %q", got)
	}

	var q quant.Quantity[quant.Mass]
	if err := json.Unmarshal(data, &q); err != nil {
		t.Fatalf("UnmarshalJSON error: %v", err)
	}

	if got, want := q.To(quant.Kilogram), 4.5359237; !almostEqual(got, want) {
		t.Fatalf("UnmarshalJSON mismatch: got %.10f want %.10f", got, want)
	}
}

func TestJSONMarshalingInStruct(t *testing.T) {
	type payload struct {
		Mass quant.Quantity[quant.Mass] `json:"mass"`
	}

	input := payload{Mass: quant.Kilograms(5)}
	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Marshal struct error: %v", err)
	}

	if got := string(data); got != `{"mass":5}` {
		t.Fatalf("Marshal struct mismatch: got %q", got)
	}

	var out payload
	if err := json.Unmarshal([]byte(`{"mass":2.5}`), &out); err != nil {
		t.Fatalf("Unmarshal struct error: %v", err)
	}

	if got, want := out.Mass.To(quant.Kilogram), 2.5; !almostEqual(got, want) {
		t.Fatalf("Unmarshal struct mismatch: got %.10f want %.10f", got, want)
	}
}

func TestTextMarshalingUsesBaseUnit(t *testing.T) {
	data, err := quant.Pounds(10).MarshalText()
	if err != nil {
		t.Fatalf("MarshalText error: %v", err)
	}

	if got := string(data); got != "4.5359237" {
		t.Fatalf("MarshalText mismatch: got %q", got)
	}

	var q quant.Quantity[quant.Mass]
	if err := q.UnmarshalText([]byte("5")); err != nil {
		t.Fatalf("UnmarshalText error: %v", err)
	}

	if got, want := q.To(quant.Kilogram), 5.0; !almostEqual(got, want) {
		t.Fatalf("UnmarshalText mismatch: got %.10f want %.10f", got, want)
	}
}

func TestSQLQuantityInterop(t *testing.T) {
	value, err := quant.SQL(quant.Kilograms(5)).Value()
	if err != nil {
		t.Fatalf("driver.Valuer error: %v", err)
	}

	gotValue, ok := value.(float64)
	if !ok {
		t.Fatalf("driver.Valuer type mismatch: got %T", value)
	}
	if !almostEqual(gotValue, 5) {
		t.Fatalf("driver.Valuer value mismatch: got %.10f", gotValue)
	}

	var q quant.SQLQuantity[quant.Mass]
	if err := q.Scan(float64(2.5)); err != nil {
		t.Fatalf("Scan float64 error: %v", err)
	}
	if got, want := q.To(quant.Kilogram), 2.5; !almostEqual(got, want) {
		t.Fatalf("Scan float64 mismatch: got %.10f want %.10f", got, want)
	}

	if err := q.Scan([]byte("3.75")); err != nil {
		t.Fatalf("Scan []byte error: %v", err)
	}
	if got, want := q.To(quant.Kilogram), 3.75; !almostEqual(got, want) {
		t.Fatalf("Scan []byte mismatch: got %.10f want %.10f", got, want)
	}
}

func TestSQLQuantityImplementsValuer(t *testing.T) {
	var _ driver.Valuer = quant.SQL(quant.Kilograms(1))
}

func Example_massConversion() {
	kg := quant.Amount[quant.Mass](10).From(quant.Pound).To(quant.Kilogram)
	fmt.Printf("%.6f\n", kg)
	// Output:
	// 4.535924
}

func Example_lengthConversion() {
	km := quant.New[quant.Length](3, quant.Mile).To(quant.Kilometer)
	fmt.Printf("%.6f\n", km)
	// Output:
	// 4.828032
}

func Example_arithmetic() {
	total := quant.New[quant.Length](750, quant.Meter).Add(quant.New[quant.Length](1.25, quant.Kilometer))
	fmt.Printf("%.0f\n", total.To(quant.Meter))
	// Output:
	// 2000
}

func almostEqual(got, want float64) bool {
	return math.Abs(got-want) < 1e-9
}

func assertConversion[D any, UFrom quant.Unit[D], UTo quant.Unit[D]](t *testing.T, input float64, from UFrom, want float64, to UTo) {
	t.Helper()

	got := quant.New[D](input, from).To(to)
	if !almostEqual(got, want) {
		t.Fatalf("conversion mismatch: got %.12f want %.12f", got, want)
	}
}

// Type safety is enforced at compile time:
//
//	length := quant.New[quant.Length](5, quant.Meter)
//	mass := quant.New[quant.Mass](2, quant.Kilogram)
//	_ = length.Add(mass)        // does not compile
//	_ = length.To(quant.Pound)  // does not compile
