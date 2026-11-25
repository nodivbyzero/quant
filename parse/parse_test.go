package parse_test

import (
	"encoding/json"
	"errors"
	"math"
	"testing"

	"github.com/nodivbyzero/quant"
	"github.com/nodivbyzero/quant/parse"
)

func TestBuiltInParsers(t *testing.T) {
	length, err := parse.ParseLength("5 km")
	if err != nil || length.To(quant.Meter) != 5000 {
		t.Fatalf("length parse: value=%v err=%v", length.To(quant.Meter), err)
	}

	mass, err := parse.ParseMass("10lb")
	if err != nil || math.Abs(mass.To(quant.Kilogram)-4.5359237) > 1e-12 {
		t.Fatalf("mass parse: value=%v err=%v", mass.To(quant.Kilogram), err)
	}

	temperature, err := parse.ParseTemperature("32 F")
	if err != nil || math.Abs(temperature.To(quant.Kelvin)-273.15) > 1e-12 {
		t.Fatalf("temperature parse: value=%v err=%v", temperature.To(quant.Kelvin), err)
	}
}

func TestCustomRegistry(t *testing.T) {
	units := parse.UnitSet[Widget]{
		"box": BoxUnit{},
	}
	q, err := parse.Parse("2 box", units)
	if err != nil || q.To(BoxUnit{}) != 2 {
		t.Fatalf("custom parse: value=%v err=%v", q.To(BoxUnit{}), err)
	}
}

func TestRegistryIsStrictByDefaultAndSupportsExplicitDefaults(t *testing.T) {
	strict := parse.Lengths()
	if _, err := strict.Parse("5"); !errors.Is(err, parse.ErrMissingUnit) {
		t.Fatalf("strict parse error = %v, want ErrMissingUnit", err)
	}

	withDefault := strict.WithDefault(quant.Meter)
	q, err := withDefault.Parse("5")
	if err != nil || q.To(quant.Meter) != 5 {
		t.Fatalf("default parse: value=%v err=%v", q.To(quant.Meter), err)
	}
	if _, err := strict.Parse("5"); !errors.Is(err, parse.ErrMissingUnit) {
		t.Fatalf("WithDefault mutated original registry: %v", err)
	}
}

func TestRegistryErrors(t *testing.T) {
	registry := parse.Lengths()
	_, err := registry.Parse("5 furlong")
	if !errors.Is(err, parse.ErrUnknownUnit) {
		t.Fatalf("error = %v, want ErrUnknownUnit", err)
	}
	var parseErr *parse.Error
	if !errors.As(err, &parseErr) || parseErr.Pos != 1 {
		t.Fatalf("error = %#v, want position 1", err)
	}

	if _, err := registry.Parse("five m"); !errors.Is(err, parse.ErrSyntax) {
		t.Fatalf("error = %v, want ErrSyntax", err)
	}
}

func TestRegistryAliasesAndCollisions(t *testing.T) {
	registry := parse.NewRegistry[quant.Length]().With("km", quant.Kilometer, "kilometer", "kilometers")
	q, err := registry.Parse("2 kilometers")
	if err != nil || q.To(quant.Meter) != 2000 {
		t.Fatalf("alias parse: value=%v err=%v", q.To(quant.Meter), err)
	}

	deferred := func() (panicked bool) {
		defer func() { panicked = recover() != nil }()
		_ = registry.With("km", quant.Kilometer)
		return false
	}()
	if !deferred {
		t.Fatal("expected duplicate symbol to panic")
	}
}

