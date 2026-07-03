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
- Text marshaling and `database/sql` interop helpers
- Support for affine temperature conversions

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

func (q Quantity[D]) To(u Unit[D]) float64
func (q Quantity[D]) Value(u Unit[D]) float64
func (q Quantity[D]) Format(u Unit[D]) string
func (q Quantity[D]) MarshalJSON() ([]byte, error)
func (q *Quantity[D]) UnmarshalJSON(data []byte) error
func (q Quantity[D]) MarshalText() ([]byte, error)
func (q *Quantity[D]) UnmarshalText(text []byte) error
func (q Quantity[D]) Add(other Quantity[D]) Quantity[D]
func (q Quantity[D]) Sub(other Quantity[D]) Quantity[D]
```

`Amount[D](v).From(unit)` is also available as a fluent constructor.

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

JSON marshaling stores the scalar value in the base unit for the dimension:

```go
data, _ := json.Marshal(quant.Pounds(10))
fmt.Println(string(data)) // 4.535923700000001
```

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
- Angle: radian
- Area: square meter
- Apparent power: volt-ampere
- Charge: coulomb
- Current: ampere
- Digital: bit
- Energy: joule
- Force: newton
- Frequency: hertz
- Illuminance: lux
- Length: meter
- Mass: kilogram
- Pace: second per meter
- Parts-per: ratio
- Pieces: piece
- Power: watt
- Pressure: pascal
- Reactive energy: volt-ampere reactive hour
- Reactive power: volt-ampere reactive
- Speed: meter per second
- Temperature: kelvin
- Time: second
- Torque: newton-meter
- Voltage: volt
- Volume: cubic meter
- Volume flow rate: cubic meter per second

## Examples By Dimension

- Acceleration: `ms2 := quant.GForces(1).To(quant.MeterPerSecondSquared)`
- Angle: `rad := quant.Degrees(180).To(quant.Radian)`
- Area: `m2 := quant.Acres(1).To(quant.SquareMeter)`
- Apparent power: `va := quant.MegaVoltAmperes(1).To(quant.VoltAmpere)`
- Charge: `mc := quant.Coulombs(1).To(quant.Millicoulomb)`
- Current: `a := quant.Kiloamperes(1).To(quant.Ampere)`
- Digital: `bytes := quant.Kibibytes(1).To(quant.Byte)`
- Energy: `j := quant.KilowattHours(1).To(quant.Joule)`
- Force: `n := quant.KilogramsForce(1).To(quant.Newton)`
- Frequency: `hz := quant.RevolutionsPerMinute(60).To(quant.Hertz)`
- Illuminance: `lx := quant.FootCandles(1).To(quant.Lux)`
- Length: `km := quant.Miles(3).To(quant.Kilometer)`
- Mass: `kg := quant.Pounds(10).To(quant.Kilogram)`
- Pace: `spm := quant.MinutesPerKilometer(5).To(quant.SecondPerMeter)`
- Parts-per: `ppb := quant.PartsPerMillion(1).To(quant.PPB)`
- Pieces: `pcs := quant.Dozens(1).To(quant.Piece)`
- Power: `w := quant.HorsepowerValues(1).To(quant.Watt)`
- Pressure: `pa := quant.Bars(1).To(quant.Pascal)`
- Reactive energy: `varh := quant.MegaVoltAmpereReactiveHours(1).To(quant.KiloVoltAmpereReactiveHour)`
- Reactive power: `vars := quant.MegaVoltAmpereReactives(1).To(quant.KiloVoltAmpereReactive)`
- Speed: `kmh := quant.MetersPerSecond(10).To(quant.KilometerPerHour)`
- Temperature: `k := quant.DegreesCelsius(25).To(quant.Kelvin)`
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

### Derived units

```go
distance := quant.Kilometers(5)
duration := quant.Minutes(30)
speed := distance.Div(duration)

