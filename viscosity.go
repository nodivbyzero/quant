package quant

// DynamicViscosity is the dimension marker for dynamic viscosity quantities.
type DynamicViscosity struct{}

// KinematicViscosity is the dimension marker for kinematic viscosity quantities.
type KinematicViscosity struct{}

// Dynamic viscosity units.
var (
	PascalSecond = scaleUnit[DynamicViscosity]{factor: 1}
	Poise        = scaleUnit[DynamicViscosity]{factor: 0.1}
	Centipoise   = scaleUnit[DynamicViscosity]{factor: 0.001}
)

// Kinematic viscosity units.
var (
	SquareMeterPerSecond = scaleUnit[KinematicViscosity]{factor: 1}
	Stokes               = scaleUnit[KinematicViscosity]{factor: 1e-4}
	Centistokes          = scaleUnit[KinematicViscosity]{factor: 1e-6}
)
