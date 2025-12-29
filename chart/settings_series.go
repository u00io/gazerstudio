package chart

import "github.com/u00io/nuiforms/ui"

type SettingsSeries struct {
	xOffset               int
	yOffset               int
	yOffsetOfHeader       float64
	width                 int
	height                int
	verticalScaleWidth111 float64
	selected              bool
	displayName           string

	vScale       *VerticalScale
	itemHistory  []*DataItemValue
	loadingTasks []*LoadingTask

	// Props
	showZero bool
}

func NewSettingsSeries() *SettingsSeries {
	var c SettingsSeries
	return &c
}

func (c *SettingsSeries) Calc(x, y, w, h, vsWidth int, vSc *VerticalScale, yHeaderOffset float64) {
	c.xOffset = x
	c.yOffset = y
	c.width = w
	c.height = h
	c.yOffsetOfHeader = yHeaderOffset
	c.vScale = vSc
}

func (c *SettingsSeries) draw(cnv *ui.Canvas, width int, height int, hScale *HorizontalScale, settings *Settings, smooth bool, index, totalSeriesCount int) {

	// DrawLineOnCanvas(cnv, 0, 0, 100, 100, 1, ui.ColorFromHex("#654654"))

	cnv.Save()
	//cnv.ClipRect(c.xOffset+int(c.verticalScaleWidth111), 0, c.width-int(c.verticalScaleWidth111), c.height)

	c.vScale.animation()

	var points []Point
	var pointsQuality []Point
	var lastHasGood bool
	var lastHasBad bool

	var funcDrawPoints = func() {
		if len(points) == 1 {
			FillCircleOnCanvas(cnv, int(points[0].X), int(points[0].Y), 1, ui.ColorFromHex("#FF0000"))
			//canvas.drawPoints(PointMode.points, points, paint);
		} else {
			DrawPolygonOnCanvas(cnv, points, 1, ui.ColorFromHex("#FF0000"))
		}
		points = []Point{}
	}

	for i := 0; i < len(c.itemHistory); i++ {
		firstPoint := i == 0
		item := c.itemHistory[i]
		posX := hScale.horValueToPixel(float64(item.DatetimeFirst))
		if item.HasGood {
			if lastHasGood || firstPoint {
				points = append(points, Point{X: posX + float64(c.xOffset), Y: c.vScale.VerValueToPixel(item.FirstValue)})
			}
			if item.MinValue != item.FirstValue {
				points = append(points, Point{X: posX + float64(c.xOffset), Y: c.vScale.VerValueToPixel(item.MinValue)})
			}
			if item.MaxValue != item.MinValue {
				points = append(points, Point{X: posX + float64(c.xOffset), Y: c.vScale.VerValueToPixel(item.MaxValue)})
			}
			if item.LastValue != item.MaxValue {
				points = append(points, Point{X: posX + float64(c.xOffset), Y: c.vScale.VerValueToPixel(item.LastValue)})
			}
		} else {
			funcDrawPoints()
		}
		if !item.HasBad {
			if len(pointsQuality) == 0 || (item.HasBad != lastHasBad || i == len(c.itemHistory)-1) {
				pointsQuality = append(pointsQuality, Point{X: posX + float64(c.xOffset), Y: c.yOffsetOfHeader})
			}
		} else {
			if len(pointsQuality) == 0 || (item.HasBad != lastHasBad) || i == len(c.itemHistory)-1 {
				pointsQuality = append(pointsQuality, Point{X: posX + float64(c.xOffset), Y: c.yOffsetOfHeader})
			}
		}
		if item.HasBad != lastHasBad {
		}
		lastHasGood = item.HasGood
		lastHasBad = item.HasBad
	}

	funcDrawPoints()

	cnv.Restore()

}

func (c *SettingsSeries) drawDetails(cnv *ui.Canvas, width int, height int, hScale *HorizontalScale, settings *Settings, smooth bool, index, totalSeriesCount int) {
	/*cnv.SetFontSize(16)
	cnv.SetColor(ui.ColorFromHex("#887744"))
	cnv.DrawText(c.xOffset+10, int(c.yOffsetOfHeader), c.width, c.height, "DETAILS")*/
}

func (c *SettingsSeries) ShowZero() bool {
	return c.showZero
}

