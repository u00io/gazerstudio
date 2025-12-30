package chart

import (
	"image/color"
	"math"

	"github.com/u00io/nuiforms/ui"
)

type Settings struct {
	areas    []*SettingsArea
	horScale *HorizontalScale

	selectionIsStarted  bool
	selectionMin        float64
	selectionMax        float64
	selectionIsFinished bool

	selectionMovingIsStarted           bool
	selectionMovingStartPositionPixels float64
	selectionMovingOriginalMin         float64
	selectionMovingOriginalMax         float64
	selectionMovingIsFinished          bool

	selectionResizingLeftIsStarted            bool
	selectionResizingLeftStartPositionPixels  float64
	selectionResizingLeftIsFinished           bool
	selectionResizingRightIsStarted           bool
	selectionResizingRightStartPositionPixels float64
	selectionResizingRightIsFinished          bool

	selectionForZoomIsStarted  bool
	selectionForZoomMin        float64
	selectionForZoomMax        float64
	selectionIsForZoomFinished bool

	movingIsStarted           bool
	movingStartPositionPixels float64
	movingOriginalDisplayMin  float64
	movingOriginalDisplayMax  float64

	scalingIsStarted           bool
	scalingStartPositionPixels float64
	scalingOriginalDisplayMin  float64
	scalingOriginalDisplayMax  float64

	hoverPosX float64
	hoverPosY float64

	editing bool

	showTimeScale     bool
	showVerticalScale bool
	backColor         color.Color

	keyControl bool
	keyAlt     bool
	keyShift   bool

	// Props
	legendItemWidth   float64
	legendItemHeight  float64
	legendItemXOffset float64
	legendItemYOffset float64

	showLegend bool
}

func NewSettings() *Settings {
	var c Settings

	c.showLegend = true

	c.legendItemWidth = 250
	c.legendItemHeight = 22
	c.legendItemXOffset = 0
	c.legendItemYOffset = 0
	c.horScale = NewHorizontalScale()

	// Demo data 1
	{
		area := NewSettingsArea()
		series := NewSettingsSeries()
		series.displayName = "Demo Series"
		area.series = append(area.series, series)
		c.areas = append(c.areas, area)

		itemHistory := make([]*DataItemValue, 0)
		for i := int64(100); i < 3500; i++ {
			t := i * 1000000
			v := math.Sin(float64(i)/180*3.14159) * 50.0
			item := &DataItemValue{
				DatetimeFirst: t,
				DatetimeLast:  t + 999999,
				FirstValue:    v,
				LastValue:     v,
				MinValue:      v,
				MaxValue:      v,
				AvgValue:      v,
				SumValue:      v,
				CountOfValues: 1,
				Qualities:     []int{1},
				HasGood:       true,
				HasBad:        false,
				Uom:           "units",
			}
			itemHistory = append(itemHistory, item)
		}
		series.itemHistory = itemHistory
	}

	// Demo data 2
	{
		area := NewSettingsArea()
		series := NewSettingsSeries()
		series.displayName = "Demo Series"
		area.series = append(area.series, series)
		c.areas = append(c.areas, area)

		itemHistory := make([]*DataItemValue, 0)
		for i := int64(100); i < 3500; i++ {
			t := i * 1000000
			v := math.Sin(float64(i)/100*3.14159)*50.0 + 50.0
			item := &DataItemValue{
				DatetimeFirst: t,
				DatetimeLast:  t + 999999,
				FirstValue:    v,
				LastValue:     v,
				MinValue:      v,
				MaxValue:      v,
				AvgValue:      v,
				SumValue:      v,
				CountOfValues: 1,
				Qualities:     []int{1},
				HasGood:       true,
				HasBad:        false,
				Uom:           "units",
			}
			itemHistory = append(itemHistory, item)
		}
		series.itemHistory = itemHistory
	}

	c.backColor = color.RGBA{15, 15, 15, 255}

	c.horScale.SetDefaultDisplayRange(0, 3600*1000000)
	c.horScale.ResetToDefaultDisplayRange()

	return &c
}

