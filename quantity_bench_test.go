package quant_test

import (
	"testing"

	"github.com/nodivbyzero/quant"
)

var (
	benchmarkQuantity quant.Quantity[quant.Length]
	benchmarkSpeed    quant.Quantity[quant.Speed]
	benchmarkFloat    float64
	benchmarkString   string
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkQuantity = quant.New[quant.Length](5, quant.Kilometer)
	}
}

func BenchmarkTo(b *testing.B) {
	q := quant.Kilometers(5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkFloat = q.To(quant.Mile)
	}
}

func BenchmarkAdd(b *testing.B) {
	a := quant.Kilometers(5)
	c := quant.Meters(250)
	result := a
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		result = result.Add(c)
	}
	benchmarkQuantity = result
}

func BenchmarkDiv(b *testing.B) {
	distance := quant.Kilometers(5)
	duration := quant.Minutes(30)
	result := distance
	step := quant.Meters(250)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkSpeed = result.Div(duration)
		result = result.Add(step)
	}
}

func BenchmarkString(b *testing.B) {
	q := quant.Pounds(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkString = q.String()
	}
}

func BenchmarkFormat(b *testing.B) {
	q := quant.Pounds(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchmarkString = q.Format(quant.Pound)
	}
}
