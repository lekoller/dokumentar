package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ControlPanel struct {
	InputList   *InputList
	ProjectInfo *ProjectInfo
	Box         *fyne.Container
}

func NewControlPanel(info *ProjectInfo, list *InputList) (control *ControlPanel) {
	control = &ControlPanel{
		InputList:   list,
		ProjectInfo: info,
		Box:         container.New(layout.NewVBoxLayout(), container.NewPadded(widget.NewSeparator())),
	}
	control.setup()
	return
}

func (c *ControlPanel) setup() {
	renderButton := widget.NewButton("render documentation", mountRenderCallback(c))
	renderButton.Icon = theme.ConfirmIcon()

	clearButton := widget.NewButton("clear", func() {
		c.InputList.Clear()
	})
	clearButton.Icon = theme.DeleteIcon()

	horizontal := container.NewHBox(
		layout.NewSpacer(),
		container.NewPadded(clearButton),
		container.NewPadded(renderButton),
	)
	c.Box.Add(horizontal)
}
