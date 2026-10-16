package parse_test

import (
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

type Widget struct{}
type BoxUnit struct{}

func (BoxUnit) ToBase(v float64) float64   { return v * 12 }
func (BoxUnit) FromBase(v float64) float64 { return v / 12 }
func (BoxUnit) Dimension() Widget          { return Widget{} }
