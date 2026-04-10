package quant

// Mul is a phantom dimension representing the product of two dimensions.
type Mul[A, B any] struct{}

// Div is a phantom dimension representing the quotient of two dimensions.
type Div[A, B any] struct{}

// Div divides a length quantity by a time quantity and returns speed.
func (q Quantity[Length]) Div(other Quantity[Time]) Quantity[Speed] {
	return Quantity[Speed]{value: q.value / other.value}
}

// DivSpeed divides a speed quantity by a time quantity and returns acceleration.
func (q Quantity[Speed]) DivSpeed(other Quantity[Time]) Quantity[Acceleration] {
	return Quantity[Acceleration]{value: q.value / other.value}
}
