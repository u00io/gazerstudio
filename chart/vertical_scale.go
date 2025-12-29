package chart

import (
	"image/color"
	"math"
	"strconv"

	"github.com/u00io/nuiforms/ui"
)

const DefaultVerticalScaleWidth = 0.0
const DefaultVerticalScaleWidthInline = 50.0

type VerticalScale struct {
	xOffset int
	yOffset int
	width   int
	height  int

	verticalValuePadding01 float64

	targetDisplayedMinY float64
	targetDisplayedMaxY float64

	displayedMinY float64
	displayedMaxY float64
}

func NewVerticalScale() *VerticalScale {
	var c VerticalScale

	c.verticalValuePadding01 = 0.05

	c.targetDisplayedMinY = 0
	c.targetDisplayedMaxY = 1

	c.displayedMinY = 0
	c.displayedMaxY = 1

	return &c
}

func (c *VerticalScale) animation() {
	var diff = c.displayedMinY - c.targetDisplayedMinY
	//print("diff: $diff");
	c.displayedMinY -= diff * 0.1
	c.displayedMaxY -= diff * 0.1
}

func (c *VerticalScale) Calc(x, y, w, h int) {
	c.xOffset = x
	c.yOffset = y
	c.width = w
	c.height = h
}

/*
Canvas canvas, Size size, Color color, int index, bool showLegend,
      int totalCount
*/

func (c *VerticalScale) draw(cnv *ui.Canvas, color color.Color, index int, showLegend bool, totalCount int) {
	cnv.Save()
	if showLegend && totalCount > 1 {
		//cnv.TranslateAndClip(0, 0, )
	}

	cnv.FillRect(c.xOffset, c.yOffset, c.width, c.height, color)

	vertScalePointsCount := int(float64(c.height) / 70.0)

	verticalScale := c.getBeautifulScale(c.displayedMinY, c.displayedMaxY, vertScalePointsCount)
	for _, vertScaleItem := range verticalScale {
		posY := c.VerValueToPixel(vertScaleItem)
		if math.IsNaN(posY) {
			continue
		}
		cnv.FillRect(c.xOffset, c.yOffset+int(posY)-8, c.width, 20, ui.ColorFromHex("#55447755"))
		cnv.SetFontSize(12)
		cnv.SetColor(color)
		cnv.DrawText(c.xOffset, c.yOffset+int(posY)-8, c.width-5, 20, c.formatValue(vertScaleItem))
		cnv.SetColor(color)
		cnv.DrawLine(c.xOffset+c.width-3, c.yOffset+int(posY), c.xOffset+c.width, c.yOffset+int(posY), 1, color)
		cnv.DrawLine(c.xOffset+c.width-3, c.yOffset+int(posY), c.xOffset+c.width+cnv.Width(), c.yOffset+int(posY), 1, ui.ColorFromHex("#55774477"))
	}
	cnv.Restore()
}

/*
void draw(Canvas canvas, Size size, Color color, int index, bool showLegend,
      int totalCount) {
    canvas.save();
    if (showLegend && totalCount > 1) {
      canvas.clipRect(Rect.fromLTWH(xOffset, (index + 1) * 22, width, height));
    }
    canvas.drawRect(
        Rect.fromLTWH(xOffset, 0, width, height),
        Paint()
          ..style = PaintingStyle.fill
          ..strokeWidth = 1
          ..color = color.withOpacity(0.3));


    var vertScalePointsCount = (height / 70).round();

    var verticalScale =
        getBeautifulScale(displayedMinY, displayedMaxY, vertScalePointsCount);
    for (var vertScaleItem in verticalScale) {
      var posY = verValueToPixel(vertScaleItem);
      if (posY.isNaN) {
        continue;
      }
      canvas.drawRect(
          Rect.fromLTWH(xOffset, posY - 8, width, 20),
          Paint()
            ..style = PaintingStyle.fill
            ..strokeWidth = 1
            ..color = Colors.black.withOpacity(0.5));

      drawText(canvas, xOffset, posY - 8, width - 5, 20,
          formatValue(vertScaleItem), 12, color, TextAlign.right);
      canvas.drawLine(
          Offset(xOffset + width - 3, posY),
          Offset(xOffset + width, posY),
          Paint()
            ..style = PaintingStyle.stroke
            ..strokeWidth = 2
            ..color = color);

      canvas.drawLine(
          Offset(xOffset + width - 3, posY),
          Offset(xOffset + width + size.width, posY),
          Paint()
            ..style = PaintingStyle.stroke
            ..strokeWidth = 1
            ..color = color.withOpacity(0.2));
    }

    canvas.restore();
  }

*/

