package main

import (
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/leandro-koller-bft/doku/features"
	"github.com/leandro-koller-bft/doku/requests"
)

var client *http.Client

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("Get Useless Fact")
	win.Resize(fyne.NewSize(1200, 800))

	title := canvas.NewText("Get Your Useless Facts", color.RGBA{R: 83, G: 89, B: 147})
	title.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	progress := features.NewProgressBar()

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	button := widget.NewButton("Get Fact", func() {
		progress.Start()
		fact, err := requests.GetRandomFact(client)
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		progress.Finish()
		factText.SetText(fact.Text)
	})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), progress.Bar, factText)

	win.SetContent(vBox)
	win.ShowAndRun()
}
