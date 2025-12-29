package chart

type DataItemValue struct {
	DatetimeFirst int64
	DatetimeLast  int64
	FirstValue    float64
	LastValue     float64
	MinValue      float64
	MaxValue      float64
	AvgValue      float64
	SumValue      float64
	CountOfValues int64
	Qualities     []int
	HasGood       bool
	HasBad        bool
	Uom           string
}
