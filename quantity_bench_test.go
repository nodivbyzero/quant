package quant_test

import (
	"testing"

	"github.com/nodivbyzero/quant"
)

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = quant.New[quant.Length](5, quant.Kilometer)
	}
}

func BenchmarkTo(b *testing.B) {
	q := quant.Kilometers(5)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = q.To(quant.Mile)
	}
}

func BenchmarkAdd(b *testing.B) {
	a := quant.Kilometers(5)
	c := quant.Meters(250)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = a.Add(c)
	}
}

func BenchmarkDiv(b *testing.B) {
	distance := quant.Kilometers(5)
	duration := quant.Minutes(30)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = distance.Div(duration)
	}
}

func BenchmarkString(b *testing.B) {
	q := quant.Pounds(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = q.String()
	}
}

func BenchmarkFormat(b *testing.B) {
	q := quant.Pounds(10)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = q.Format(quant.Pound)
	}
}
