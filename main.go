package main

import (
	"image/color"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lekoller/dokumentar/features"
	"github.com/lekoller/dokumentar/my_theme"
)

var client *http.Client

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("DOKUMENTAR")
	win.Resize(fyne.NewSize(1200, 800))

	a.Settings().SetTheme(my_theme.MyTheme{})

	title := canvas.NewText("DOKUMENTAR", color.RGBA{R: 104, G: 112, B: 132})
	title.TextStyle = fyne.TextStyle{
		Bold:      true,
		Monospace: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	subTitle := canvas.NewText("Paste your json data and comment it.", color.RGBA{R: 104, G: 112, B: 132})
	subTitle.Alignment = fyne.TextAlignCenter

	il := features.NewInputList()
	tableHead := features.NewGridHead("paste your JSON here", "add your commentary here")
	// log.Println(space.Size())
	// space.Resize(fyne.Size{Width: 1, Height: 0.1})
	// log.Println(space.Size())

	titleBlock := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		title,
		layout.NewSpacer(),
		subTitle,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	)
	addBlock := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), il.AddButton, layout.NewSpacer())
	vBox := container.New(
		layout.NewVBoxLayout(),
		titleBlock,
		tableHead,
		il.Box,
		widget.NewSeparator(),
		addBlock,
	)

	win.SetContent(vBox)
	win.ShowAndRun()
}
