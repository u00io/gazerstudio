package modeproject

import (
	"math"
	"strconv"

	"github.com/u00io/gazerstudio/chart"
	"github.com/u00io/gazerstudio/project"
	"github.com/u00io/gazerstudio/utils"
	"github.com/u00io/nuiforms/ui"
)

type DataViewWidget struct {
	ui.Widget

	lvItems *ui.Table
	//dataItem  *project.DataItem
	lvResults *ui.Table

	chartSignal *chart.Chart
	chartFFT    *chart.Chart

	values []project.DataItemValue

	period float64
}

func NewDataViewWidget() *DataViewWidget {
	var c DataViewWidget
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)

	c.period = 1

	c.chartSignal = chart.NewChart()
	c.chartSignal.SetXExpandable(true)
	//c.chartSignal.Settings().HorizontalScale().SetScaleType(chart.HorizontalScaleTypeValue)
	c.chartFFT = chart.NewChart()
	c.chartFFT.SetXExpandable(true)
	c.chartFFT.Settings().HorizontalScale().SetScaleType(chart.HorizontalScaleTypeValue)

	customWidgets := map[string]ui.Widgeter{
		"chartSignal": c.chartSignal,
		"chartFFT":    c.chartFFT,
	}

	c.SetLayout(`
		<column>
			<button text="FFT" onclick="BtnFFT" />
			<button text="Update Data" onclick="UpdateData" />
			<button text="-" onclick="MinusData" />
			<button text="+" onclick="PlusData" />
			<spacer height="10" />
			<row>
				<table id="lvItems" />
				<widget id="chartSignal" />
			</row>
			<row>
				<table id="lvResults" />
				<widget id="chartFFT" />
			</row>
		</column>		
	`, &c, customWidgets)

	c.lvItems = c.FindWidgetByName("lvItems").(*ui.Table)
	c.lvItems.SetColumnCount(2)
	c.lvItems.SetColumnName(0, "DT")
	c.lvItems.SetColumnName(1, "Value")
	c.lvItems.SetColumnWidth(0, 300)
	c.lvItems.SetColumnWidth(1, 300)

	c.lvResults = c.FindWidgetByName("lvResults").(*ui.Table)
	c.lvResults.SetColumnCount(2)
	c.lvResults.SetColumnName(0, "Freq")
	c.lvResults.SetColumnName(1, "Magnitude")
	c.lvResults.SetColumnWidth(0, 300)
	c.lvResults.SetColumnWidth(1, 300)
	return &c
}

func (c *DataViewWidget) SetDataItemId(dataItem *project.DataItem) {
	c.lvItems.SetRowCount(0)
	if dataItem == nil {
		return
	}

	//c.dataItem = dataItem

	values := dataItem.GetValues()

	c.lvItems.SetRowCount(len(values))

	for i, v := range values {
		c.lvItems.SetCellText2(i, 0, strconv.FormatFloat(float64(v.DT), 'f', 6, 64))
		c.lvItems.SetCellText2(i, 1, v.Value)
	}

	c.updateChart()
	c.BtnFFT()
}

func (c *DataViewWidget) PlusData() {
	c.period += 1
	if c.period > 10 {
		c.period = 10
	}
	c.UpdateData()
}

func (c *DataViewWidget) MinusData() {
	c.period -= 1
	if c.period < 1 {
		c.period = 1
	}
	c.UpdateData()
}

func (c *DataViewWidget) UpdateData() {
	//dataItem := project.NewDataItem("0")
	values := make([]project.DataItemValue, 0)

	count := 4090

	for i := range count {
		valueRad := float64(i) / float64(count) * math.Pi * 2
		vSin := math.Sin(valueRad*c.period)*1000 + 1000
		vSin += math.Sin(valueRad*c.period*3) * 300
		vSin += math.Sin(valueRad*c.period*5) * 100
		dt := i * 1000000
		value := project.DataItemValue{
			DT:    int64(dt),
			Value: strconv.FormatFloat(vSin, 'f', 1, 64),
		}
		values = append(values, value)
	}
	c.values = values
	c.updateChart()
	c.BtnFFT()
}

