# quant

`quant` is a small, type-safe Go library for physical unit conversions using generics.

It uses phantom types to model dimensions like length, area, mass, volume, flow, time, temperature, electrical values, and more, so incompatible units cannot be mixed by accident at compile time.

## Features

- Compile-time dimension safety with `Quantity[D]`
- Zero-dependency, idiomatic Go API
- Internal storage in base units for predictable conversions
- Fluent construction with `Amount[D](v).From(unit)`
- Convenience constructors like `quant.Pounds(10)` and `quant.Kilometers(5)`
- Safe arithmetic on matching dimensions only
- Base-unit `String()` formatting and explicit unit-aware `Format(...)`
- Derived dimensions with type-safe division for built-in combinations like length/time
- JSON marshaling based on the scalar base-unit value
- JSON scalar output rounded to 15 significant digits, plus `MarshalJSONIn` for a chosen unit
- Text marshaling and `database/sql` interop helpers
- `time.Duration` interop with `FromDuration` and `ToDuration`
- Support for affine temperature conversions
- Separate temperature-delta quantities for safe temperature differences

## Installation

```bash
	go get github.com/nodivbyzero/quant
```

## Quick Start

```go
package main

import (
	"fmt"

	"github.com/nodivbyzero/quant"
)

func main() {
	kg := quant.Amount[quant.Mass](10).From(quant.Pound).To(quant.Kilogram)
	fmt.Printf("%.6f\n", kg)
}
```

Output:

```text
4.535924
```

## Core API

```go
type Quantity[D any] struct

func New[D any, U Unit[D]](v float64, u U) Quantity[D]
func From[D any](v float64, u BuiltinUnit[D]) Quantity[D]

func (q Quantity[D]) To(u Unit[D]) float64
func (q Quantity[D]) Value(u Unit[D]) float64
func (q Quantity[D]) Format(u Unit[D]) string
func (q Quantity[D]) MarshalJSON() ([]byte, error)
func (q *Quantity[D]) UnmarshalJSON(data []byte) error
func (q Quantity[D]) MarshalText() ([]byte, error)
func (q *Quantity[D]) UnmarshalText(text []byte) error
func (q Quantity[D]) Add(other Quantity[D]) Quantity[D]
func (q Quantity[D]) Sub(other Quantity[D]) Quantity[D]
func (q Quantity[D]) Mul(scalar float64) Quantity[D]
func (q Quantity[D]) DivScalar(scalar float64) Quantity[D]
func (q Quantity[D]) Neg() Quantity[D]
func (q Quantity[D]) Abs() Quantity[D]
func (q Quantity[D]) LessThan(other Quantity[D]) bool
func (q Quantity[D]) GreaterThan(other Quantity[D]) bool
func (q Quantity[D]) EqualWithin(other, tolerance Quantity[D]) bool
func (q Quantity[D]) IsZero() bool
func (q Quantity[D]) IsPositive() bool
func (q Quantity[D]) IsNegative() bool
```

`Amount[D](v).From(unit)` is also available as a fluent constructor.

`From(v, unit)` is a terser constructor for built-in units:

```go
mass := quant.From(10, quant.Pound)
```

Convenience constructors are also available for direct creation from specific units, for example:

```go
mass := quant.Pounds(10)
distance := quant.Kilometers(5)
temp := quant.DegreesCelsius(21)
```

`Quantity[D]` also implements `fmt.Stringer` in the base unit for its dimension, and supports explicit formatting in a chosen unit:

```go
fmt.Println(quant.Pounds(10))                  // 4.5359237 kg
fmt.Println(quant.Pounds(10).Format(quant.Pound)) // 10 lb
```

JSON marshaling stores the scalar value in the base unit for the dimension and
uses 15 significant digits to avoid exposing binary floating-point noise:

```go
data, _ := json.Marshal(quant.Pounds(10))
fmt.Println(string(data)) // 4.5359237
```

For an API whose contract is a particular unit, use the explicit helper:

```go
data, _ := quant.Pounds(10).MarshalJSONIn(quant.Pound)
fmt.Println(string(data)) // 10
```

The default scalar form is intentionally compact, but it does not carry a
unit label. Use a surrounding object or `MarshalJSONIn` when the consumer
cannot know the base unit from the schema.

