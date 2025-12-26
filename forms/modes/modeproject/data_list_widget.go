package modeproject

import (
	"github.com/u00io/gazerstudio/project"
	"github.com/u00io/gazerstudio/system"
	"github.com/u00io/nuiforms/ui"
)

type DataListWidget struct {
	ui.Widget
	lvItems *ui.Table

	OnSelected func(dataItem *project.DataItem)
}

func NewDataListWidget() *DataListWidget {
	var c DataListWidget
	c.InitWidget()
	c.SetLayout(`
		<column>
			<table id="lvItems" />
		</column>		
	`, &c, nil)

	c.SetMaxWidth(300)

	c.lvItems = c.FindWidgetByName("lvItems").(*ui.Table)
	c.lvItems.SetColumnCount(2)

	c.lvItems.SetColumnName(0, "Name")
	c.lvItems.SetColumnName(1, "Size")

	c.lvItems.SetColumnWidth(0, 180)
	c.lvItems.SetColumnWidth(1, 80)

	c.lvItems.SetRowCount(0)

	c.lvItems.SetOnSelectionChanged(func(row, col int) {
		if c.OnSelected != nil {
			project := system.Instance.GetProject()
			if project == nil {
				return
			}
			dataItems := project.DataItems()
			if row >= 0 && row < len(dataItems) {
				c.OnSelected(dataItems[row])
			} else {
				c.OnSelected(nil)
			}
		}
	})

	c.Load()

	return &c
}

func (c *DataListWidget) Load() {
	project := system.Instance.GetProject()
	if project == nil {
		return
	}

	dataItems := project.DataItems()

	c.lvItems.SetRowCount(len(dataItems))
	for i, dataItem := range dataItems {
		c.lvItems.SetCellText2(i, 0, dataItem.Name)
		c.lvItems.SetCellText2(i, 1, "0 B")
	}
}
