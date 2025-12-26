package modeproject

import (
	"github.com/u00io/gazerstudio/project"
	"github.com/u00io/nuiforms/ui"
)

type ProjectWidget struct {
	ui.Widget

	dataListWidget *DataListWidget
	dataViewWidget *DataViewWidget
}

func NewProjectWidget() *ProjectWidget {
	var c ProjectWidget
	c.InitWidget()

	c.dataListWidget = NewDataListWidget()
	c.dataViewWidget = NewDataViewWidget()

	c.SetXExpandable(true)
	c.SetYExpandable(true)

	customWidgets := map[string]ui.Widgeter{
		"dataListWidget": c.dataListWidget,
		"dataViewWidget": c.dataViewWidget,
	}
	c.SetLayout(`
		<row>
			<widget id="dataListWidget" />
			<widget id="dataViewWidget" />
		</row>
	`, &c, customWidgets)

	c.dataListWidget.OnSelected = func(dataItem *project.DataItem) {
		c.dataViewWidget.SetDataItemId(dataItem)
	}

	return &c
}