func TestBuiltInRegistryAccessorsAndNormalization(t *testing.T) {
	registries := []struct {
		name  string
		parse func(string) error
	}{
		{"areas", func(s string) error { _, err := parse.Areas().Parse(s); return err }},
		{"volumes", func(s string) error { _, err := parse.Volumes().Parse(s); return err }},
		{"speeds", func(s string) error { _, err := parse.Speeds().Parse(s); return err }},
		{"accelerations", func(s string) error { _, err := parse.Accelerations().Parse(s); return err }},
		{"forces", func(s string) error { _, err := parse.Forces().Parse(s); return err }},
		{"pressures", func(s string) error { _, err := parse.Pressures().Parse(s); return err }},
		{"energies", func(s string) error { _, err := parse.Energies().Parse(s); return err }},
		{"powers", func(s string) error { _, err := parse.Powers().Parse(s); return err }},
		{"frequencies", func(s string) error { _, err := parse.Frequencies().Parse(s); return err }},
		{"angles", func(s string) error { _, err := parse.Angles().Parse(s); return err }},
		{"digitals", func(s string) error { _, err := parse.Digitals().Parse(s); return err }},
		{"data rates", func(s string) error { _, err := parse.DataRates().Parse(s); return err }},
		{"voltages", func(s string) error { _, err := parse.Voltages().Parse(s); return err }},
		{"currents", func(s string) error { _, err := parse.Currents().Parse(s); return err }},
		{"resistances", func(s string) error { _, err := parse.Resistances().Parse(s); return err }},
		{"capacitances", func(s string) error { _, err := parse.Capacitances().Parse(s); return err }},
		{"charges", func(s string) error { _, err := parse.Charges().Parse(s); return err }},
		{"temperature deltas", func(s string) error { _, err := parse.TemperatureDeltas().Parse(s); return err }},
		{"torques", func(s string) error { _, err := parse.Torques().Parse(s); return err }},
		{"densities", func(s string) error { _, err := parse.Densities().Parse(s); return err }},
		{"volume flow rates", func(s string) error { _, err := parse.VolumeFlowRates().Parse(s); return err }},
		{"angular velocities", func(s string) error { _, err := parse.AngularVelocities().Parse(s); return err }},
		{"angular accelerations", func(s string) error { _, err := parse.AngularAccelerations().Parse(s); return err }},
		{"paces", func(s string) error { _, err := parse.Paces().Parse(s); return err }},
		{"parts per", func(s string) error { _, err := parse.PartsPer().Parse(s); return err }},
		{"pieces", func(s string) error { _, err := parse.Pieces().Parse(s); return err }},
		{"acidities", func(s string) error { _, err := parse.Acidities().Parse(s); return err }},
		{"concentrations", func(s string) error { _, err := parse.Concentrations().Parse(s); return err }},
		{"molalities", func(s string) error { _, err := parse.Molalities().Parse(s); return err }},
		{"dynamic viscosities", func(s string) error { _, err := parse.DynamicViscosities().Parse(s); return err }},
		{"kinematic viscosities", func(s string) error { _, err := parse.KinematicViscosities().Parse(s); return err }},
		{"thermal conductivities", func(s string) error { _, err := parse.ThermalConductivities().Parse(s); return err }},
		{"specific heat capacities", func(s string) error { _, err := parse.SpecificHeatCapacities().Parse(s); return err }},
		{"radioactivities", func(s string) error { _, err := parse.Radioactivities().Parse(s); return err }},
		{"absorbed doses", func(s string) error { _, err := parse.AbsorbedDoses().Parse(s); return err }},
		{"equivalent doses", func(s string) error { _, err := parse.EquivalentDoses().Parse(s); return err }},
		{"catalytic activities", func(s string) error { _, err := parse.CatalyticActivities().Parse(s); return err }},
		{"illuminances", func(s string) error { _, err := parse.Illuminances().Parse(s); return err }},
		{"luminous fluxes", func(s string) error { _, err := parse.LuminousFluxes().Parse(s); return err }},
		{"luminous intensities", func(s string) error { _, err := parse.LuminousIntensities().Parse(s); return err }},
		{"magnetic fluxes", func(s string) error { _, err := parse.MagneticFluxes().Parse(s); return err }},
		{"magnetic flux densities", func(s string) error { _, err := parse.MagneticFluxDensities().Parse(s); return err }},
	}
	for _, registry := range registries {
		t.Run(registry.name, func(t *testing.T) {
			if err := registry.parse("1 "); err == nil {
				t.Fatal("expected missing unit")
			}
		})
	}

	if _, err := parse.Lengths().Parse("2 Kilometers"); err != nil {
		t.Fatalf("long-name alias: %v", err)
	}
	if _, err := parse.Lengths().Parse("2\u00a0km"); err != nil {
		t.Fatalf("non-breaking space: %v", err)
	}
	if _, err := parse.Lengths().Parse("2\u2009km"); err != nil {
		t.Fatalf("thin space: %v", err)
	}
	if _, err := parse.Temperatures().Parse("20 °C"); err != nil {
		t.Fatalf("degree normalization: %v", err)
	}
	if _, err := parse.Capacitances().Parse("2 µF"); err != nil {
		t.Fatalf("micro normalization: %v", err)
	}
	if _, err := parse.Resistances().Parse("2 Ω"); err != nil {
		t.Fatalf("ohm normalization: %v", err)
	}
	if _, err := parse.ParseTemperatureDelta("5 °C-delta"); err != nil {
		t.Fatalf("temperature delta parse: %v", err)
	}
}

