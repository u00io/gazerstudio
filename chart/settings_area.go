package chart

import (
	"image/color"

	"github.com/u00io/nuiforms/ui"
)

type SettingsArea struct {
	verticalScalesWidth float64
	xOffset             int
	yOffset             int
	width               int
	height              int
	selected            bool
	series              []*SettingsSeries

	unitedVScale *VerticalScale

	// Props
	showVerticalScale   bool
	unitedVerticalScale bool
	showZero            bool
}

func NewSettingsArea() *SettingsArea {
	var c SettingsArea
	c.showVerticalScale = true
	return &c
}

func (c *SettingsArea) AddSeries(s *SettingsSeries) {
	c.series = append(c.series, s)
}

func (c *SettingsArea) Calc(settings *Settings, x, y, w, h, vsWidth int) {
	c.xOffset = x
	c.yOffset = y
	c.width = w
	c.height = h

	if c.UnitedVerticalScale() {
		vScale := NewVerticalScale()
		for seriesIndex, s := range c.series {
			vScale.UpdateVerticalScaleValues(s.itemHistory, true)
			s.Calc(x, y, w, h, vsWidth, vScale, float64(seriesIndex)*settings.legendItemHeight+settings.legendItemHeight/2+settings.legendItemYOffset)
			s.vScale.Calc(0, y, DefaultVerticalScaleWidthInline, h)
		}
		if c.ShowZero() {
			vScale.ExpandToZero()
		}
	} else {
		for seriesIndex, s := range c.series {
			vScale := NewVerticalScale()
			vScale.UpdateVerticalScaleValues(s.itemHistory, false)
			s.Calc(x, y, w, h, vsWidth, vScale, float64(seriesIndex)*settings.legendItemHeight+settings.legendItemHeight/2+settings.legendItemYOffset)
			//s.vScale.Calc(seriesIndex*DefaultVerticalScaleWidthInline, y, 50, h)
			s.vScale.Calc(seriesIndex*DefaultVerticalScaleWidthInline, 0, 50, h)
			if c.ShowZero() || s.ShowZero() {
				s.vScale.ExpandToZero()
			}
		}
	}
}

func (c *SettingsArea) UnitedVerticalScale() bool {
	return c.unitedVerticalScale
}

func (c *SettingsArea) ShowZero() bool {
	return c.showZero
}

func (c *SettingsArea) Draw(cnv *ui.Canvas, width int, height int, hScale *HorizontalScale, settings *Settings, last bool) {
	cnv.Save()
	cnv.TranslateAndClip(0, c.yOffset, width, c.height)
	someSeriesSelected := false
	for _, s := range c.series {
		if s.selected {
			someSeriesSelected = true
			break
		}
	}

	for seriesIndex, s := range c.series {
		smooth := false
		if settings.editing && someSeriesSelected {
			if !s.selected {
				smooth = true
			}
		}
		s.draw(cnv, width, height, hScale, settings, smooth, seriesIndex, len(c.series))
	}

	for seriesIndex, s := range c.series {
		s.drawDetails(cnv, width, height, hScale, settings, false, seriesIndex, len(c.series))
	}

	if c.showVerticalScale && len(c.series) > 0 {
		if c.UnitedVerticalScale() {
			s := c.series[0]
			vScaleColor := ui.ColorFromHex("#88339933")
			if len(c.series) == 1 {
				// vScaleColor = s.getColor("stroke_color")
			}
			s.vScale.draw(cnv, vScaleColor, 0, settings.showLegend, 1)
		} else {
			for seriesIndex, s := range c.series {
				vScaleColor := ui.ColorFromHex("#88339933")
				s.vScale.draw(cnv, vScaleColor, seriesIndex, settings.showLegend, len(c.series))
			}
		}
	}

	if !last {
		cnv.DrawLine(c.xOffset, c.yOffset+c.height-1, c.xOffset+c.width, c.yOffset+c.height-1, 1, color.RGBA{100, 100, 100, 255})
	}

	seriesInAreaSelected := false

	for _, s := range c.series {
		if s.selected {
			seriesInAreaSelected = true
		}
	}

	if c.selected || seriesInAreaSelected {
		cnv.SetColor(ui.ColorFromHex("#561726"))
		cnv.DrawRect(10, 10, c.width-20, c.height-20)
	}

	cnv.Restore()
}

/*
  void draw(Canvas canvas, Size size, TimeChartHorizontalScale hScale,
      TimeChartSettings settings, bool last) {
    canvas.save();
    canvas.clipRect(Rect.fromLTWH(0, yOffset, size.width, height));
    canvas.translate(0, yOffset);
    bool someSeriesSelected = false;
    for (int seriesIndex = 0; seriesIndex < series.length; seriesIndex++) {
      var s = series[seriesIndex];
      if (s.selected) {
        someSeriesSelected = true;
        break;
      }
    }

    for (int seriesIndex = 0; seriesIndex < series.length; seriesIndex++) {
      var s = series[seriesIndex];
      bool smooth = false;
      if (settings.editing() && someSeriesSelected) {
        if (!s.selected) {
          smooth = true;
        }
      }
      s.draw(
          canvas, size, hScale, settings, smooth, seriesIndex, series.length);
    }

    for (int seriesIndex = 0; seriesIndex < series.length; seriesIndex++) {
      var s = series[seriesIndex];
      s.drawDetails(
          canvas, size, hScale, settings, false, seriesIndex, series.length);
    }

    if (settings.showVerticalScale && series.isNotEmpty) {
      if (unitedVerticalScale()) {
        var s = series[0];
        var vScaleColor = DesignColors.fore();
        if (series.length == 1) {
          vScaleColor = s.getColor("stroke_color");
        }

        s.vScale.draw(canvas, size, vScaleColor, 0, getBool("show_legend"), 1);
      } else {
        for (int seriesIndex = 0; seriesIndex < series.length; seriesIndex++) {
          var s = series[seriesIndex];
          s.vScale.draw(canvas, size, s.getColor("stroke_color"), seriesIndex,
              getBool("show_legend"), series.length);
        }
      }
    }

    if (!last) {
      canvas.drawLine(
          Offset(xOffset, height - 1),
          Offset(xOffset + width, height - 1),
          Paint()
            ..color = Colors.blueGrey
            ..strokeWidth = 2);
    }

    bool seriesInAreaSelected = false;

    for (var s in series) {
      if (s.selected) {
        seriesInAreaSelected = true;
      }
    }

    if (selected || seriesInAreaSelected) {
      canvas.drawRect(
          Rect.fromLTWH(10, 10, width - 20, height - 20),
          Paint()
            ..color = Colors.yellowAccent.withOpacity(0.2)
            ..strokeWidth = 5
            ..style = PaintingStyle.stroke);
    }

    canvas.restore();
  }

*/
