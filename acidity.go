package quant

// Acidity is the dimension marker for acidity quantities.
//
// PH is the base unit. POH conversions use the standard 25 C relationship:
// pH + pOH = 14.
type Acidity struct{}

// Acidity units.
var (
	PH  = scaleUnit[Acidity]{factor: 1}
	POH = affineUnit[Acidity]{factor: -1, offset: -14}
)