Text marshaling follows the same base-unit rule:

```go
text, _ := quant.Pounds(10).MarshalText()
fmt.Println(string(text)) // 4.5359237
```

For `database/sql`, use the adapter helper:

```go
value, _ := quant.SQL(quant.Kilograms(5)).Value()
fmt.Println(value) // 5
```

All quantities are stored internally in a base unit for their dimension:

- Acceleration: meter per second squared
- Acidity: pH
- Absorbed dose: gray
- Angle: radian
- Angular acceleration: radian per second squared
- Angular velocity: radian per second
- Area: square meter
- Apparent power: volt-ampere
- Capacitance: farad
- Catalytic activity: katal
- Charge: coulomb
- Concentration: mole per cubic meter
- Current: ampere
- Data rate: bit per second
- Density: kilogram per cubic meter
- Digital: bit
- Dynamic viscosity: pascal-second
- Electrical conductivity: siemens per meter
- Electric field: volt per meter
- Energy: joule
- Equivalent dose: sievert
- Force: newton
- Frequency: hertz
- Illuminance: lux
- Inductance: henry
- Kinematic viscosity: square meter per second
- Length: meter
- Luminous flux: lumen
- Luminous intensity: candela
- Magnetic flux: weber
- Magnetic flux density: tesla
- Mass: kilogram
- Molality: mole per kilogram
- Momentum: kilogram-meter per second
- Pace: second per meter
- Parts-per: ratio
- Pieces: piece
- Power: watt
- Pressure: pascal
- Radioactivity: becquerel
- Reactive energy: volt-ampere reactive hour
- Reactive power: volt-ampere reactive
- Resistance: ohm
- Speed: meter per second
- Specific heat capacity: joule per kilogram kelvin
- Surface tension: newton per meter
- Temperature: kelvin
- Temperature delta: kelvin delta
- Thermal conductivity: watt per meter kelvin
- Time: second
- Torque: newton-meter
- Voltage: volt
- Volume: cubic meter
- Volume flow rate: cubic meter per second

## Examples By Dimension

