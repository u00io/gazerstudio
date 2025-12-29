package contentpanel

import (
	"github.com/u00io/gazerstudio/chart"
	"github.com/u00io/nuiforms/ui"
)

type ContentPanel struct {
	ui.Widget

	//projectWidget *modeproject.ProjectWidget
	chart *chart.Chart
}

func NewContentPanel() *ContentPanel {
	var c ContentPanel
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)

	//c.projectWidget = modeproject.NewProjectWidget()
	c.chart = chart.NewChart()

	customWidgets := map[string]ui.Widgeter{
		"chart": c.chart,
	}
	c.SetLayout(`
		<column>
			<widget id="chart" />
		</column>		
	`, &c, customWidgets)
	return &c
}
