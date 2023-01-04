package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func NewTitleComponent() *fyne.Container {
	title := canvas.NewText("DOKUMENTAR", color.RGBA{R: 104, G: 112, B: 132})
	title.TextStyle = fyne.TextStyle{
		Bold:      true,
		Monospace: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	subTitle := canvas.NewText("Paste your json data and comment it.", color.RGBA{R: 104, G: 112, B: 132})
	subTitle.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		&layout.Spacer{FixVertical: true},
		&layout.Spacer{FixVertical: true},
		container.New(
			layout.NewHBoxLayout(),
			layout.NewSpacer(),
			title,
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			subTitle,
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
			layout.NewSpacer(),
		),
		&layout.Spacer{FixVertical: true},
		&layout.Spacer{FixVertical: true},
		&layout.Spacer{FixVertical: true},
		&layout.Spacer{FixVertical: true},
	)
}
