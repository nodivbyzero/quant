package quant

// Digital is the dimension marker for digital information quantities.
type Digital struct{}

// Digital units.
var (
	Bit      = scaleUnit[Digital]{factor: 1}
	Kilobit  = scaleUnit[Digital]{factor: 1e3}
	Megabit  = scaleUnit[Digital]{factor: 1e6}
	Gigabit  = scaleUnit[Digital]{factor: 1e9}
	Terabit  = scaleUnit[Digital]{factor: 1e12}
	Byte     = scaleUnit[Digital]{factor: 8}
	Kilobyte = scaleUnit[Digital]{factor: 8e3}
	Megabyte = scaleUnit[Digital]{factor: 8e6}
	Gigabyte = scaleUnit[Digital]{factor: 8e9}
	Terabyte = scaleUnit[Digital]{factor: 8e12}
	Petabyte = scaleUnit[Digital]{factor: 8e15}
	Kibibyte = scaleUnit[Digital]{factor: 8 * 1024}
	Mebibyte = scaleUnit[Digital]{factor: 8 * 1024 * 1024}
	Gibibyte = scaleUnit[Digital]{factor: 8 * 1024 * 1024 * 1024}
	Tebibyte = scaleUnit[Digital]{factor: 8 * 1024 * 1024 * 1024 * 1024}
)
