package quant

// Energy is the dimension marker for energy quantities.
type Energy struct{}

// ReactiveEnergy is the dimension marker for reactive energy quantities.
type ReactiveEnergy struct{}

// Energy units.
var (
	Joule         = scaleUnit[Energy]{factor: 1}
	Kilojoule     = scaleUnit[Energy]{factor: 1e3}
	Megajoule     = scaleUnit[Energy]{factor: 1e6}
	Gigajoule     = scaleUnit[Energy]{factor: 1e9}
	WattSecond    = scaleUnit[Energy]{factor: 1}
	WattMinute    = scaleUnit[Energy]{factor: 60}
	MilliwattHour = scaleUnit[Energy]{factor: 3.6}
	WattHour      = scaleUnit[Energy]{factor: 3600}
	KilowattHour  = scaleUnit[Energy]{factor: 3.6e6}
	MegawattHour  = scaleUnit[Energy]{factor: 3.6e9}
	GigawattHour  = scaleUnit[Energy]{factor: 3.6e12}
)

// Reactive energy units.
var (
	VoltAmpereReactiveHour      = scaleUnit[ReactiveEnergy]{factor: 1}
	MilliVoltAmpereReactiveHour = scaleUnit[ReactiveEnergy]{factor: 1e-3}
	KiloVoltAmpereReactiveHour  = scaleUnit[ReactiveEnergy]{factor: 1e3}
	MegaVoltAmpereReactiveHour  = scaleUnit[ReactiveEnergy]{factor: 1e6}
	GigaVoltAmpereReactiveHour  = scaleUnit[ReactiveEnergy]{factor: 1e9}
)
