package quant

// Resistance is the dimension marker for electrical resistance quantities.
type Resistance struct{}

// Capacitance is the dimension marker for electrical capacitance quantities.
type Capacitance struct{}

// Inductance is the dimension marker for electrical inductance quantities.
type Inductance struct{}

// ElectricField is the dimension marker for electric field strength quantities.
type ElectricField struct{}

// Resistance units.
var (
	Ohm      = scaleUnit[Resistance]{factor: 1}
	Milliohm = scaleUnit[Resistance]{factor: 1e-3}
	Kiloohm  = scaleUnit[Resistance]{factor: 1e3}
	Megaohm  = scaleUnit[Resistance]{factor: 1e6}
)

// Capacitance units.
var (
	Farad      = scaleUnit[Capacitance]{factor: 1}
	Microfarad = scaleUnit[Capacitance]{factor: 1e-6}
	Nanofarad  = scaleUnit[Capacitance]{factor: 1e-9}
	Picofarad  = scaleUnit[Capacitance]{factor: 1e-12}
)

// Inductance units.
var (
	Henry      = scaleUnit[Inductance]{factor: 1}
	Millihenry = scaleUnit[Inductance]{factor: 1e-3}
	Microhenry = scaleUnit[Inductance]{factor: 1e-6}
)

// Electric field units.
var (
	VoltPerMeter     = scaleUnit[ElectricField]{factor: 1}
	NewtonPerCoulomb = scaleUnit[ElectricField]{factor: 1}
)
