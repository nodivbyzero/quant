package quant

// Pieces is the dimension marker for count-based quantities.
type Pieces struct{}

// Pieces units.
var (
	Piece       = scaleUnit[Pieces]{factor: 1}
	BakersDozen = scaleUnit[Pieces]{factor: 13}
	Couple      = scaleUnit[Pieces]{factor: 2}
	DozenDozen  = scaleUnit[Pieces]{factor: 144}
	Dozen       = scaleUnit[Pieces]{factor: 12}
	GreatGross  = scaleUnit[Pieces]{factor: 1728}
	Gross       = scaleUnit[Pieces]{factor: 144}
	HalfDozen   = scaleUnit[Pieces]{factor: 6}
	LongHundred = scaleUnit[Pieces]{factor: 120}
	Ream        = scaleUnit[Pieces]{factor: 500}
	Score       = scaleUnit[Pieces]{factor: 20}
	SmallGross  = scaleUnit[Pieces]{factor: 120}
	Trio        = scaleUnit[Pieces]{factor: 3}
)