- Acceleration: `ms2 := quant.GForces(1).To(quant.MeterPerSecondSquared)`
- Acidity: `poh := quant.PHValue(3).To(quant.POH)`
- Absorbed dose: `gy := quant.Rads(100).To(quant.Gray)`
- Angle: `rad := quant.Degrees(180).To(quant.Radian)`
- Angular acceleration: `radS2 := quant.AngularDegreesPerSecondSquared(180).To(quant.AngularRadianPerSecondSquared)`
- Angular velocity: `radS := quant.AngularRevolutionsPerMinute(60).To(quant.AngularRadianPerSecond)`
- Area: `m2 := quant.Acres(1).To(quant.SquareMeter)`
- Apparent power: `va := quant.MegaVoltAmperes(1).To(quant.VoltAmpere)`
- Capacitance: `uf := quant.Nanofarads(1000).To(quant.Microfarad)`
- Catalytic activity: `kat := quant.Katals(1).To(quant.Katal)`
- Charge: `mc := quant.Coulombs(1).To(quant.Millicoulomb)`
- Concentration: `molm3 := quant.Molars(1).To(quant.MolePerCubicMeter)`
- Current: `a := quant.Kiloamperes(1).To(quant.Ampere)`
- Data rate: `mbps := quant.MegabytesPerSecond(1).To(quant.MegabitPerSecond)`
- Density: `kgm3 := quant.GramsPerCubicCentimeter(1).To(quant.KilogramPerCubicMeter)`
- Digital: `bytes := quant.Kibibytes(1).To(quant.Byte)`
- Dynamic viscosity: `cp := quant.Poises(1).To(quant.Centipoise)`
- Electrical conductivity: `sm := quant.MillisiemensPerCentimeters(1).To(quant.SiemensPerMeter)`
- Electric field: `nc := quant.VoltsPerMeter(1).To(quant.NewtonPerCoulomb)`
- Energy: `j := quant.KilowattHours(1).To(quant.Joule)`
- Equivalent dose: `sv := quant.Rems(100).To(quant.Sievert)`
- Force: `n := quant.KilogramsForce(1).To(quant.Newton)`
- Frequency: `hz := quant.RevolutionsPerMinute(60).To(quant.Hertz)`
- Illuminance: `lx := quant.FootCandles(1).To(quant.Lux)`
- Inductance: `mh := quant.Microhenrys(1000).To(quant.Millihenry)`
- Kinematic viscosity: `cst := quant.StokesValues(1).To(quant.Centistokes)`
- Length: `km := quant.Miles(3).To(quant.Kilometer)`
- Luminous flux: `lm := quant.Lumens(800).To(quant.Lumen)`
- Luminous intensity: `cd := quant.Candelas(1).To(quant.Candela)`
- Magnetic flux: `mx := quant.Webers(1).To(quant.Maxwell)`
- Magnetic flux density: `gauss := quant.Teslas(1).To(quant.Gauss)`
- Mass: `kg := quant.Pounds(10).To(quant.Kilogram)`
- Molality: `molkg := quant.Molals(1).To(quant.MolePerKilogram)`
- Momentum: `ns := quant.KilogramMetersPerSecond(1).To(quant.NewtonSecond)`
- Pace: `spm := quant.MinutesPerKilometer(5).To(quant.SecondPerMeter)`
- Parts-per: `ppb := quant.PartsPerMillion(1).To(quant.PPB)`
- Pieces: `pcs := quant.Dozens(1).To(quant.Piece)`
- Power: `w := quant.HorsepowerValues(1).To(quant.Watt)`
- Pressure: `pa := quant.Bars(1).To(quant.Pascal)`
- Radioactivity: `bq := quant.Curies(1).To(quant.Becquerel)`
- Reactive energy: `varh := quant.MegaVoltAmpereReactiveHours(1).To(quant.KiloVoltAmpereReactiveHour)`
- Reactive power: `vars := quant.MegaVoltAmpereReactives(1).To(quant.KiloVoltAmpereReactive)`
- Resistance: `ohm := quant.Kiloohms(1).To(quant.Ohm)`
- Speed: `kmh := quant.MetersPerSecond(10).To(quant.KilometerPerHour)`
- Specific heat capacity: `jkgk := quant.JoulesPerKilogramKelvin(1).To(quant.JoulePerKilogramKelvin)`
- Surface tension: `nm := quant.DynesPerCentimeter(1).To(quant.NewtonPerMeter)`
- Temperature: `k := quant.DegreesCelsius(25).To(quant.Kelvin)`
- Temperature delta: `dk := quant.FahrenheitDeltas(18).To(quant.KelvinDelta)`
- Thermal conductivity: `wmk := quant.WattsPerMeterKelvin(1).To(quant.WattPerMeterKelvin)`
- Time: `years := quant.Decades(1).To(quant.Year)`
- Torque: `nm := quant.PoundForceFeet(1).To(quant.NewtonMeter)`
- Voltage: `v := quant.Kilovolts(1).To(quant.Volt)`
- Volume: `l := quant.Gallons(1).To(quant.Liter)`
- Volume flow rate: `ls := quant.LitersPerMinute(60).To(quant.LiterPerSecond)`

### Arithmetic

```go
total := quant.Meters(750).Add(quant.Kilometers(1.25))
fmt.Println(total.To(quant.Meter)) // 2000
```

### Scalar Operations

```go
doubled := quant.Meters(5).Mul(2)
halved := quant.Meters(5).DivScalar(2)
neg := quant.Meters(5).Neg()
abs := quant.Meters(-5).Abs()
```

### Comparisons

```go
short := quant.Meters(100)
long := quant.Kilometers(1)

fmt.Println(short.LessThan(long)) // true
fmt.Println(quant.Min(short, long).To(quant.Meter)) // 100
fmt.Println(long.EqualWithin(quant.Meters(1000.0001), quant.Millimeters(1))) // true
```

### Derived units

```go
distance := quant.Kilometers(5)
duration := quant.Minutes(30)
speed := distance.Div(duration)

fmt.Println(speed.To(quant.KilometerPerHour)) // 10

area := quant.Meters(3).MulLength(quant.Meters(4))
energy := quant.Watts(100).MulTime(quant.Seconds(2))
power := quant.Volts(12).MulCurrent(quant.Amperes(2))
voltage := quant.Amperes(2).MulResistance(quant.Ohms(3))
```

