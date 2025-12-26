package centerpanel

import (
	_ "embed"

	"github.com/u00io/gazerstudio/forms/panels/contentpanel"
	"github.com/u00io/gazerstudio/forms/panels/leftpanel"
	"github.com/u00io/nuiforms/ui"
)

type CenterPanel struct {
	ui.Widget

	leftPanel    *leftpanel.LeftPanel
	contentPanel *contentpanel.ContentPanel
}

func NewCenterPanel() *CenterPanel {
	var c CenterPanel
	c.InitWidget()

	c.leftPanel = leftpanel.NewLeftPanel()
	c.contentPanel = contentpanel.NewContentPanel()

	curstomWidgets := map[string]ui.Widgeter{
		"leftpanel":    c.leftPanel,
		"contentpanel": c.contentPanel,
	}
	c.SetLayout(`
		<row>
			<widget id="leftpanel" />
			<widget id="contentpanel" />
		</row>
	`, &c, curstomWidgets)

	return &c
}