func TestNumberSyntax(t *testing.T) {
	for _, input := range []string{"NaN m", "Inf m", "+Inf m", "1_000 m", "1e m"} {
		if _, err := parse.Lengths().Parse(input); !errors.Is(err, parse.ErrSyntax) {
			t.Errorf("Parse(%q) error = %v, want ErrSyntax", input, err)
		}
	}
	q, err := parse.Lengths().Parse("1.25e3 m")
	if err != nil || q.To(quant.Meter) != 1250 {
		t.Fatalf("exponent parse: q=%v err=%v", q, err)
	}
}

func TestValueBoundary(t *testing.T) {
	v := parse.NewValue(parse.Lengths())
	if err := v.Set("2 kilometers"); err != nil {
		t.Fatalf("Set: %v", err)
	}
	q, ok := v.Quantity()
	if !ok || q.To(quant.Meter) != 2000 {
		t.Fatalf("Quantity: q=%v set=%v", q.To(quant.Meter), ok)
	}
	if got := v.String(); got != "2 kilometers" {
		t.Fatalf("String: got %q", got)
	}
	text, err := v.MarshalText()
	if err != nil || string(text) != "2 kilometers" {
		t.Fatalf("MarshalText: got %q err=%v", text, err)
	}
	jsonText, err := json.Marshal(v)
	if err != nil || string(jsonText) != `"2 kilometers"` {
		t.Fatalf("JSON marshal: got %q err=%v", jsonText, err)
	}
	if err := v.UnmarshalText([]byte("3 m")); err != nil {
		t.Fatalf("UnmarshalText: %v", err)
	}
	q, _ = v.Quantity()
	if q.To(quant.Meter) != 3 {
		t.Fatalf("UnmarshalText value: got %v", q.To(quant.Meter))
	}
	if err := v.Set("bad"); err == nil || !errors.Is(err, parse.ErrSyntax) {
		t.Fatalf("invalid Set error: %v", err)
	}
	q, _ = v.Quantity()
	if q.To(quant.Meter) != 3 {
		t.Fatalf("failed Set changed value: got %v", q.To(quant.Meter))
	}
}

func TestValueDefaultAndSetQuantity(t *testing.T) {
	v := parse.NewValue(parse.Lengths().WithDefault(quant.Meter))
	if err := v.Set("5"); err != nil {
		t.Fatalf("default Set: %v", err)
	}
	q, _ := v.Quantity()
	if q.To(quant.Meter) != 5 {
		t.Fatalf("default value: got %v", q.To(quant.Meter))
	}
	v.SetQuantity(quant.Kilometers(2))
	if got := v.String(); got != "2000 m" {
		t.Fatalf("SetQuantity String: got %q", got)
	}
}