fmt.Println(speed.To(quant.KilometerPerHour)) // 10
```

## Supported Dimensions

- Acceleration: `MeterPerSecondSquared`, `GForce`, `StandardGravity`
- Angle: `Degree`, `Radian`, `Gradian`, `ArcMinute`, `ArcSecond`
- Area: `SquareMillimeter` (`mm2`), `SquareCentimeter` (`cm2`), `SquareMeter` (`m2`), `Hectare` (`ha`), `SquareKilometer` (`km2`), `SquareInch` (`in2`), `SquareFoot` (`ft2`), `Acre` (`ac`), `SquareMile` (`mi2`)
- Apparent power: `VoltAmpere`, `MilliVoltAmpere`, `KiloVoltAmpere`, `MegaVoltAmpere`, `GigaVoltAmpere`
- Charge: `Coulomb`, `Millicoulomb`, `Microcoulomb`, `Nanocoulomb`, `Picocoulomb`
- Current: `Ampere`, `Milliampere`, `Kiloampere`
- Digital: `Bit`, `Kilobit`, `Megabit`, `Gigabit`, `Terabit`, `Byte`, `Kilobyte`, `Megabyte`, `Gigabyte`, `Terabyte`, `Petabyte`, `Kibibyte`, `Mebibyte`, `Gibibyte`, `Tebibyte`
- Energy: `WattSecond`, `WattMinute`, `MilliwattHour`, `WattHour`, `KilowattHour`, `MegawattHour`, `GigawattHour`, `Joule`, `Kilojoule`, `Megajoule`, `Gigajoule`
- Force: `Newton`, `Kilonewton`, `PoundForce`, `KilogramForce`
- Frequency: `Hertz`, `Millihertz`, `Kilohertz`, `Megahertz`, `Gigahertz`, `Terahertz`, `RevolutionPerMinute`, `DegreePerSecond`, `RadianPerSecond`
- Illuminance: `Lux`, `FootCandle`
- Length: `Nanometer`, `Micrometer`, `Millimeter`, `Centimeter`, `Meter`, `Inch`, `Yard`, `USFoot`, `Foot`, `Fathom`, `Kilometer`, `Mile`, `NauticalMile`
- Mass: `Microgram` (`mcg`), `Milligram` (`mg`), `Gram` (`g`), `Kilogram` (`kg`), `Ounce` (`oz`), `Pound` (`lb`), `MetricTon` (`mt`), `Stone` (`st`), `Tonne` (`t`)
- Pace: `SecondPerMeter`, `MinutePerKilometer`, `SecondPerFoot`, `MinutePerMile`
- Parts-per: `PPM`, `PPB`, `PPT`, `PPQ`
- Pieces: `Piece`, `BakersDozen`, `Couple`, `DozenDozen`, `Dozen`, `GreatGross`, `Gross`, `HalfDozen`, `LongHundred`, `Ream`, `Score`, `SmallGross`, `Trio`
- Pressure: `Pascal`, `Hectopascal`, `Kilopascal`, `Megapascal`, `Bar`, `Torr`, `MeterOfWater`, `MillimeterOfMercury`, `PSI`, `KSI`
- Power: `Watt`, `Milliwatt`, `Kilowatt`, `Megawatt`, `Gigawatt`, `MetricHorsepower`, `BTUPerSecond`, `FootPoundForcePerSecond`, `Horsepower`
- Reactive power: `VoltAmpereReactive`, `MilliVoltAmpereReactive`, `KiloVoltAmpereReactive`, `MegaVoltAmpereReactive`, `GigaVoltAmpereReactive`
- Reactive energy: `VoltAmpereReactiveHour`, `MilliVoltAmpereReactiveHour`, `KiloVoltAmpereReactiveHour`, `MegaVoltAmpereReactiveHour`, `GigaVoltAmpereReactiveHour`
- Speed: `MeterPerSecond`, `KilometerPerHour`, `MilePerHour`, `MeterPerHour`, `Knot`, `FootPerSecond`, `InchPerHour`, `MillimeterPerHour`
- Temperature: `Celsius`, `Fahrenheit`, `Kelvin`, `Rankine`
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
