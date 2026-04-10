package quant

// VolumeFlowRate is the dimension marker for volumetric flow rate quantities.
type VolumeFlowRate struct{}

// Volume flow rate units.
var (
	CubicMillimeterPerSecond = scaleUnit[VolumeFlowRate]{factor: 1e-9}
	CubicCentimeterPerSecond = scaleUnit[VolumeFlowRate]{factor: 1e-6}
	MilliliterPerSecond      = scaleUnit[VolumeFlowRate]{factor: 1e-6}
	CentiliterPerSecond      = scaleUnit[VolumeFlowRate]{factor: 1e-5}
	DeciliterPerSecond       = scaleUnit[VolumeFlowRate]{factor: 1e-4}
	LiterPerSecond           = scaleUnit[VolumeFlowRate]{factor: 1e-3}
	LiterPerMinute           = scaleUnit[VolumeFlowRate]{factor: 1e-3 / 60}
	LiterPerHour             = scaleUnit[VolumeFlowRate]{factor: 1e-3 / 3600}
	KiloliterPerSecond       = scaleUnit[VolumeFlowRate]{factor: 1}
	KiloliterPerMinute       = scaleUnit[VolumeFlowRate]{factor: 1.0 / 60.0}
	KiloliterPerHour         = scaleUnit[VolumeFlowRate]{factor: 1.0 / 3600.0}
	CubicMeterPerSecond      = scaleUnit[VolumeFlowRate]{factor: 1}
	CubicMeterPerMinute      = scaleUnit[VolumeFlowRate]{factor: 1.0 / 60.0}
	CubicMeterPerHour        = scaleUnit[VolumeFlowRate]{factor: 1.0 / 3600.0}
	CubicKilometerPerSecond  = scaleUnit[VolumeFlowRate]{factor: 1e9}
	TeaspoonPerSecond        = scaleUnit[VolumeFlowRate]{factor: 4.92892159375e-6}
	TablespoonPerSecond      = scaleUnit[VolumeFlowRate]{factor: 1.478676478125e-5}
	CubicInchPerSecond       = scaleUnit[VolumeFlowRate]{factor: 1.6387064e-5}
	CubicInchPerMinute       = scaleUnit[VolumeFlowRate]{factor: 1.6387064e-5 / 60}
	CubicInchPerHour         = scaleUnit[VolumeFlowRate]{factor: 1.6387064e-5 / 3600}
	FluidOuncePerSecond      = scaleUnit[VolumeFlowRate]{factor: 2.95735295625e-5}
	FluidOuncePerMinute      = scaleUnit[VolumeFlowRate]{factor: 2.95735295625e-5 / 60}
	FluidOuncePerHour        = scaleUnit[VolumeFlowRate]{factor: 2.95735295625e-5 / 3600}
	CupPerSecond             = scaleUnit[VolumeFlowRate]{factor: 2.365882365e-4}
	PintPerSecond            = scaleUnit[VolumeFlowRate]{factor: 4.73176473e-4}
	PintPerMinute            = scaleUnit[VolumeFlowRate]{factor: 4.73176473e-4 / 60}
	PintPerHour              = scaleUnit[VolumeFlowRate]{factor: 4.73176473e-4 / 3600}
	QuartPerSecond           = scaleUnit[VolumeFlowRate]{factor: 9.46352946e-4}
	GallonPerSecond          = scaleUnit[VolumeFlowRate]{factor: 3.785411784e-3}
	GallonPerMinute          = scaleUnit[VolumeFlowRate]{factor: 3.785411784e-3 / 60}
	GallonPerHour            = scaleUnit[VolumeFlowRate]{factor: 3.785411784e-3 / 3600}
	CubicFootPerSecond       = scaleUnit[VolumeFlowRate]{factor: 0.028316846592}
	CubicFootPerMinute       = scaleUnit[VolumeFlowRate]{factor: 0.028316846592 / 60}
	CubicFootPerHour         = scaleUnit[VolumeFlowRate]{factor: 0.028316846592 / 3600}
	CubicYardPerSecond       = scaleUnit[VolumeFlowRate]{factor: 0.764554857984}
	CubicYardPerMinute       = scaleUnit[VolumeFlowRate]{factor: 0.764554857984 / 60}
	CubicYardPerHour         = scaleUnit[VolumeFlowRate]{factor: 0.764554857984 / 3600}
)
