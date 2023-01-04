package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewGridHead(headers ...string) (grid *fyne.Container) {
	grid = container.New(layout.NewGridLayout(len(headers)))

	for _, h := range headers {
		c := container.New(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			canvas.NewText(h, color.RGBA{R: 64, G: 94, B: 94}),
			layout.NewSpacer(),
		)
		grid.Add(c)
	}

	return
}
