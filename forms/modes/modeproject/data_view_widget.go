package modeproject

import (
	"sort"
	"strconv"

	"github.com/u00io/gazerstudio/project"
	"github.com/u00io/nuiforms/ui"
)

type DataViewWidget struct {
	ui.Widget

	lvItems   *ui.Table
	dataItem  *project.DataItem
	lvResults *ui.Table
}

func NewDataViewWidget() *DataViewWidget {
	var c DataViewWidget
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)
	c.SetLayout(`
		<column>
			<button text="FFT" onclick="BtnFFT" />
			<table id="lvItems" />
			<table id="lvResults" />
		</column>		
	`, &c, nil)

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

	c.dataItem = dataItem

	values := dataItem.GetValues()

	c.lvItems.SetRowCount(len(values))

	for i, v := range values {
		c.lvItems.SetCellText2(i, 0, v.DateTimeString())
		c.lvItems.SetCellText2(i, 1, v.Value)
	}
}

func (c *DataViewWidget) BtnFFT() {
	fft := c.dataItem.FFT()

	c.lvResults.SetRowCount(len(fft) / 2)

	sampleRate := 1000.0
	n := len(fft)

	type Item struct {
		Freq      float64
		Magnitude float64
	}

	res := make([]Item, 0)

	for i := 0; i < n/2; i++ {
		freq := float64(i) * sampleRate / float64(n)
		magnitude := fft[i]
		res = append(res, Item{
			Freq:      freq,
			Magnitude: magnitude,
		})
		//c.lvResults.SetCellText2(i, 0, strconv.FormatFloat(freq, 'f', 6, 64))
		//c.lvResults.SetCellText2(i, 1, strconv.FormatFloat(magnitude, 'f', 6, 64))
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Magnitude > res[j].Magnitude
	})

	for i, item := range res {
		c.lvResults.SetCellText2(i, 0, strconv.FormatFloat(item.Freq, 'f', 6, 64))
		c.lvResults.SetCellText2(i, 1, strconv.FormatFloat(item.Magnitude, 'f', 6, 64))
	}
}