func (c *DataViewWidget) updateChart() {
	///////////////////////////////////////
	// Signal
	c.chartSignal.Settings().RemoveAllAreas()
	area := chart.NewSettingsArea()
	series := chart.NewSettingsSeries()
	items := make([]*chart.DataItemValue, 0)
	minT := int64(0)
	maxT := int64(0)

	for i, v := range c.values {
		fv, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			continue
		}
		items = append(items, &chart.DataItemValue{
			DatetimeFirst: int64(i * 1000000),
			FirstValue:    fv,
			MinValue:      fv,
			MaxValue:      fv,
			LastValue:     fv,
			HasGood:       true,
			HasBad:        false,
			AvgValue:      fv,
		})

		if minT == 0 || int64(i*1000000) < minT {
			minT = int64(i * 1000000)
		}

		if int64(i*1000000) > maxT {
			maxT = int64(i * 1000000)
		}
	}
	series.SetData(items)
	area.AddSeries(series)
	c.chartSignal.Settings().AddArea(area)
	c.chartSignal.Settings().HorizontalScale().SetDefaultDisplayRange(float64(minT), float64(maxT))

	///////////////////////////////////////
	// FFT
	c.chartFFT.Settings().RemoveAllAreas()
	areaFFT := chart.NewSettingsArea()
	seriesFFT := chart.NewSettingsSeries()
	itemsFFT := make([]*chart.DataItemValue, 0)

	minT = 0
	maxT = 0

	for i := range c.lvResults.RowCount() {
		if i == 0 {
			continue
		}
		tStr := c.lvResults.GetCellText2(i, 0)
		t, _ := strconv.ParseFloat(tStr, 64)
		//t = t * 1000
		vStr := c.lvResults.GetCellText2(i, 1)
		v, _ := strconv.ParseFloat(vStr, 64)
		itemsFFT = append(itemsFFT, &chart.DataItemValue{
			DatetimeFirst: int64(t),
			DatetimeLast:  int64(t),
			FirstValue:    v,
			MinValue:      v,
			MaxValue:      v,
			LastValue:     v,
			HasGood:       true,
			HasBad:        false,
			AvgValue:      v,
		})

		if minT == 0 || int64(t) < minT {
			minT = int64(t)
		}
		if int64(t) > maxT {
			maxT = int64(t)
		}
	}
	seriesFFT.SetData(itemsFFT)
	areaFFT.AddSeries(seriesFFT)
	c.chartFFT.Settings().AddArea(areaFFT)
	maxT = minT + (maxT-minT)/50
	c.chartFFT.Settings().HorizontalScale().SetDefaultDisplayRange(float64(minT), float64(maxT))
}

func (c *DataViewWidget) FFT() []float64 {
	var data []float64
	for _, v := range c.values {
		fv, err := strconv.ParseFloat(v.Value, 64)
		if err != nil {
			continue
		}
		data = append(data, fv)
	}

	// add zero padding to the next power of two
	n := len(data)
	power := 1
	for power < n {
		power <<= 1
	}
	for len(data) < power {
		data = append(data, 0)
	}

	fft := utils.FFTDouble(data)
	return fft
}

func (c *DataViewWidget) BtnFFT() {
	fft := c.FFT()

	c.lvResults.SetRowCount(len(fft) / 2)

	sampleRate := float64(len(fft))
	n := len(fft)

	type Item struct {
		Freq      float64
		Magnitude float64
	}

	res := make([]Item, 0)

	for i := 0; i < n/2; i++ {
		freq := float64(i) * sampleRate / float64(n)
		magnitude := fft[i]
		magnitude = math.Abs(magnitude)
		res = append(res, Item{
			Freq:      freq,
			Magnitude: magnitude,
		})
		//c.lvResults.SetCellText2(i, 0, strconv.FormatFloat(freq, 'f', 6, 64))
		//c.lvResults.SetCellText2(i, 1, strconv.FormatFloat(magnitude, 'f', 6, 64))
	}

	/*sort.Slice(res, func(i, j int) bool {
		return res[i].Magnitude > res[j].Magnitude
	})*/

	for i, item := range res {
		c.lvResults.SetCellText2(i, 0, strconv.FormatFloat(item.Freq, 'f', 6, 64))
		c.lvResults.SetCellText2(i, 1, strconv.FormatFloat(item.Magnitude, 'f', 6, 64))
	}

	c.updateChart()
}
