package chart

import (
	"image/color"

	"github.com/fogleman/gg"
	"github.com/u00io/nuiforms/ui"
)

type LoadingTask struct {
	minTime int
	maxTime int
}

func DrawLineOnCanvas(cnv *ui.Canvas, x1, y1, x2, y2 int, lineWidth float64, lineColor color.Color) {
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetLineWidth(lineWidth)
	ctx.SetColor(lineColor)
	ctx.DrawLine(float64(x1), float64(y1), float64(x2), float64(y2))
	ctx.Stroke()
}

func DrawRectOnCanvas(cnv *ui.Canvas, x, y, w, h int, lineWidth float64, lineColor color.Color) {
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetLineWidth(lineWidth)
	ctx.SetColor(lineColor)
	ctx.DrawRectangle(float64(x), float64(y), float64(w), float64(h))
	ctx.Stroke()
}

func FillRectOnCanvas(cnv *ui.Canvas, x, y, w, h int, fillColor color.Color) {
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetColor(fillColor)
	ctx.DrawRectangle(float64(x), float64(y), float64(w), float64(h))
	ctx.Fill()
}

func DrawCircleOnCanvas(cnv *ui.Canvas, x, y, r int, lineColor color.Color) {
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetLineWidth(1)
	ctx.SetColor(lineColor)
	ctx.DrawCircle(float64(x), float64(y), float64(r))
	ctx.Stroke()
}

func FillCircleOnCanvas(cnv *ui.Canvas, x, y, r int, fillColor color.Color) {
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetColor(fillColor)
	ctx.DrawCircle(float64(x), float64(y), float64(r))
	ctx.Fill()
}

type Point struct {
	X float64
	Y float64
}

func DrawPolygonOnCanvas(cnv *ui.Canvas, points []Point, lineWidth float64, lineColor color.Color) {
	if len(points) < 2 {
		return
	}
	ctx := gg.NewContextForRGBA(cnv.RGBA())
	ctx.Translate(float64(cnv.TranslatedX()), float64(cnv.TranslatedY()))
	ctx.SetLineWidth(lineWidth)
	ctx.SetColor(lineColor)
	ctx.MoveTo(points[0].X, points[0].Y)
	for i := 1; i < len(points); i++ {
		ctx.LineTo(points[i].X, points[i].Y)
	}
	ctx.Stroke()
}
