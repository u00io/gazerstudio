package chart

import (
	"fmt"
	"time"

	"github.com/u00io/nuiforms/ui"
)

type HorizontalScale struct {
	xOffset int
	yOffset int
	width   int
	height  int

	displayMin float64
	displayMax float64

	defaultDisplayMin float64
	defaultDisplayMax float64

	fixedHorScale bool
}

func NewHorizontalScale() *HorizontalScale {
	var c HorizontalScale
	c.height = 30
	return &c
}

func (c *HorizontalScale) ResetToDefaultDisplayRange() {
	c.displayMin = c.defaultDisplayMin
	c.displayMax = c.defaultDisplayMax
}

func (c *HorizontalScale) SetDefaultDisplayRange(min, max float64) {
	c.defaultDisplayMin = min
	c.defaultDisplayMax = max
	if !c.fixedHorScale {
		c.ResetToDefaultDisplayRange()
	}
}

func (c *HorizontalScale) SetDisplayRange(min, max float64) {
	c.displayMin = min
	c.displayMax = max
	if !c.fixedHorScale {
		c.ResetToDefaultDisplayRange()
	}
}

func (c *HorizontalScale) SetFixedHorScale(fixed bool) {
	c.fixedHorScale = fixed
	if !c.fixedHorScale {
		c.ResetToDefaultDisplayRange()
	}
}

func (c *HorizontalScale) Calc(x, y, w, h int) {
	c.xOffset = x
	c.yOffset = y
	c.width = w
	c.height = h
}

func (c *HorizontalScale) draw(cnv *ui.Canvas, width, height int) {
	if c.height < 1 {
		return
	}

	if c.xOffset < 0 || c.yOffset < 0 || c.width < 0 || c.height < 0 {
		return
	}

	cnv.SetFontSize(14)

	countOfValues := int64(c.width / 50)
	diapasonX := c.displayMax - c.displayMin

	dateTextWidth := 100
	displayDatesBlocks := true
	countOfDays := diapasonX / (24 * 3600 * 1000000)
	//maxCountOfDaysForDisplay := c.width / dateTextWidth

	/*if countOfDays > float64(maxCountOfDaysForDisplay) {
		displayDatesBlocks = false
	}*/

	beautifulScale := c.getHorBeautifulScale(int64(c.displayMin), int64(c.displayMax), countOfValues)

	for _, t := range beautifulScale {
		dt := time.UnixMicro(t).UTC()
		dateStr := dt.Format("2006-01-02")
		timeStr := dt.Format("15:04:05")
		ms := dt.Nanosecond() / 1000000
		msStr := ""

		if len(beautifulScale) > 1 {
			if beautifulScale[1]-beautifulScale[0] >= 60*1000000 {
				timeStr = dt.Format("15:04")
			}
			if beautifulScale[1]-beautifulScale[0] < 1000000 {
				msStr = fmt.Sprintf("%d ms", ms)
			}
		}

		_ = msStr

		if diapasonX > 0 {
			xPos := ((float64(t) - c.displayMin) / diapasonX) * float64(c.width)
			yPos := 0.0

			cnv.Save()
			//cnv.ClipRect(c.xOffset, c.yOffset, c.width, c.height)
			cnv.DrawLine(int(c.xOffset+int(xPos)), c.yOffset+int(yPos), int(c.xOffset+int(xPos)), c.yOffset+int(yPos)+5, 1, ui.ColorFromHex("#FFFF00"))

			yOffsetInScale := c.yOffset + 3

			timeWidth := 150

			cnv.SetColor(ui.ColorFromHex("#777777"))
			cnv.SetHAlign(ui.HAlignCenter)
			cnv.DrawText(int(c.xOffset+int(xPos))-timeWidth/2, yOffsetInScale, timeWidth, 100, timeStr)

			yOffsetInScale += 12
			if !displayDatesBlocks {
				cnv.SetColor(ui.ColorFromHex("#777777"))
				cnv.SetHAlign(ui.HAlignCenter)
				cnv.DrawText(int(c.xOffset+int(xPos))-timeWidth/2, yOffsetInScale, timeWidth, 100, dateStr)
			}
			cnv.Restore()
		}
	}

	cnv.Save()
	//cnv.ClipRect(c.xOffset, c.yOffset, c.width, c.height)

	if displayDatesBlocks && diapasonX > 0 {
		beautifulScaleForDates := c.getHorBeautifulScale(int64(c.displayMin), int64(c.displayMax), int64(countOfDays))
		off := c.yOffset + 14
		for _, d := range beautifulScaleForDates {
			dt := time.UnixMicro(d)
			currentColor := ui.ColorFromHex("#FFFFFF38")

			isToday := false
			dateNow := time.Now()
			if dateNow.Year() == dt.Year() && dateNow.Month() == dt.Month() && dateNow.Day() == dt.Day() {
				isToday = true
			}
			if isToday {
				currentColor = ui.ColorFromHex("#00FF00")
			}
			dateStr := dt.Format("2006-01-02")
			xPos1 := c.horValueToPixel(float64(d))
			xPos2 := c.horValueToPixel(float64(d + 24*3600*1000000))
			xPos1Visible := xPos1
			if xPos1Visible < 0 {
				xPos1Visible = 0
			}
			xPos2Visible := xPos2
			if xPos2Visible > float64(c.width) {
				xPos2Visible = float64(c.width)
			}
			cnv.DrawLine(int(c.xOffset+int(xPos1)), off+5, int(c.xOffset+int(xPos1)), off+20, 1, currentColor)
			cnv.DrawLine(int(c.xOffset+int(xPos2)), off+5, int(c.xOffset+int(xPos2)), off+20, 1, currentColor)
			cnv.DrawLine(int(c.xOffset+int(xPos1))+5, off+10, int(c.xOffset+int(xPos2))-5, off+10, 1, currentColor)
			dateTextHeight := 12
			dateTextPosX := xPos1Visible + (xPos2Visible-xPos1Visible)/2 - (float64(dateTextWidth) / 2)
			dateTextPosY := float64(off + 3)
			if dateTextPosX+float64(dateTextWidth) > float64(c.width) {
				dateTextPosX = float64(c.width - dateTextWidth)
			}
			if dateTextPosX < xPos1 {
				dateTextPosX = xPos1
			}
			if dateTextPosX < 0 {
				dateTextPosX = 0
			}
			if dateTextPosX+float64(dateTextWidth) > xPos2 {
				dateTextPosX = xPos2 - float64(dateTextWidth)
			}
			cnv.FillRect(int(c.xOffset+int(dateTextPosX)), int(dateTextPosY), dateTextWidth, dateTextHeight, ui.ColorFromHex("#00FF00"))
			cnv.SetColor(ui.ColorFromHex("#00FF00"))
			cnv.DrawRect(int(c.xOffset+int(dateTextPosX)), int(dateTextPosY), dateTextWidth, dateTextHeight)
			cnv.SetColor(ui.ColorFromHex("#FFFFFF"))
			cnv.DrawText(int(c.xOffset+int(dateTextPosX)), int(dateTextPosY)-1, dateTextWidth, dateTextHeight, dateStr)
		}
	}

	cnv.DrawLine(c.xOffset, c.yOffset, c.xOffset+c.width, c.yOffset, 1, ui.ColorFromHex("#777777"))
	cnv.Restore()
}