func (c *Settings) RemoveAllAreas() {
	c.areas = make([]*SettingsArea, 0)
}

func (c *Settings) AddArea(area *SettingsArea) {
	c.areas = append(c.areas, area)
}

func (c *Settings) Areas() []*SettingsArea {
	return c.areas
}

func (c *Settings) HorizontalScale() *HorizontalScale {
	return c.horScale
}

func (c *Settings) draw(cnv *ui.Canvas, width, height int) {
	verticalScalesWidth := 0

	if len(c.areas) < 1 {
		return
	}

	cnvHeight := height
	cnvHeight -= c.horScale.height
	areaHeight := cnvHeight / len(c.areas)

	cnv.FillRect(0, 0, width, height, c.backColor)

	for areaIndex := 0; areaIndex < len(c.areas); areaIndex++ {
		c.areas[areaIndex].Calc(c, 0, areaHeight*areaIndex, width, areaHeight, verticalScalesWidth)
	}

	c.horScale.Calc(verticalScalesWidth, areaHeight*len(c.areas), width-verticalScalesWidth, 30)

	for areaIndex := 0; areaIndex < len(c.areas); areaIndex++ {
		area := c.areas[areaIndex]
		area.Draw(cnv, width, areaHeight, c.horScale, c, areaIndex == len(c.areas)-1)
	}

	c.horScale.draw(cnv, width, height)
}

