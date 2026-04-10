package quant

// builder is a fluent helper for constructing quantities from unit values.
type builder[D any] struct {
	value float64
}

// Amount starts a fluent quantity construction.
func Amount[D any](v float64) builder[D] {
	return builder[D]{value: v}
}

// From constructs a quantity from the builder value expressed in the given unit.
func (b builder[D]) From(u Unit[D]) Quantity[D] {
	return New[D](b.value, u)
}