### Why this one?

`quant` is aimed at Go programs that want compile-time separation between
dimensions without a runtime registry or string parser in the core package.
Unlike APIs built from dedicated numeric types and one method per unit,
`quant` uses one generic `Quantity[D]` plus typed units. That makes custom
dimensions and generic helpers natural, while keeping conversions in small
value types. The trade-off is that the generic API cannot express every
physically valid product as an operator overload; named methods such as
`MulLength` make the supported derived relationships explicit.

`time.Duration` is an excellent standard-library example of a scalar with a
well-defined unit, but it is intentionally limited to time and does not carry
dimension information for arithmetic with length, mass, or electrical values.
`quant` follows that same value-oriented style while extending the type safety
to physical dimensions.

### Semantics and custom dimensions

`pH` is logarithmic, so `Add`, `Sub`, and scalar multiplication on
`Quantity[Acidity]` are only numeric operations, not general physical laws.
`PartsPer` is a ratio and `Pieces` is a count; treat arithmetic on them as
domain-specific. Absolute temperatures should be changed with `AddDelta` and
`SubDelta`; `Add` on `Quantity[Temperature]` remains available for generic
code but is not a physically meaningful temperature operation.

Frequency describes cycles per time, while `AngularVelocity` describes angle
per time. They have overlapping units (RPM and degrees/second) but represent
different physical quantities and therefore remain separate dimensions.
Torque and energy can both have N·m dimensions, but torque is kept separate
to prevent silently treating a rotational moment as transferred energy.

Custom dimensions are struct markers, and custom units implement `ToBase` and
`FromBase`:

```go
type Widget struct{}
type DozenWidgets struct{}

func (DozenWidgets) ToBase(v float64) float64   { return v * 12 }
func (DozenWidgets) FromBase(v float64) float64 { return v / 12 }

widgets := quant.New[Widget](2, DozenWidgets{}).To(DozenWidgets{}) // 2
```

Custom dimensions do not automatically acquire a base-unit symbol for
`String`; use `To` or `Format` with an application-owned formatter.

## Supported Dimensions

