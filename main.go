package main

import (
	"image/color"
	"log"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
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
		Bold: true,
	}
	title.Alignment = fyne.TextAlignCenter
	title.TextSize = 24

	factText := widget.NewLabel("")
	factText.Wrapping = fyne.TextWrapWord

	var jsonContents []binding.String
	var jsonIndex int

	box := container.New(layout.NewVBoxLayout())
	button := widget.NewButton("+", func() {
		jsonContent := binding.NewString()
		jsonContents = append(jsonContents, jsonContent)

		jsonInput := widget.NewMultiLineEntry()
		println(jsonIndex)
		log.Println(jsonContents)
		jsonInput.Bind(jsonContents[jsonIndex])
		box.Add(jsonInput)
		jsonIndex++
	})

	hBox := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), button, layout.NewSpacer())
	vBox := container.New(layout.NewVBoxLayout(), title, hBox, widget.NewSeparator(), box)

	win.SetContent(vBox)
	win.ShowAndRun()
}
