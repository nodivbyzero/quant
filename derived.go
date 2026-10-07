package quant

// Mul is a phantom dimension representing the product of two dimensions.
type Mul[A, B Dimension] struct{}

// Div is a phantom dimension representing the quotient of two dimensions.
type Div[A, B Dimension] struct{}

// Div divides a length quantity by a time quantity and returns speed.
func (q Quantity[Length]) Div(other Quantity[Time]) Quantity[Speed] {
	return Quantity[Speed]{value: q.value / other.value}
}

// DivSpeed divides a speed quantity by a time quantity and returns acceleration.
func (q Quantity[Speed]) DivSpeed(other Quantity[Time]) Quantity[Acceleration] {
	return Quantity[Acceleration]{value: q.value / other.value}
}

// MulLength multiplies two lengths to produce an area.
func (q Quantity[Length]) MulLength(other Quantity[Length]) Quantity[Area] {
	return Quantity[Area]{value: q.value * other.value}
}

// MulTime multiplies power by time to produce energy.
func (q Quantity[Power]) MulTime(other Quantity[Time]) Quantity[Energy] {
	return Quantity[Energy]{value: q.value * other.value}
}

// MulCurrent multiplies voltage by current to produce active power.
func (q Quantity[Voltage]) MulCurrent(other Quantity[Current]) Quantity[Power] {
	return Quantity[Power]{value: q.value * other.value}
}

// MulResistance multiplies current by resistance to produce voltage.
func (q Quantity[Current]) MulResistance(other Quantity[Resistance]) Quantity[Voltage] {
	return Quantity[Voltage]{value: q.value * other.value}
}
