package main

import (
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/lekoller/dokumentar/components"
	"github.com/lekoller/dokumentar/my_theme"
)

var client *http.Client

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	a := app.New()
	win := a.NewWindow("DOKUMENTAR")
	win.Resize(fyne.NewSize(1200, 800))

	a.Settings().SetTheme(my_theme.MyTheme{})

	il := components.NewInputList()
	tableHead := components.NewGridHead("paste your JSON here", "add your commentary here")

	titleBlock := components.NewTitleComponent()
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
