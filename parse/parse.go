// Package parse provides optional string parsing for quant quantities.
//
// The core quant package deliberately has no unit-name parser. Callers opt in
// here and provide the unit registry they want to accept.
package parse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/nodivbyzero/quant"
)

var (
	// ErrSyntax reports malformed numeric input.
	ErrSyntax = errors.New("invalid syntax")
	// ErrUnknownUnit reports a unit symbol absent from a registry.
	ErrUnknownUnit = errors.New("unknown unit")
	// ErrMissingUnit reports input without a unit when no default is configured.
	ErrMissingUnit = errors.New("missing unit")
)

// Error adds the original input and byte position to a parser error.
type Error struct {
	Input string
	Pos   int
	Err   error
}

func (e *Error) Error() string {
	return fmt.Sprintf("parse: %v at position %d in %q", e.Err, e.Pos, e.Input)
}

func (e *Error) Unwrap() error { return e.Err }

// UnitSet maps accepted symbols to units for one dimension.
//
// Deprecated: use Registry and the dimension-specific accessor functions. It
// remains available for source compatibility with the v1 parser API.
type UnitSet[D quant.Dimension] map[string]quant.Unit[D]

// Registry is an immutable, per-dimension unit registry. With and WithDefault
// return copies; parsing never mutates a registry.
type Registry[D quant.Dimension] struct {
	units       map[string]quant.Unit[D]
	aliases     map[string]quant.Unit[D]
	defaultUnit quant.Unit[D]
	hasDefault  bool
}

func NewRegistry[D quant.Dimension]() Registry[D] {
	return Registry[D]{units: make(map[string]quant.Unit[D]), aliases: make(map[string]quant.Unit[D])}
}

// With returns a copy containing symbol and any aliases for unit. Duplicate
// symbols panic so registry construction cannot silently select the wrong unit.
func (r Registry[D]) With(symbol string, unit quant.Unit[D], aliases ...string) Registry[D] {
	next := r.clone()
	next.addSymbol(symbol, unit)
	for _, alias := range aliases {
		next.addAlias(alias, unit)
	}
	return next
}

// WithDefault returns a copy that uses unit when Parse receives no symbol.
func (r Registry[D]) WithDefault(unit quant.Unit[D]) Registry[D] {
	next := r.clone()
	next.defaultUnit = unit
	next.hasDefault = true
	return next
}

// Parse parses a scalar followed by a unit symbol. Both "5 km" and "5km" are
// accepted. A missing symbol is accepted only when WithDefault was used.
func (r Registry[D]) Parse(input string) (quant.Quantity[D], error) {
	value, symbol, pos, err := split(input)
	if err != nil {
		return quant.Quantity[D]{}, err
	}
	if symbol == "" {
		if !r.hasDefault {
			return quant.Quantity[D]{}, &Error{Input: input, Pos: pos, Err: ErrMissingUnit}
		}
		return quant.New[D](value, r.defaultUnit), nil
	}
	normalized := normalizeSymbol(symbol)
	u, ok := r.units[normalized]
	if !ok {
		u, ok = r.aliases[strings.ToLower(normalized)]
	}
	if !ok {
		return quant.Quantity[D]{}, &Error{Input: input, Pos: pos, Err: ErrUnknownUnit}
	}
	return quant.New[D](value, u), nil
}

func (r Registry[D]) clone() Registry[D] {
	next := r
	next.units = make(map[string]quant.Unit[D], len(r.units))
	next.aliases = make(map[string]quant.Unit[D], len(r.aliases))
	for symbol, unit := range r.units {
		next.units[symbol] = unit
	}
	for alias, unit := range r.aliases {
		next.aliases[alias] = unit
	}
	return next
}

func (r *Registry[D]) addSymbol(symbol string, unit quant.Unit[D]) {
	symbol = normalizeSymbol(symbol)
	if strings.TrimSpace(symbol) == "" {
		panic("parse: registry symbols cannot be empty")
	}
	if _, exists := r.units[symbol]; exists {
		panic(fmt.Sprintf("parse: duplicate unit symbol %q", symbol))
	}
	if _, exists := r.aliases[strings.ToLower(symbol)]; exists {
		panic(fmt.Sprintf("parse: duplicate unit symbol %q", symbol))
	}
	r.units[symbol] = unit
}

