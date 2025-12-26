package bottompanel

import (
	"time"

	"github.com/u00io/gazerstudio/system"
	"github.com/u00io/nuiforms/ui"
)

type BottomPanel struct {
	ui.Widget

	dtButtonCopyChanged time.Time
}

func NewBottomPanel() *BottomPanel {
	var c BottomPanel
	c.InitWidget()
	c.SetLayout(`
		<row>
			<hspacer />
			<button text="About" onclick="OnAboutClicked" />
		</row>
	`, &c, nil)

	c.SetElevation(5)

	c.AddTimer(1000, c.OnTimerUpdate)
	return &c
}

func (c *BottomPanel) HandleSystemEvent(event system.Event) {
}

func (c *BottomPanel) OnAboutClicked() {
	ui.ShowAboutDialog("About", "Gazer Studio v0.2.2", "", "", "")
}

func (c *BottomPanel) OnTimerUpdate() {
	if time.Since(c.dtButtonCopyChanged) > time.Second {
		btnCopy, ok := c.FindWidgetByName("btnCopy").(*ui.Button)
		if ok {
			btnCopy.SetText("Copy")
			btnCopy.SetBackgroundColor(nil)
		}
	}
}
