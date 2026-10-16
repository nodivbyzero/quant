package codata_test

import (
	"testing"

	"github.com/nodivbyzero/quant/codata"
)

func TestDefiningConstants(t *testing.T) {
	if codata.SpeedOfLight != 299792458 {
		t.Fatalf("speed of light changed: got %g", codata.SpeedOfLight)
	}
	if codata.ElementaryCharge != 1.602176634e-19 {
		t.Fatalf("elementary charge changed: got %g", codata.ElementaryCharge)
	}
	if codata.AvogadroConstant != 6.02214076e23 {
		t.Fatalf("Avogadro constant changed: got %g", codata.AvogadroConstant)
	}
}