func (r *Registry[D]) addAlias(alias string, unit quant.Unit[D]) {
	alias = strings.ToLower(normalizeSymbol(alias))
	if strings.TrimSpace(alias) == "" {
		panic("parse: registry aliases cannot be empty")
	}
	if _, exists := r.units[alias]; exists {
		panic(fmt.Sprintf("parse: duplicate unit alias %q", alias))
	}
	if _, exists := r.aliases[alias]; exists {
		panic(fmt.Sprintf("parse: duplicate unit alias %q", alias))
	}
	r.aliases[alias] = unit
}

// Parse parses a scalar and looks up its symbol in a legacy UnitSet. The empty
// symbol behavior is retained for compatibility with the original API.
func Parse[D quant.Dimension](input string, units UnitSet[D]) (quant.Quantity[D], error) {
	value, symbol, pos, err := split(input)
	if err != nil {
		return quant.Quantity[D]{}, err
	}
	symbol = normalizeSymbol(symbol)
	u, ok := units[symbol]
	if !ok {
		kind := ErrUnknownUnit
		if symbol == "" {
			kind = ErrMissingUnit
		}
		return quant.Quantity[D]{}, &Error{Input: input, Pos: pos, Err: kind}
	}
	return quant.New[D](value, u), nil
}

func ParseLength(input string) (quant.Quantity[quant.Length], error) {
	return Parse(input, LengthUnits)
}

func ParseMass(input string) (quant.Quantity[quant.Mass], error) {
	return Parse(input, MassUnits)
}

func ParseTime(input string) (quant.Quantity[quant.Time], error) {
	return Parse(input, TimeUnits)
}

func ParseTemperature(input string) (quant.Quantity[quant.Temperature], error) {
	return Parse(input, TemperatureUnits)
}

// Lengths returns an immutable length registry with no implicit default.
func Lengths() Registry[quant.Length] {
	return NewRegistry[quant.Length]().
		With(quant.UnitSymbol(quant.Meter), quant.Meter, "meter", "meters").
		With(quant.UnitSymbol(quant.Kilometer), quant.Kilometer, "kilometer", "kilometers").
		With(quant.UnitSymbol(quant.Centimeter), quant.Centimeter, "centimeter", "centimeters").
		With(quant.UnitSymbol(quant.Millimeter), quant.Millimeter, "millimeter", "millimeters").
		With(quant.UnitSymbol(quant.Mile), quant.Mile, "mile", "miles").
		With(quant.UnitSymbol(quant.Foot), quant.Foot, "foot", "feet").
		With(quant.UnitSymbol(quant.Inch), quant.Inch, "inch", "inches")
}

// Masses returns an immutable mass registry with no implicit default.
func Masses() Registry[quant.Mass] {
	return NewRegistry[quant.Mass]().
		With(quant.UnitSymbol(quant.Kilogram), quant.Kilogram, "kilogram", "kilograms").
		With(quant.UnitSymbol(quant.Gram), quant.Gram, "gram", "grams").
		With(quant.UnitSymbol(quant.Milligram), quant.Milligram, "milligram", "milligrams").
		With(quant.UnitSymbol(quant.Pound), quant.Pound, "pound", "pounds").
		With(quant.UnitSymbol(quant.Ounce), quant.Ounce, "ounce", "ounces")
}

// Times returns an immutable time registry with no implicit default.
func Times() Registry[quant.Time] {
	return NewRegistry[quant.Time]().
		With(quant.UnitSymbol(quant.Second), quant.Second).
		With(quant.UnitSymbol(quant.Millisecond), quant.Millisecond).
		With(quant.UnitSymbol(quant.Minute), quant.Minute).
		With(quant.UnitSymbol(quant.Hour), quant.Hour).
		With(quant.UnitSymbol(quant.Day), quant.Day)
}

// Temperatures returns an immutable temperature registry with no implicit
// default.
func Temperatures() Registry[quant.Temperature] {
	return NewRegistry[quant.Temperature]().
		With(quant.UnitSymbol(quant.Kelvin), quant.Kelvin).
		With(quant.UnitSymbol(quant.Celsius), quant.Celsius).
		With(quant.UnitSymbol(quant.Fahrenheit), quant.Fahrenheit)
}

