// Package quant provides type-safe physical quantities and unit conversions
// using generics.
//
// Quantities are stored in dimension-specific base units, and conversions are
// checked at compile time through phantom type parameters. The package includes
// built-in dimensions for common physical quantities, convenience constructors,
// same-dimension arithmetic, derived units such as speed from length divided by
// time, formatting helpers, and JSON, text, and SQL interop based on base-unit
// scalar values.
package quant
