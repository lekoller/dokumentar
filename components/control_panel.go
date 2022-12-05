package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ControlPanel struct {
	DataList    []*ItemData
	ProjectInfo *ProjectInfo
	Box         *fyne.Container
}

func NewControlPanel(info *ProjectInfo, list []*ItemData) (control *ControlPanel) {
	control = &ControlPanel{
		DataList:    list,
		ProjectInfo: info,
		Box:         container.New(layout.NewVBoxLayout(), container.NewPadded(widget.NewSeparator())),
	}
	control.setup()
	return
}

func (c *ControlPanel) setup() {
	renderButton := widget.NewButton("render documentation", func() {})
	renderButton.Icon = theme.ConfirmIcon()

	horizontal := container.NewHBox(layout.NewSpacer(), container.NewPadded(renderButton))
	c.Box.Add(horizontal)
}
