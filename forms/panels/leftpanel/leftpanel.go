package leftpanel

import "github.com/u00io/nuiforms/ui"

type LeftPanel struct {
	ui.Widget
}

func NewLeftPanel() *LeftPanel {
	var c LeftPanel
	c.InitWidget()
	c.SetMinWidth(50)
	c.SetElevation(2)
	c.SetAutoFillBackground(true)

	c.SetLayout(`
		<column>
			<button id="btnDataSources" text="Data Sources" />
			<button id="btnProcessing" text="Processing" />
			<button id="btnResults" text="Results" />
			<vspacer />
			<button id="btnData" text="Data" />
			<button id="btnSettings" text="Settings" />
		</column>
	`, &c, nil)
	return &c
}
