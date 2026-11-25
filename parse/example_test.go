package parse_test

import (
	"fmt"

	"github.com/nodivbyzero/quant"
	"github.com/nodivbyzero/quant/parse"
)

func ExampleRegistry() {
	lengths := parse.Lengths()
	q, err := lengths.Parse("5 kilometers")
	if err != nil {
		return
	}
	fmt.Println(q.To(quant.Meter))
	// Output: 5000
}

func ExampleRegistry_WithDefault() {
	seconds := parse.Times().WithDefault(quant.Second)
	q, err := seconds.Parse("30")
	if err != nil {
		return
	}
	fmt.Println(q.To(quant.Second))
	// Output: 30
}

func ExampleValue() {
	timeout := parse.NewValue(parse.Times())
	if err := timeout.Set("30 s"); err != nil {
		return
	}
	q, _ := timeout.Quantity()
	fmt.Println(q.To(quant.Second))
	// Output: 30
}
