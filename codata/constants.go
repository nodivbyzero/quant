// Package codata contains selected 2022 CODATA recommended constants.
//
// Values are exposed as float64 constants in the SI units documented by NIST.
// Uncertainties are not represented; consult the source when uncertainty or
// correlations matter. This package intentionally does not invent Quantity
// dimensions for compound units that are not modeled by quant.
package codata

const (
	// SpeedOfLight is the speed of light in vacuum, in m/s (exact).
	SpeedOfLight = 299792458.0
	// PlanckConstant is Planck's constant, in J s (exact).
	PlanckConstant = 6.62607015e-34
	// ReducedPlanckConstant is the reduced Planck constant, in J s.
	ReducedPlanckConstant = 1.054571817e-34
	// ElementaryCharge is the elementary charge, in C (exact).
	ElementaryCharge = 1.602176634e-19
	// BoltzmannConstant is the Boltzmann constant, in J/K (exact).
	BoltzmannConstant = 1.380649e-23
	// AvogadroConstant is the Avogadro constant, in 1/mol (exact).
	AvogadroConstant = 6.02214076e23
	// GravitationalConstant is Newton's constant of gravitation, in m^3 kg^-1 s^-2.
	GravitationalConstant = 6.67430e-11
	// FineStructureConstant is the fine-structure constant, dimensionless.
	FineStructureConstant = 7.2973525643e-3
	// ElectronMass is the electron mass, in kg.
	ElectronMass = 9.1093837139e-31
	// ProtonMass is the proton mass, in kg.
	ProtonMass = 1.67262192595e-27
	// MolarGasConstant is the molar gas constant, in J/(mol K).
	MolarGasConstant = 8.314462618
	// VacuumPermittivity is the vacuum electric permittivity, in F/m.
	VacuumPermittivity = 8.8541878188e-12
)
