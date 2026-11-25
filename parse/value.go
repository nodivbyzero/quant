package parse

import (
	"encoding"
	"errors"
	"flag"

	"github.com/nodivbyzero/quant"
)

// ErrNoRegistry reports use of a zero Value without a configured registry.
var ErrNoRegistry = errors.New("parse: value has no registry")

// Value is a typed, unit-aware boundary value. It implements
// encoding.TextMarshaler, encoding.TextUnmarshaler, and flag.Value.
//
// Value deliberately belongs to parse rather than quant: accepting unit names
// is an application-boundary decision, while the core quantity remains a
// numeric value in its base unit.
type Value[D quant.Dimension] struct {
	registry Registry[D]
	quantity quant.Quantity[D]
	text     string
	set      bool
}

// NewValue creates a boundary value backed by registry.
func NewValue[D quant.Dimension](registry Registry[D]) *Value[D] {
	return &Value[D]{registry: registry}
}

// Set parses and stores a unit-bearing value. On error, the previous value is
// retained.
func (v *Value[D]) Set(text string) error {
	if v == nil {
		return ErrNoRegistry
	}
	if v.registry.units == nil && v.registry.aliases == nil && !v.registry.hasDefault {
		return ErrNoRegistry
	}
	q, err := v.registry.Parse(text)
	if err != nil {
		return err
	}
	v.quantity = q
	v.text = text
	v.set = true
	return nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (v *Value[D]) UnmarshalText(text []byte) error {
	return v.Set(string(text))
}

// MarshalText implements encoding.TextMarshaler. Values parsed from text
// retain that parseable representation, including its unit. A value assigned
// with SetQuantity uses the quantity's base-unit String representation.
func (v Value[D]) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

// SetQuantity stores q and uses its base-unit representation when marshaled.
func (v *Value[D]) SetQuantity(q quant.Quantity[D]) {
	v.quantity = q
	v.text = q.String()
	v.set = true
}

// Quantity returns the parsed quantity and whether a value has been set.
func (v Value[D]) Quantity() (quant.Quantity[D], bool) {
	return v.quantity, v.set
}

// String implements flag.Value. It returns a parseable representation.
func (v Value[D]) String() string {
	if !v.set {
		return ""
	}
	return v.text
}

var (
	_ encoding.TextMarshaler   = (*Value[quant.Length])(nil)
	_ encoding.TextUnmarshaler = (*Value[quant.Length])(nil)
	_ flag.Value               = (*Value[quant.Length])(nil)
)