func TestCompoundUnitsAreExplicitRegistryEntries(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  float64
		to    quant.Unit[quant.Speed]
	}{
		{name: "kilometers per hour", input: "90 km/h", want: 90, to: quant.KilometerPerHour},
		{name: "feet per second", input: "88 ft/s", want: 60, to: quant.MilePerHour},
		{name: "inches per hour", input: "63360 in/h", want: 1, to: quant.MilePerHour},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q, err := parse.Speeds().Parse(test.input)
			if err != nil {
				t.Fatalf("Parse: %v", err)
			}
			if got := q.To(test.to); math.Abs(got-test.want) > 1e-10 {
				t.Fatalf("value: got %.15g want %.15g", got, test.want)
			}
		})
	}

	pace, err := parse.Paces().Parse("1 min/mi")
	if err != nil {
		t.Fatalf("pace Parse: %v", err)
	}
	if got := pace.To(quant.MinutePerMile); got != 1 {
		t.Fatalf("pace value: got %v want 1", got)
	}
}

func TestBuiltInRegistryFormatRoundTrips(t *testing.T) {
	tests := []struct {
		name string
		run  func(t *testing.T)
	}{
		{"length", func(t *testing.T) {
			for _, unit := range []quant.Unit[quant.Length]{quant.Meter, quant.Kilometer, quant.Centimeter, quant.Millimeter, quant.Mile, quant.Foot, quant.Inch} {
				q := quant.New[quant.Length](12.5, unit)
				parsed, err := parse.Lengths().Parse(q.Format(unit))
				if err != nil || math.Abs(parsed.To(unit)-12.5) > 1e-10 {
					t.Fatalf("unit %q: parsed=%v err=%v", quant.UnitSymbol(unit), parsed.To(unit), err)
				}
			}
		}},
		{"mass", func(t *testing.T) {
			for _, unit := range []quant.Unit[quant.Mass]{quant.Kilogram, quant.Gram, quant.Milligram, quant.Pound, quant.Ounce} {
				q := quant.New[quant.Mass](12.5, unit)
				parsed, err := parse.Masses().Parse(q.Format(unit))
				if err != nil || math.Abs(parsed.To(unit)-12.5) > 1e-10 {
					t.Fatalf("unit %q: parsed=%v err=%v", quant.UnitSymbol(unit), parsed.To(unit), err)
				}
			}
		}},
		{"time", func(t *testing.T) {
			for _, unit := range []quant.Unit[quant.Time]{quant.Second, quant.Millisecond, quant.Minute, quant.Hour, quant.Day} {
				q := quant.New[quant.Time](12.5, unit)
				parsed, err := parse.Times().Parse(q.Format(unit))
				if err != nil || math.Abs(parsed.To(unit)-12.5) > 1e-10 {
					t.Fatalf("unit %q: parsed=%v err=%v", quant.UnitSymbol(unit), parsed.To(unit), err)
				}
			}
		}},
		{"temperature", func(t *testing.T) {
			for _, unit := range []quant.Unit[quant.Temperature]{quant.Kelvin, quant.Celsius, quant.Fahrenheit} {
				q := quant.New[quant.Temperature](12.5, unit)
				parsed, err := parse.Temperatures().Parse(q.Format(unit))
				if err != nil || math.Abs(parsed.To(unit)-12.5) > 1e-10 {
					t.Fatalf("unit %q: parsed=%v err=%v", quant.UnitSymbol(unit), parsed.To(unit), err)
				}
			}
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, tt.run)
	}
}

type Widget struct{}
type BoxUnit struct{}

func (BoxUnit) ToBase(v float64) float64   { return v * 12 }
func (BoxUnit) FromBase(v float64) float64 { return v / 12 }
func (BoxUnit) Dimension() Widget          { return Widget{} }
