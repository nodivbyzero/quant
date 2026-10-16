// Package parse provides optional string parsing for quant quantities.
//
// The core quant package deliberately has no unit-name parser. Callers opt in
// here and provide the unit registry they want to accept.
package parse

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nodivbyzero/quant"
)

// UnitSet maps accepted symbols to units for one dimension.
type UnitSet[D quant.Dimension] map[string]quant.Unit[D]

// Parse parses a scalar followed by an optional unit symbol. Both "5 km" and
// "5km" are accepted. A missing symbol uses the empty-string entry in units.
func Parse[D quant.Dimension](input string, units UnitSet[D]) (quant.Quantity[D], error) {
	value, symbol, err := split(input)
	if err != nil {
		return quant.Quantity[D]{}, err
	}

	u, ok := units[symbol]
	if !ok {
		return quant.Quantity[D]{}, fmt.Errorf("parse: unknown unit %q", symbol)
	}

	return quant.New[D](value, u), nil
}

// ParseLength parses a length using the built-in length symbols in LengthUnits.
func ParseLength(input string) (quant.Quantity[quant.Length], error) {
	return Parse(input, LengthUnits)
}

// ParseMass parses a mass using the built-in mass symbols in MassUnits.
func ParseMass(input string) (quant.Quantity[quant.Mass], error) {
	return Parse(input, MassUnits)
}

// ParseTime parses a time using the built-in time symbols in TimeUnits.
func ParseTime(input string) (quant.Quantity[quant.Time], error) {
	return Parse(input, TimeUnits)
}

// ParseTemperature parses a temperature using the built-in temperature symbols
// in TemperatureUnits.
func ParseTemperature(input string) (quant.Quantity[quant.Temperature], error) {
	return Parse(input, TemperatureUnits)
}

var LengthUnits = UnitSet[quant.Length]{
	"": quant.Meter, "m": quant.Meter, "km": quant.Kilometer, "cm": quant.Centimeter,
	"mm": quant.Millimeter, "mi": quant.Mile, "ft": quant.Foot, "in": quant.Inch,
}

var MassUnits = UnitSet[quant.Mass]{
	"": quant.Kilogram, "kg": quant.Kilogram, "g": quant.Gram, "mg": quant.Milligram,
	"lb": quant.Pound, "oz": quant.Ounce,
}

var TimeUnits = UnitSet[quant.Time]{
	"": quant.Second, "s": quant.Second, "ms": quant.Millisecond, "min": quant.Minute,
	"h": quant.Hour, "d": quant.Day,
}

var TemperatureUnits = UnitSet[quant.Temperature]{
	"": quant.Kelvin, "K": quant.Kelvin, "C": quant.Celsius, "F": quant.Fahrenheit,
}

func split(input string) (float64, string, error) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, "", fmt.Errorf("parse: empty input")
	}

	end := 0
	if s[0] == '+' || s[0] == '-' {
		end++
	}
	digits := false
	for end < len(s) && s[end] >= '0' && s[end] <= '9' {
		end++
		digits = true
	}
	if end < len(s) && s[end] == '.' {
		end++
		for end < len(s) && s[end] >= '0' && s[end] <= '9' {
			end++
			digits = true
		}
	}
	if !digits {
		return 0, "", fmt.Errorf("parse: invalid number in %q", input)
	}
	if end < len(s) && (s[end] == 'e' || s[end] == 'E') {
		end++
		if end < len(s) && (s[end] == '+' || s[end] == '-') {
			end++
		}
		expStart := end
		for end < len(s) && s[end] >= '0' && s[end] <= '9' {
			end++
		}
		if expStart == end {
			return 0, "", fmt.Errorf("parse: invalid exponent in %q", input)
		}
	}

	value, err := strconv.ParseFloat(s[:end], 64)
	if err != nil {
		return 0, "", fmt.Errorf("parse: %w", err)
	}
	symbol := strings.TrimSpace(s[end:])
	if strings.ContainsAny(symbol, " \t\r\n") {
		return 0, "", fmt.Errorf("parse: invalid unit %q", symbol)
	}
	return value, symbol, nil
}
