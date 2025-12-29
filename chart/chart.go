package chart

import "github.com/u00io/nuiforms/ui"

type Chart struct {
	ui.Widget

	s *Settings
}

func NewChart() *Chart {
	var c Chart
	c.InitWidget()
	c.SetOnPaint(c.draw)

	c.s = NewSettings()

	return &c
}

func (c *Chart) draw(cnv *ui.Canvas) {
	cnv.SetColor(ui.ColorFromHex("#777777"))
	cnv.SetFontFamily(c.FontFamily())
	cnv.SetFontSize(c.FontSize())
	c.s.draw(cnv, c.Width(), c.Height())
}
