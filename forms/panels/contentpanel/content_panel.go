package contentpanel

import (
	"github.com/u00io/gazerstudio/forms/modes/modeproject"
	"github.com/u00io/nuiforms/ui"
)

type ContentPanel struct {
	ui.Widget

	projectWidget *modeproject.ProjectWidget
}

func NewContentPanel() *ContentPanel {
	var c ContentPanel
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)

	c.projectWidget = modeproject.NewProjectWidget()

	customWidgets := map[string]ui.Widgeter{
		"projectWidget": c.projectWidget,
	}
	c.SetLayout(`
		<column>
			<widget id="projectWidget" />
		</column>		
	`, &c, customWidgets)
	return &c
}
