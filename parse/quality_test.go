package parse_test

import (
	"errors"
	"math"
	"sync"
	"testing"

	"github.com/nodivbyzero/quant"
	"github.com/nodivbyzero/quant/parse"
)

func FuzzRegistryParse(f *testing.F) {
	for _, seed := range []string{"5 m", "2 kilometers", "1e3 mm", "-4.5 ft", "NaN m", "5 furlongs", ""} {
		f.Add(seed)
	}
	registry := parse.Lengths()
	f.Fuzz(func(t *testing.T, input string) {
		q, err := registry.Parse(input)
		if err == nil && (math.IsNaN(q.To(quant.Meter)) || math.IsInf(q.To(quant.Meter), 0)) {
			t.Fatalf("successful parse produced non-finite quantity for %q", input)
		}
	})
}

func TestParserErrorsAcrossDimensions(t *testing.T) {
	tests := []struct {
		name  string
		parse func(string) error
	}{
		{"length", func(s string) error { _, err := parse.Lengths().Parse(s); return err }},
		{"mass", func(s string) error { _, err := parse.Masses().Parse(s); return err }},
		{"time", func(s string) error { _, err := parse.Times().Parse(s); return err }},
		{"temperature", func(s string) error { _, err := parse.Temperatures().Parse(s); return err }},
		{"speed", func(s string) error { _, err := parse.Speeds().Parse(s); return err }},
		{"energy", func(s string) error { _, err := parse.Energies().Parse(s); return err }},
		{"voltage", func(s string) error { _, err := parse.Voltages().Parse(s); return err }},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.parse("not-a-number m"); !errors.Is(err, parse.ErrSyntax) {
				t.Fatalf("syntax error = %v", err)
			}
			if err := test.parse("1 unknown-unit"); !errors.Is(err, parse.ErrUnknownUnit) {
				t.Fatalf("unknown-unit error = %v", err)
			}
		})
	}
}

func TestRegistryParseAllocations(t *testing.T) {
	registry := parse.Lengths()
	allocs := testing.AllocsPerRun(1000, func() {
		if _, err := registry.Parse("12.5 km"); err != nil {
			t.Fatal(err)
		}
	})
	if allocs != 0 {
		t.Fatalf("successful parse allocations = %v, want 0", allocs)
	}
}

func TestRegistryConcurrentParse(t *testing.T) {
	registry := parse.Lengths()
	const workers = 16
	const iterations = 1000
	var wg sync.WaitGroup
	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				q, err := registry.Parse("2 kilometers")
				if err != nil || q.To(quant.Meter) != 2000 {
					t.Errorf("parse: q=%v err=%v", q.To(quant.Meter), err)
					return
				}
			}
		}()
	}
	wg.Wait()
}

// Compile-time boundary checks: a registry is parameterized by its dimension,
// so a mass unit cannot be passed to a length registry. The negative cases are
// kept as compile-fail examples in quantity_test.go.
var (
	_ parse.Registry[quant.Length] = parse.Lengths()
	_ parse.Registry[quant.Mass]   = parse.Masses()
)
