package toppanel

import (
	"github.com/u00io/gazerstudio/system"
	"github.com/u00io/nui/nuikey"
	"github.com/u00io/nuiforms/ui"
)

type TopPanel struct {
	ui.Widget
}

func NewTopPanel() *TopPanel {
	var c TopPanel
	c.InitWidget()
	c.SetElevation(5)
	c.SetLayout(`
	<column>
		<row>
			<button id="btnOpenProject" text="Open Project" onclick="OnOpenProjectClick"/>
			<hspacer />
			<button id="btnStart" text="Start" onclick="OnStartClick"/>
			<button id="btnStop" text="Stop" onclick="OnStopClick"/>
		</row>
	</column>
	`, &c, nil)

	txtTarget, ok := c.FindWidgetByName("txtTarget").(*ui.TextBox)
	if ok {
		txtTarget.SetOnTextBoxKeyDown(func() {
			ev := ui.CurrentEvent().Parameter.(*ui.EventTextboxKeyDown)
			if ev.Key == nuikey.KeyEnter {
				c.OnStartClick()
				ev.Processed = true
				return
			}
			if ev.Key == nuikey.KeyEsc {
				c.OnStopClick()
				ev.Processed = true
				return
			}
		})
	}

	c.AddTimer(100, c.timerUpdate)
	return &c
}

func (c *TopPanel) timerUpdate() {
	c.updateButtons()
}

func (c *TopPanel) updateButtons() {
}

func (c *TopPanel) OnStartClick() {
	c.updateButtons()
}

func (c *TopPanel) OnStopClick() {
	c.updateButtons()
}

func (c *TopPanel) HandleSystemEvent(event system.Event) {
}
