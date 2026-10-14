package quant_test

import (
	"math"
	"testing"

	"github.com/nodivbyzero/quant"
)

// The non-SI factors in this table are checked against NIST SP 811 Appendix
// B.8/B.9. The SI prefix cases follow the BIPM SI prefix table.
func TestReferenceConversionFactors(t *testing.T) {
	tests := []struct {
		name string
		got  float64
		want float64
		tol  float64
	}{
		{"inch to meter", quant.New[quant.Length](1, quant.Inch).To(quant.Meter), 0.0254, 1e-15},
		{"foot to meter", quant.New[quant.Length](1, quant.Foot).To(quant.Meter), 0.3048, 1e-15},
		{"mile to meter", quant.New[quant.Length](1, quant.Mile).To(quant.Meter), 1609.344, 1e-12},
		{"pound to kilogram", quant.New[quant.Mass](1, quant.Pound).To(quant.Kilogram), 0.45359237, 1e-15},
		{"us gallon to liter", quant.New[quant.Volume](1, quant.Gallon).To(quant.Liter), 3.785411784, 1e-12},
		{"psi to pascal", quant.New[quant.Pressure](1, quant.PSI).To(quant.Pascal), 6894.757293168, 1e-9},
		{"standard gravity", quant.New[quant.Acceleration](1, quant.StandardGravity).To(quant.MeterPerSecondSquared), 9.80665, 1e-14},
		{"fahrenheit to kelvin", quant.New[quant.Temperature](32, quant.Fahrenheit).To(quant.Kelvin), 273.15, 1e-12},
		{"rpm to hertz", quant.New[quant.Frequency](60, quant.RevolutionPerMinute).To(quant.Hertz), 1, 1e-15},
		{"degree to radian", quant.New[quant.Angle](180, quant.Degree).To(quant.Radian), math.Pi, 1e-14},
		{"kilowatt-hour to joule", quant.New[quant.Energy](1, quant.KilowattHour).To(quant.Joule), 3.6e6, 1e-8},
		{"byte to bit", quant.New[quant.Digital](1, quant.Byte).To(quant.Bit), 8, 1e-15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if math.Abs(tt.got-tt.want) > tt.tol {
				t.Fatalf("got %.17g, want %.17g ± %.2g", tt.got, tt.want, tt.tol)
			}
		})
	}
}

func TestUnitRoundTrips(t *testing.T) {
	tests := []struct {
		name string
		run  func() float64
	}{
		{"mass/pound", func() float64 { return quant.New[quant.Mass](12.5, quant.Pound).To(quant.Pound) }},
		{"mass/ounce", func() float64 { return quant.New[quant.Mass](12.5, quant.Ounce).To(quant.Ounce) }},
		{"length/inch", func() float64 { return quant.New[quant.Length](12.5, quant.Inch).To(quant.Inch) }},
		{"length/mile", func() float64 { return quant.New[quant.Length](12.5, quant.Mile).To(quant.Mile) }},
		{"area/acre", func() float64 { return quant.New[quant.Area](12.5, quant.Acre).To(quant.Acre) }},
		{"volume/gallon", func() float64 { return quant.New[quant.Volume](12.5, quant.Gallon).To(quant.Gallon) }},
		{"temperature/celsius", func() float64 { return quant.New[quant.Temperature](12.5, quant.Celsius).To(quant.Celsius) }},
		{"temperature/fahrenheit", func() float64 { return quant.New[quant.Temperature](12.5, quant.Fahrenheit).To(quant.Fahrenheit) }},
		{"time/hour", func() float64 { return quant.New[quant.Time](12.5, quant.Hour).To(quant.Hour) }},
		{"speed/mph", func() float64 { return quant.New[quant.Speed](12.5, quant.MilePerHour).To(quant.MilePerHour) }},
		{"pressure/psi", func() float64 { return quant.New[quant.Pressure](12.5, quant.PSI).To(quant.PSI) }},
		{"energy/kwh", func() float64 { return quant.New[quant.Energy](12.5, quant.KilowattHour).To(quant.KilowattHour) }},
		{"power/horsepower", func() float64 { return quant.New[quant.Power](12.5, quant.Horsepower).To(quant.Horsepower) }},
		{"angle/degree", func() float64 { return quant.New[quant.Angle](12.5, quant.Degree).To(quant.Degree) }},
		{"frequency/rpm", func() float64 {
			return quant.New[quant.Frequency](12.5, quant.RevolutionPerMinute).To(quant.RevolutionPerMinute)
		}},
		{"digital/kibibyte", func() float64 { return quant.New[quant.Digital](12.5, quant.Kibibyte).To(quant.Kibibyte) }},
		{"density/gcm3", func() float64 {
			return quant.New[quant.Density](12.5, quant.GramPerCubicCentimeter).To(quant.GramPerCubicCentimeter)
		}},
		{"resistance/kiloohm", func() float64 { return quant.New[quant.Resistance](12.5, quant.Kiloohm).To(quant.Kiloohm) }},
		{"voltage/kilovolt", func() float64 { return quant.New[quant.Voltage](12.5, quant.Kilovolt).To(quant.Kilovolt) }},
		{"current/kiloampere", func() float64 { return quant.New[quant.Current](12.5, quant.Kiloampere).To(quant.Kiloampere) }},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.run(); math.Abs(got-12.5) > 1e-12 {
				t.Fatalf("round trip changed value: got %.17g, want 12.5", got)
			}
		})
	}
}