- Acceleration: `MeterPerSecondSquared`, `GForce`, `StandardGravity`
- Acidity: `PH`, `POH`
- Absorbed dose: `Gray`, `Rad`
- Angle: `Degree`, `Radian`, `Gradian`, `ArcMinute`, `ArcSecond`
- Angular acceleration: `AngularRadianPerSecondSquared`, `AngularDegreePerSecondSquared`
- Angular velocity: `AngularRadianPerSecond`, `AngularDegreePerSecond`, `AngularRevolutionPerMinute`
- Area: `SquareMillimeter` (`mm2`), `SquareCentimeter` (`cm2`), `SquareMeter` (`m2`), `Hectare` (`ha`), `SquareKilometer` (`km2`), `SquareInch` (`in2`), `SquareFoot` (`ft2`), `Acre` (`ac`), `SquareMile` (`mi2`)
- Apparent power: `VoltAmpere`, `MilliVoltAmpere`, `KiloVoltAmpere`, `MegaVoltAmpere`, `GigaVoltAmpere`
- Capacitance: `Farad`, `Microfarad`, `Nanofarad`, `Picofarad`
- Catalytic activity: `Katal`
- Charge: `Coulomb`, `Millicoulomb`, `Microcoulomb`, `Nanocoulomb`, `Picocoulomb`
- Concentration: `MolePerCubicMeter`, `Molar`, `Millimolar`
- Current: `Ampere`, `Milliampere`, `Kiloampere`
- Data rate: `BitPerSecond`, `KilobitPerSecond`, `MegabitPerSecond`, `GigabitPerSecond`, `TerabitPerSecond`, `BytePerSecond`, `KilobytePerSecond`, `MegabytePerSecond`, `GigabytePerSecond`, `TerabytePerSecond`, `KibibytePerSecond`, `MebibytePerSecond`, `GibibytePerSecond`, `TebibytePerSecond`
- Density: `KilogramPerCubicMeter`, `GramPerCubicCentimeter`, `PoundPerCubicFoot`, `KilogramPerLiter`
- Digital: `Bit`, `Kilobit`, `Megabit`, `Gigabit`, `Terabit`, `Byte`, `Kilobyte`, `Megabyte`, `Gigabyte`, `Terabyte`, `Petabyte`, `Kibibyte`, `Mebibyte`, `Gibibyte`, `Tebibyte`
- Dynamic viscosity: `PascalSecond`, `Poise`, `Centipoise`
- Electrical conductivity: `SiemensPerMeter`, `MillisiemensPerMeter`, `MicrosiemensPerMeter`, `SiemensPerCentimeter`, `MillisiemensPerCentimeter`, `MicrosiemensPerCentimeter`
- Electric field: `VoltPerMeter`, `NewtonPerCoulomb`
- Energy: `WattSecond`, `WattMinute`, `MilliwattHour`, `WattHour`, `KilowattHour`, `MegawattHour`, `GigawattHour`, `Joule`, `Kilojoule`, `Megajoule`, `Gigajoule`
- Equivalent dose: `Sievert`, `Rem`
- Force: `Newton`, `Kilonewton`, `PoundForce`, `KilogramForce`
- Frequency: `Hertz`, `Millihertz`, `Kilohertz`, `Megahertz`, `Gigahertz`, `Terahertz`, `RevolutionPerMinute`, `DegreePerSecond`, `RadianPerSecond`
- Illuminance: `Lux`, `FootCandle`
- Inductance: `Henry`, `Millihenry`, `Microhenry`
- Kinematic viscosity: `SquareMeterPerSecond`, `Stokes`, `Centistokes`
- Length: `Nanometer`, `Micrometer`, `Millimeter`, `Centimeter`, `Meter`, `Inch`, `Yard`, `USFoot`, `Foot`, `Fathom`, `Kilometer`, `Mile`, `NauticalMile`
- Luminous flux: `Lumen`
- Luminous intensity: `Candela`
- Magnetic flux: `Weber`, `Maxwell`
- Magnetic flux density: `Tesla`, `Millitesla`, `Gauss`
- Mass: `Microgram` (`mcg`), `Milligram` (`mg`), `Gram` (`g`), `Kilogram` (`kg`), `Ounce` (`oz`), `Pound` (`lb`), `MetricTon` (`mt`), `Stone` (`st`), `Tonne` (`t`)
- Molality: `MolePerKilogram`, `Molal`
- Momentum: `KilogramMeterPerSecond`, `NewtonSecond`
- Pace: `SecondPerMeter`, `MinutePerKilometer`, `SecondPerFoot`, `MinutePerMile`
- Parts-per: `PPM`, `PPB`, `PPT`, `PPQ`
- Pieces: `Piece`, `BakersDozen`, `Couple`, `DozenDozen`, `Dozen`, `GreatGross`, `Gross`, `HalfDozen`, `LongHundred`, `Ream`, `Score`, `SmallGross`, `Trio`
- Pressure: `Pascal`, `Hectopascal`, `Kilopascal`, `Megapascal`, `Bar`, `Torr`, `MeterOfWater`, `MillimeterOfMercury`, `PSI`, `KSI`
- Power: `Watt`, `Milliwatt`, `Kilowatt`, `Megawatt`, `Gigawatt`, `MetricHorsepower`, `BTUPerSecond`, `FootPoundForcePerSecond`, `Horsepower`
- Radioactivity: `Becquerel`, `Curie`
- Reactive energy: `VoltAmpereReactiveHour`, `MilliVoltAmpereReactiveHour`, `KiloVoltAmpereReactiveHour`, `MegaVoltAmpereReactiveHour`, `GigaVoltAmpereReactiveHour`
- Reactive power: `VoltAmpereReactive`, `MilliVoltAmpereReactive`, `KiloVoltAmpereReactive`, `MegaVoltAmpereReactive`, `GigaVoltAmpereReactive`
- Resistance: `Ohm`, `Milliohm`, `Kiloohm`, `Megaohm`
- Speed: `MeterPerSecond`, `KilometerPerHour`, `MilePerHour`, `MeterPerHour`, `Knot`, `FootPerSecond`, `InchPerHour`, `MillimeterPerHour`
- Specific heat capacity: `JoulePerKilogramKelvin`
- Surface tension: `NewtonPerMeter`, `DynePerCentimeter`
- Temperature: `Celsius`, `Fahrenheit`, `Kelvin`, `Rankine`
- Temperature delta: `KelvinDelta`, `CelsiusDelta`, `FahrenheitDelta`, `RankineDelta`
- Thermal conductivity: `WattPerMeterKelvin`
- Time: `Nanosecond`, `Microsecond`, `Millisecond`, `Second`, `Minute`, `Hour`, `Day`, `Week`, `Month`, `Year`, `Decade`, `Century`
- Torque: `NewtonMeter`, `PoundForceFoot`
- Voltage: `Volt`, `Millivolt`, `Kilovolt`
- Volume: `CubicMillimeter`, `CubicCentimeter`, `Milliliter`, `Liter`, `Kiloliter`, `Megaliter`, `Gigaliter`, `CubicMeter`, `CubicKilometer`, `Teaspoon`, `Tablespoon`, `CubicInch`, `FluidOunce`, `Cup`, `Pint`, `Quart`, `Gallon`, `CubicFoot`, `CubicYard`
- Volume flow rate: `CubicMillimeterPerSecond`, `CubicCentimeterPerSecond`, `MilliliterPerSecond`, `CentiliterPerSecond`, `DeciliterPerSecond`, `LiterPerSecond`, `LiterPerMinute`, `LiterPerHour`, `KiloliterPerSecond`, `KiloliterPerMinute`, `KiloliterPerHour`, `CubicMeterPerSecond`, `CubicMeterPerMinute`, `CubicMeterPerHour`, `CubicKilometerPerSecond`, `TeaspoonPerSecond`, `TablespoonPerSecond`, `CubicInchPerSecond`, `CubicInchPerMinute`, `CubicInchPerHour`, `FluidOuncePerSecond`, `FluidOuncePerMinute`, `FluidOuncePerHour`, `CupPerSecond`, `PintPerSecond`, `PintPerMinute`, `PintPerHour`, `QuartPerSecond`, `GallonPerSecond`, `GallonPerMinute`, `GallonPerHour`, `CubicFootPerSecond`, `CubicFootPerMinute`, `CubicFootPerHour`, `CubicYardPerSecond`, `CubicYardPerMinute`, `CubicYardPerHour`

