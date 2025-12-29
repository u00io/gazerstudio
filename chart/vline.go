package chart

type VLine struct {
	X         int
	HasValues bool
	MinYValue float64
	MaxYValue float64
	MinY      float64
	MaxY      float64
	FirstY    float64
	LastY     float64

	MinYp int
	MaxYp int

	HasY          bool
	HasBegin      bool
	HasEnd        bool
	HasBadQuality bool
}

func NewVLine() *VLine {
	var c VLine
	return &c
}