/*

  void draw(Canvas canvas, Size size) {
    double verticalScalesWidth = 0;
    double areaHeight = (size.height - horScale.height) / areas.length;

    canvas.drawRect(
        Offset.zero & size,
        Paint()
          ..style = PaintingStyle.fill
          ..color = backColor);

    for (int areaIndex = 0; areaIndex < areas.length; areaIndex++) {
      areas[areaIndex].calc(this, 0, areaHeight * areaIndex.toDouble(),
          size.width, areaHeight, verticalScalesWidth);
    }

    horScale.calc(verticalScalesWidth, areaHeight * areas.length,
        size.width - verticalScalesWidth, showTimeScale ? 30 : 0);

    for (int areaIndex = 0; areaIndex < areas.length; areaIndex++) {
      var area = areas[areaIndex];
      area.draw(canvas, size, horScale, this, areaIndex == areas.length - 1);

      if (_editing) {
        canvas.save();
        canvas
            .clipRect(Rect.fromLTWH(0, area.yOffset, size.width, area.height));
        canvas.translate(0, area.yOffset);
        if (area.selected) {
          drawEditButton(
              canvas, 0, "Area #${areaIndex + 1}", Colors.white, true);
        } else {
          drawEditButton(
              canvas, 0, "Area #${areaIndex + 1}", Colors.white, false);
        }
        int index = 1;
        for (var s in area.series) {
          if (s.selected) {
            drawEditButton(canvas, index, s.getDisplayName(),
                s.getColor("stroke_color"), true);
          } else {
            drawEditButton(canvas, index, s.getDisplayName(),
                s.getColor("stroke_color"), false);
          }
          index++;
        }

        canvas.restore();
      } else {
        if (area.getBool("show_legend")) {
          int index = 0;
          canvas.save();
          canvas.clipRect(
              Rect.fromLTWH(0, area.yOffset, size.width, area.height));
          canvas.translate(0, area.yOffset);
          for (var s in area.series) {
            if (s.selected) {
              drawLegendItem(
                  canvas,
                  index,
                  s.getDisplayName(),
                  s.getColor("stroke_color"),
                  area.series.length,
                  area.unitedVerticalScale() || area.series.length == 1);
            } else {
              drawLegendItem(
                  canvas,
                  index,
                  s.getDisplayName(),
                  s.getColor("stroke_color"),
                  area.series.length,
                  area.unitedVerticalScale() || area.series.length == 1);
            }
            index++;
          }
          canvas.restore();
        }
      }
    }

    horScale.draw(canvas, size);

    if (selectionForZoomIsStarted) {
      canvas.drawRect(
          Rect.fromLTRB(horScale.horValueToPixel(selectionForZoomMin), 0,
              horScale.horValueToPixel(selectionForZoomMax), size.height),
          Paint()
            ..style = PaintingStyle.fill
            ..color = Colors.white30
            ..strokeWidth = 1);
    }

    if (selectionMin != selectionMax) {
      canvas.save();
      canvas.clipRect(Rect.fromLTWH(0, 0, size.width, size.height));
      canvas.drawRect(
          Rect.fromLTRB(horScale.horValueToPixel(selectionMin), 0,
              horScale.horValueToPixel(selectionMax), size.height),
          Paint()
            ..style = PaintingStyle.fill
            ..color = Colors.yellow.withOpacity(0.1)
            ..strokeWidth = 1);

      Duration duration =
          DateTime.fromMicrosecondsSinceEpoch(selectionMax.toInt()).difference(
              DateTime.fromMicrosecondsSinceEpoch(selectionMin.toInt()));

      String durationStringCommon = duration.toString();
      String durationStringSeconds = "In Seconds: ${duration.inSeconds}";
      String durationStringMinutes =
          "In Minutes: ${(duration.inSeconds / 60).toStringAsFixed(1)}";
      String durationStringHours =
          "In Hours: ${(duration.inSeconds / 3600).toStringAsFixed(2)}";

      drawSelectionResizeArea(
          canvas,
          horScale.horValueToPixel(selectionMin) - selectionResizingPadding,
          size.height,
          selectionResizingPadding,
          true);
      drawSelectionResizeArea(canvas, horScale.horValueToPixel(selectionMax),
          size.height, selectionResizingPadding, false);

      double leftPadding = 10;
      drawText(
          canvas,
          leftPadding + horScale.horValueToPixel(selectionMin),
          0,
          timeToPixels(selectionMax - selectionMin),
          size.height,
          durationStringCommon,
          14,
          Colors.yellow,
          TextAlign.left,
          false);
      drawText(
          canvas,
          leftPadding + horScale.horValueToPixel(selectionMin),
          20,
          timeToPixels(selectionMax - selectionMin),
          size.height,
          durationStringSeconds,
          14,
          Colors.yellow,
          TextAlign.left,
          false);
      drawText(
          canvas,
          leftPadding + horScale.horValueToPixel(selectionMin),
          40,
          timeToPixels(selectionMax - selectionMin),
          size.height,
          durationStringMinutes,
          14,
          Colors.yellow,
          TextAlign.left,
          false);
      drawText(
          canvas,
          leftPadding + horScale.horValueToPixel(selectionMin),
          60,
          timeToPixels(selectionMax - selectionMin),
          size.height,
          durationStringHours,
          14,
          Colors.yellow,
          TextAlign.left,
          false);

      canvas.restore();
    }

    if (hoverPos.dy > 0 &&
        hoverPos.dx > 0 &&
        !selectionMovingIsStarted &&
        !selectionForZoomIsStarted &&
        !movingIsStarted) {
      canvas.drawLine(
          Offset(0, hoverPos.dy),
          Offset(size.width, hoverPos.dy),
          Paint()
            ..color = Colors.white38
            ..strokeWidth = 0.3);
      canvas.drawLine(
          Offset(hoverPos.dx, 0),
          Offset(hoverPos.dx, size.height),
          Paint()
            ..color = Colors.white38
            ..strokeWidth = 0.3);
    }

    canvas.drawRect(
        const Offset(0, 0) & Size(size.width, size.height),
        Paint()
          ..style = PaintingStyle.stroke
          ..color = DesignColors.fore1()
          ..strokeWidth = 1);
  }

*/