## Type Safety

Dimensions are enforced by the type system.

These compile:

```go
distance := quant.New[quant.Length](5, quant.Meter)
extra := quant.New[quant.Length](2, quant.Kilometer)
total := distance.Add(extra)
_ = total.To(quant.Mile)
```

These do not compile:

```go
distance := quant.New[quant.Length](5, quant.Meter)
mass := quant.New[quant.Mass](2, quant.Kilogram)

_ = distance.Add(mass)
_ = distance.To(quant.Pound)
```

## Temperature Conversions

Temperature uses affine conversion rather than a pure scale factor.

```go
k := quant.New[quant.Temperature](25, quant.Celsius).To(quant.Kelvin)
fmt.Printf("%.2f\n", k) // 298.15
```

## Acidity Conversions

Acidity supports `PH` and `POH`. `PH` is the base unit, and `POH` uses the standard 25 C relationship `pH + pOH = 14`.

```go
poh := quant.PHValue(3).To(quant.POH)
fmt.Println(poh) // 11
```

`Month`, `Year`, `Decade`, and `Century` use average Gregorian durations: `365.25 / 12` days, `365.25` days, `10 * 365.25` days, and `100 * 365.25` days respectively.

## Design Notes

- No reflection
- No string parsing
- No runtime registry
- No `interface{}`
- No external dependencies

The package favors correctness and simplicity first, while keeping the implementation small and allocation-free in normal use.

## Development

Run tests with:

```bash
go test ./...
```

Run benchmarks with:

```bash
go test -bench . ./...
```

Recent benchmark results on Apple M4:

```text
BenchmarkNew-10        1000000000   0.7779 ns/op
BenchmarkTo-10          175765758   6.900 ns/op
BenchmarkAdd-10        1000000000   0.2607 ns/op
BenchmarkDiv-10        1000000000   0.2594 ns/op
BenchmarkString-10       17407522  71.04 ns/op
BenchmarkFormat-10       15354643  77.52 ns/op
```
