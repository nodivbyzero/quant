package quant

// DataRate is the dimension marker for digital throughput quantities.
type DataRate struct{}

// Data rate units.
var (
	BitPerSecond      = scaleUnit[DataRate]{factor: 1}
	KilobitPerSecond  = scaleUnit[DataRate]{factor: 1e3}
	MegabitPerSecond  = scaleUnit[DataRate]{factor: 1e6}
	GigabitPerSecond  = scaleUnit[DataRate]{factor: 1e9}
	TerabitPerSecond  = scaleUnit[DataRate]{factor: 1e12}
	BytePerSecond     = scaleUnit[DataRate]{factor: 8}
	KilobytePerSecond = scaleUnit[DataRate]{factor: 8e3}
	MegabytePerSecond = scaleUnit[DataRate]{factor: 8e6}
	GigabytePerSecond = scaleUnit[DataRate]{factor: 8e9}
	TerabytePerSecond = scaleUnit[DataRate]{factor: 8e12}
	KibibytePerSecond = scaleUnit[DataRate]{factor: 8 * 1024}
	MebibytePerSecond = scaleUnit[DataRate]{factor: 8 * 1024 * 1024}
	GibibytePerSecond = scaleUnit[DataRate]{factor: 8 * 1024 * 1024 * 1024}
	TebibytePerSecond = scaleUnit[DataRate]{factor: 8 * 1024 * 1024 * 1024 * 1024}
)