/*
void draw(
      Canvas canvas,
      Size s,
      TimeChartHorizontalScale hScale,
      TimeChartSettings settings,
      bool smooth,
      int index,
      int totalSeriesCount) {
    List<DataItemHistoryChartItemValueResponse> history = itemHistory;

    canvas.save();
    canvas.clipRect(Rect.fromLTWH(xOffset + verticalScaleWidth111, 0,
        width - verticalScaleWidth111, height));

    {
      vScale.animation();

      var paint = Paint()
        ..style = PaintingStyle.stroke
        ..color = getColor("stroke_color")
        ..strokeJoin = StrokeJoin.round
        ..strokeWidth = getDouble("stroke_width");

      var paintQualityGood = Paint()
        ..style = PaintingStyle.stroke
        ..color = Colors.green
        ..strokeJoin = StrokeJoin.round
        ..strokeWidth = 0.2;

      var paintQualityBad = Paint()
        ..style = PaintingStyle.stroke
        ..color = Colors.red
        ..strokeJoin = StrokeJoin.round
        ..strokeCap = StrokeCap.round
        ..strokeWidth = 5;

      var paintLoading = Paint()
        ..style = PaintingStyle.stroke
        ..color = Colors.lightBlue
        ..strokeJoin = StrokeJoin.round
        ..strokeWidth = 1;

      List<Offset> points = [];
      List<Offset> pointsQuality = [];
      bool lastHasGood = false;
      bool lastHasBad = false;

      void funcDrawPoints() {
        //print("points: ${points.length} hislen: ${history.length}");
        if (points.length == 1) {
          canvas.drawCircle(points[0], 1, paint..style = PaintingStyle.fill);
          //canvas.drawPoints(PointMode.points, points, paint);
        } else {
          if (smooth) {
            canvas.drawPoints(PointMode.polygon, points,
                paint..color = paint.color.withOpacity(0.2));
          } else {
            canvas.drawPoints(PointMode.polygon, points, paint);
          }
        }
        points = [];
      }

      void funcDrawPointsQuality() {
        var currentPaint = paintQualityGood;
        if (lastHasBad) {
          currentPaint = paintQualityBad;
        }

        canvas.drawPoints(PointMode.polygon, pointsQuality, currentPaint);
        pointsQuality = [];
      }

      for (int i = 0; i < history.length; i++) {
        bool firstPoint = i == 0;
        var item = history[i];

        var posX = hScale.horValueToPixel(item.datetimeFirst.toDouble());
        if (item.hasGood) {
          if (lastHasGood || firstPoint) {
            points.add(Offset(posX, vScale.verValueToPixel(item.firstValue)));
          }
          if (item.minValue != item.firstValue) {
            points.add(Offset(posX, vScale.verValueToPixel(item.minValue)));
          }
          if (item.maxValue != item.minValue) {
            points.add(Offset(posX, vScale.verValueToPixel(item.maxValue)));
          }
          if (item.lastValue != item.maxValue) {
            points.add(Offset(posX, vScale.verValueToPixel(item.lastValue)));
          }
        } else {
          funcDrawPoints();
        }

        if (!item.hasBad) {
          if (pointsQuality.isEmpty ||
              (item.hasBad != lastHasBad || i == history.length - 1)) {
            pointsQuality.add(Offset(posX, yOffsetOfHeader));
          }
        } else {
          if (pointsQuality.isEmpty ||
              (item.hasBad != lastHasBad) ||
              i == history.length - 1) {
            pointsQuality.add(Offset(posX, yOffsetOfHeader));
          }
        }

        if (item.hasBad != lastHasBad) {
          funcDrawPointsQuality();
        }

        lastHasGood = item.hasGood;
        lastHasBad = item.hasBad;
      }

      funcDrawPoints();
      funcDrawPointsQuality();

      if (loadingTasks.isNotEmpty) {
        // draw loading
        for (var loadingTask in loadingTasks) {
          List<Offset> pointsLoadingTasks = [];
          var posX1 = hScale.horValueToPixel(loadingTask.minTime.toDouble());
          var posX2 = hScale.horValueToPixel(loadingTask.maxTime.toDouble());
          pointsLoadingTasks.add(Offset(posX1, height - 5));
          pointsLoadingTasks.add(Offset(posX2, height - 5));
          canvas.drawPoints(
              PointMode.polygon, pointsLoadingTasks, paintLoading);
        }
      }
    }

    //drawText(canvas, 0, 0, width - verticalScaleWidth - 10, 20, itemName, 14, Colors.yellowAccent, TextAlign.right);

    canvas.restore();

    final f = international.NumberFormat("#.##########");
    String formatValue(num n) {
      return f.format(n);
    }
  }

*/