func (c *VerticalScale) UpdateVerticalScaleValues(history []*DataItemValue, united bool) {
	if !united {
		c.targetDisplayedMinY = 1.7976931348623157e+308 // max float64
		c.targetDisplayedMaxY = -1.7976931348623157e+308
	}
	for i := 0; i < len(history); i++ {
		value := history[i]
		if value.MinValue < c.targetDisplayedMinY {
			c.targetDisplayedMinY = value.MinValue
		}
		if value.MaxValue > c.targetDisplayedMaxY {
			c.targetDisplayedMaxY = value.MaxValue
		}
	}
	if c.targetDisplayedMinY != c.targetDisplayedMaxY {
		padding := (c.targetDisplayedMaxY - c.targetDisplayedMinY) * c.verticalValuePadding01
		c.targetDisplayedMinY = c.targetDisplayedMinY - padding
		c.targetDisplayedMaxY = c.targetDisplayedMaxY + padding
	} else {
		c.targetDisplayedMinY = c.targetDisplayedMinY - 1
		c.targetDisplayedMaxY = c.targetDisplayedMaxY + 1
	}

	// without animation
	c.displayedMinY = c.targetDisplayedMinY
	c.displayedMaxY = c.targetDisplayedMaxY
}

func (c *VerticalScale) ExpandToZero() {
	if c.targetDisplayedMinY == 1.7976931348623157e+308 ||
		c.targetDisplayedMaxY == -1.7976931348623157e+308 {
		return
	}
	if c.targetDisplayedMinY > 0 {
		c.targetDisplayedMinY = 0
	}
	if c.targetDisplayedMaxY < 0 {
		c.targetDisplayedMaxY = 0
	}
}

func (c *VerticalScale) formatValue(n float64) string {
	return strconv.FormatFloat(n, 'f', -1, 64)
}

func (c *VerticalScale) getBeautifulScale(min, max float64, countOfPoints int) []float64 {
	var scale []float64
	if max < min {
		return scale
	}
	if max == min {
		scale = append(scale, min)
		return scale
	}
	var diapason = max - min
	var step = diapason / float64(countOfPoints)
	doubleLog10 := func(x float64) float64 {
		const ln10 = 2.302585092994046
		return (math.Log(x) / ln10)
	}
	var log1 = math.Round(doubleLog10(step))
	var step10 = math.Pow(10, log1)
	for diapason/step10 < float64(countOfPoints) {
		step10 = step10 / 2
	}
	for newMin := min - math.Mod(min, step10); newMin < max; newMin += step10 {
		scale = append(scale, newMin)
	}
	return scale
}

func (c *VerticalScale) VerValueToPixel(value float64) float64 {
	var diapason = c.displayedMaxY - c.displayedMinY
	var offsetOfValueFromMin = value - c.displayedMinY
	var onePixelValue = float64(c.height) / diapason
	return float64(c.height) - onePixelValue*offsetOfValueFromMin
}

func (c *VerticalScale) VerPixelToValue(pixels float64) float64 {
	var diapason = c.displayedMaxY - c.displayedMinY
	var onePixelValue = float64(c.height) / diapason
	return pixels/onePixelValue + c.displayedMinY
}