// Deprecated parser maps retain the original default-unit behavior. Their
// keys are sourced from quant.UnitSymbol so they cannot drift from formatting.
// LengthUnits is the legacy length symbol map with Meter as its default.
//
// Deprecated: use Lengths, optionally followed by WithDefault.
var LengthUnits = UnitSet[quant.Length]{
	"": quant.Meter, quant.UnitSymbol(quant.Meter): quant.Meter,
	quant.UnitSymbol(quant.Kilometer): quant.Kilometer, quant.UnitSymbol(quant.Centimeter): quant.Centimeter,
	quant.UnitSymbol(quant.Millimeter): quant.Millimeter, quant.UnitSymbol(quant.Mile): quant.Mile,
	quant.UnitSymbol(quant.Foot): quant.Foot, quant.UnitSymbol(quant.Inch): quant.Inch,
}

// MassUnits is the legacy mass symbol map with Kilogram as its default.
//
// Deprecated: use Masses, optionally followed by WithDefault.
var MassUnits = UnitSet[quant.Mass]{
	"": quant.Kilogram, quant.UnitSymbol(quant.Kilogram): quant.Kilogram,
	quant.UnitSymbol(quant.Gram): quant.Gram, quant.UnitSymbol(quant.Milligram): quant.Milligram,
	quant.UnitSymbol(quant.Pound): quant.Pound, quant.UnitSymbol(quant.Ounce): quant.Ounce,
}

// TimeUnits is the legacy time symbol map with Second as its default.
//
// Deprecated: use Times, optionally followed by WithDefault.
var TimeUnits = UnitSet[quant.Time]{
	"": quant.Second, quant.UnitSymbol(quant.Second): quant.Second,
	quant.UnitSymbol(quant.Millisecond): quant.Millisecond, quant.UnitSymbol(quant.Minute): quant.Minute,
	quant.UnitSymbol(quant.Hour): quant.Hour, quant.UnitSymbol(quant.Day): quant.Day,
}

// TemperatureUnits is the legacy temperature symbol map with Kelvin as its default.
//
// Deprecated: use Temperatures, optionally followed by WithDefault.
var TemperatureUnits = UnitSet[quant.Temperature]{
	"": quant.Kelvin, quant.UnitSymbol(quant.Kelvin): quant.Kelvin,
	quant.UnitSymbol(quant.Celsius): quant.Celsius, quant.UnitSymbol(quant.Fahrenheit): quant.Fahrenheit,
}

func split(input string) (float64, string, int, error) {
	s := strings.TrimSpace(input)
	base := len(input) - len(strings.TrimLeft(input, " \t\r\n"))
	if s == "" {
		return 0, "", base, &Error{Input: input, Pos: base, Err: ErrSyntax}
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
		return 0, "", base, &Error{Input: input, Pos: base, Err: ErrSyntax}
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
			return 0, "", base + end, &Error{Input: input, Pos: base + end, Err: ErrSyntax}
		}
	}

	value, err := strconv.ParseFloat(s[:end], 64)
	if err != nil {
		return 0, "", base, &Error{Input: input, Pos: base, Err: ErrSyntax}
	}
	symbol := strings.TrimSpace(s[end:])
	if strings.ContainsAny(symbol, " \t\r\n") {
		return 0, "", base + end, &Error{Input: input, Pos: base + end, Err: ErrSyntax}
	}
	return value, symbol, base + end, nil
}

func normalizeSymbol(symbol string) string {
	// Normalize common Unicode spellings to the ASCII symbols used by the
	// built-in registries. Canonical symbol matching remains case-sensitive.
	symbol = strings.ReplaceAll(symbol, "µ", "u")
	symbol = strings.ReplaceAll(symbol, "μ", "u")
	symbol = strings.ReplaceAll(symbol, "°C", "C")
	symbol = strings.ReplaceAll(symbol, "°F", "F")
	symbol = strings.ReplaceAll(symbol, "°K", "K")
	symbol = strings.ReplaceAll(symbol, "Ω", "Ohm")
	symbol = strings.ReplaceAll(symbol, "\u00a0", "")
	symbol = strings.ReplaceAll(symbol, "\u2009", "")
	symbol = strings.ReplaceAll(symbol, "\u202f", "")
	return symbol
}