func (c *HorizontalScale) getHorBeautifulScale(min, max int64, countOfPoints int64) []int64 {
	scale := make([]int64, 0)
	if max < min {
		return scale
	}
	if max == min {
		scale = append(scale, int64(min))
		return scale
	}
	diapason := max - min
	step := int64(1)
	if countOfPoints != 0 {
		step = diapason / countOfPoints
	}
	newMin := min
	for i := 0; i < len(allowedSteps); i++ {
		st := allowedSteps[i]
		if step < st {
			step = st
			break
		}
	}
	newMin = newMin - int64(newMin)%int64(step)

	for i := int64(0); i < countOfPoints; i++ {
		if newMin > min && newMin < max {
			scale = append(scale, newMin)
		}
		newMin += int64(step)
	}
	return scale
}

var allowedSteps = []int64{
	1,
	5,
	10,
	50,
	100,
	500,
	1000,
	5000,
	10000,
	50000,
	100000,
	500000,
	1 * 1000000,
	2 * 1000000,
	5 * 1000000,
	10 * 1000000,
	15 * 1000000,
	30 * 1000000,
	1 * 60 * 1000000,
	2 * 60 * 1000000,
	5 * 60 * 1000000,
	10 * 60 * 1000000,
	15 * 60 * 1000000,
	30 * 60 * 1000000,
	1 * 60 * 60 * 1000000,
	3 * 60 * 60 * 1000000,
	6 * 60 * 60 * 1000000,
	12 * 60 * 60 * 1000000,
	1 * 24 * 3600 * 1000000,
	2 * 24 * 3600 * 1000000,
	7 * 24 * 3600 * 1000000,
	15 * 24 * 3600 * 1000000,
	1 * 30 * 24 * 3600 * 1000000,
	2 * 30 * 24 * 3600 * 1000000,
	3 * 30 * 24 * 3600 * 1000000,
	4 * 30 * 24 * 3600 * 1000000,
	365 * 24 * 3600 * 1000000,
}

