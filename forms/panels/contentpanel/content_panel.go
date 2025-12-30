package contentpanel

import (
	"github.com/u00io/gazerstudio/forms/modes/modeproject"
	"github.com/u00io/nuiforms/ui"
)

type ContentPanel struct {
	ui.Widget

	projectWidget *modeproject.ProjectWidget
	//chart *chart.Chart
}

func NewContentPanel() *ContentPanel {
	var c ContentPanel
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)

	c.projectWidget = modeproject.NewProjectWidget()
	//c.chart = chart.NewChart()

	customWidgets := map[string]ui.Widgeter{
		//"chart": c.chart,
		"projectWidget": c.projectWidget,
	}
	c.SetLayout(`
		<column>
			<widget id="projectWidget" />
		</column>		
	`, &c, customWidgets)
	return &c
}