func (c *HorizontalScale) horValueToPixel(value float64) float64 {
	diapason := c.displayMax - c.displayMin
	offsetOfValueFromMin := value - c.displayMin
	onePixelValue := float64(c.width) / diapason
	return onePixelValue*offsetOfValueFromMin + float64(c.xOffset)
}

func (c *HorizontalScale) horPixelToValue(pixels float64) float64 {
	pixels -= float64(c.xOffset)
	diapason := c.displayMax - c.displayMin
	onePixelValue := float64(c.width) / diapason
	return pixels/onePixelValue + c.displayMin
}

/*

  void resetToDefaultDisplayRange() {
    displayMin = defaultDisplayMin;
    displayMax = defaultDisplayMax;
  }

  void setDefaultDisplayRange(double min, double max) {
    defaultDisplayMin = min;
    defaultDisplayMax = max;
    if (!fixedHorScale) {
      resetToDefaultDisplayRange();
    }
  }


  double horValueToPixel(double time) {
    var diapason = displayMax - displayMin;
    var offsetOfValueFromMin = time - displayMin;
    var onePixelValue = width / diapason;
    return onePixelValue * offsetOfValueFromMin + xOffset;
  }

  double horPixelToValue(double pixels) {
    pixels -= xOffset;
    var diapason = displayMax - displayMin;
    var onePixelValue = width / diapason;
    return pixels / onePixelValue + displayMin;
  }


  List<int> getHorBeautifulScale(
      double min, double max, int countOfPoints, int minStep) {
    List<int> scale = [];

    if (max < min) {
      return scale;
    }

    if (max == min) {
      scale.add(min.toInt());
      return scale;
    }

    var diapason = max - min;
    int step = 1;
    if (countOfPoints != 0) {
      step = (diapason / countOfPoints).round();
    }
    var newMin = min;
    for (int i = 0; i < allowedSteps.length; i++) {
      var st = allowedSteps[i];
      if (st < minStep) {
        continue;
      }
      if (step < st) {
        step = st;
        break;
      }
    }
    newMin = newMin - (newMin % step);

    for (int i = 0; i < countOfPoints; i++) {
      if (newMin > min && newMin < max) {
        scale.add(newMin.toInt());
      }
      newMin += step;
    }

    return scale;
  }


  void draw(Canvas canvas, Size size) {
    if (height < 1) {
      return;
    }

    if (xOffset.isInfinite ||
        xOffset.isNaN ||
        yOffset.isInfinite ||
        yOffset.isNaN ||
        width.isInfinite ||
        width.isNaN ||
        height.isInfinite ||
        height.isNaN) {
      return;
    }

    international.DateFormat timeFormat = international.DateFormat("HH:mm:ss");
    international.DateFormat timeShortFormat =
        international.DateFormat("HH:mm");
    international.DateFormat dateFormat =
        international.DateFormat("yyyy-MM-dd");

    var countOfValues = width / 50;
    double diapasonX = (displayMax - displayMin).toDouble();

    double dateTextWidth = 100;
    var displayDatesBlocks = true;
    var countOfDays = diapasonX / (24 * 3600 * 1000000);
    var maxCountOfDaysForDisplay = width / dateTextWidth;
    if (countOfDays > maxCountOfDaysForDisplay) {
      displayDatesBlocks = false;
    }

    List<int> beautifulScale =
        getHorBeautifulScale(displayMin, displayMax, countOfValues.toInt(), 0);

    for (int t in beautifulScale) {
      DateTime dt = DateTime.fromMicrosecondsSinceEpoch(t);

      var dateStr = dateFormat.format(dt);
      var timeStr = timeFormat.format(dt);
      var ms = dt.millisecond;
      var msStr = "";

      if (beautifulScale.length > 1) {
        if (beautifulScale[1] - beautifulScale[0] >= 60 * 1000000) {
          timeStr = timeShortFormat.format(dt);
        }
        if (beautifulScale[1] - beautifulScale[0] < 1000000) {
          msStr = '${ms} ms';
        }
      }

      if (diapasonX > 0) {
        double xPos = ((t - displayMin) / diapasonX) * width;
        double yPos = 0;

        canvas.save();
        canvas.clipRect(Rect.fromLTWH(xOffset, yOffset, width, height));

        canvas.drawLine(
            Offset(xOffset + xPos, yOffset + yPos),
            Offset(xOffset + xPos, yOffset + yPos + 5),
            Paint()
              ..style = PaintingStyle.stroke
              ..color = Colors.yellow
              ..strokeWidth = 3);

        var yOffsetInScale = yOffset + 3;

        double timeWidth = 150;

        drawText(canvas, xPos + xOffset - timeWidth / 2, yPos + yOffsetInScale,
            timeWidth, 100, timeStr, 10, Colors.blueAccent, TextAlign.center);
        yOffsetInScale += 12;

        if (!displayDatesBlocks) {
          drawText(
              canvas,
              xPos + xOffset - timeWidth / 2,
              yPos + yOffsetInScale,
              timeWidth,
              100,
              dateStr,
              12,
              Colors.blueAccent,
              TextAlign.center);
        }
        canvas.restore();
      }
    }

    canvas.save();
    canvas.clipRect(Rect.fromLTWH(xOffset, yOffset, width, height));

    if (displayDatesBlocks && diapasonX > 0) {
      var beautifulScaleForDates =
          getBeautifulScaleForDates(displayMin.toInt(), displayMax.toInt());
      var off = yOffset + 14;
      for (int d in beautifulScaleForDates) {
        DateTime dt = DateTime.fromMicrosecondsSinceEpoch(d);

        Color currentColor = Colors.white38;

        var isToday = false;
        var dateNow = DateTime.now();
        if (dateNow.year == dt.year &&
            dateNow.month == dt.month &&
            dateNow.day == dt.day) {
          isToday = true;
        }

        if (isToday) {
          currentColor = Colors.green;
        }

        var dateStr = dateFormat.format(dt);
        var xPos1 = horValueToPixel(d.toDouble());
        var xPos2 = horValueToPixel(d + 24 * 3600 * 1000000);
        var xPos1Visible = xPos1;
        if (xPos1Visible < 0) {
          xPos1Visible = 0;
        }
        var xPos2Visible = xPos2;
        if (xPos2Visible > width) {
          xPos2Visible = width;
        }
        canvas.drawLine(
            Offset(xPos1 + 2, off + 5),
            Offset(xPos1 + 2, off + 20),
            Paint()
              ..style = PaintingStyle.stroke
              ..color = currentColor
              ..strokeWidth = 1);
        canvas.drawLine(
            Offset(xPos2 - 2, off + 5),
            Offset(xPos2 - 2, off + 20),
            Paint()
              ..style = PaintingStyle.stroke
              ..color = currentColor
              ..strokeWidth = 1);
        canvas.drawLine(
            Offset(xPos1 + 5, off + 10),
            Offset(xPos2 - 5, off + 10),
            Paint()
              ..style = PaintingStyle.stroke
              ..color = currentColor
              ..strokeWidth = 1);

        double dateTextHeight = 12;
        var dateTextPosX = xPos1Visible +
            (xPos2Visible - xPos1Visible) / 2 -
            (dateTextWidth / 2);
        var dateTextPosY = off + 3;

        if (dateTextPosX + dateTextWidth > width) {
          dateTextPosX = width - dateTextWidth;
        }

        if (dateTextPosX < xPos1) {
          dateTextPosX = xPos1;
        }

        if (dateTextPosX < 0) {
          dateTextPosX = 0;
        }
        if (dateTextPosX + dateTextWidth > xPos2) {
          dateTextPosX = xPos2 - dateTextWidth;
        }
        //(this.left, this.top, this.right, this.bottom
        canvas.drawRect(
            Rect.fromLTWH(
                dateTextPosX, dateTextPosY, dateTextWidth, dateTextHeight),
            Paint()
              ..style = PaintingStyle.fill
              ..color = Colors.green);
        canvas.drawRect(
            Rect.fromLTWH(
                dateTextPosX, dateTextPosY, dateTextWidth, dateTextHeight),
            Paint()
              ..style = PaintingStyle.stroke
              ..color = Colors.green);

        drawText(canvas, dateTextPosX, dateTextPosY - 1, dateTextWidth,
            dateTextHeight, dateStr, 10, Colors.white, TextAlign.center);
      }
    }

    canvas.drawLine(
        Offset(xOffset, yOffset),
        Offset(xOffset + width, yOffset),
        Paint()
          ..style = PaintingStyle.stroke
          ..color = Colors.yellow
          ..strokeWidth = 1);

    canvas.restore();
  }

*/
