package main

import (
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
	tableHead := components.NewGridHead("Project Name", "Container Name", "Module Name")

	titleBlock := components.NewTitleComponent()
	project := components.NewProjectInfo()
	addBlock := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), il.AddButton, layout.NewSpacer())
	control := components.NewControlPanel(project, il)
	vBox := container.New(
		layout.NewVBoxLayout(),
		titleBlock,
		tableHead,
		project.Box,
		il.Box,
		addBlock,
		layout.NewSpacer(),
		control.Box,
	)
	scrollBox := container.NewVScroll(
		container.NewPadded(
			container.NewPadded(
				container.NewPadded(
					container.NewPadded(
						container.NewPadded(
							container.NewPadded(
								container.NewPadded(
									container.NewPadded(container.NewPadded(vBox)),
								),
							),
						),
					),
				),
			),
		),
	)

	win.SetContent(scrollBox)
	win.ShowAndRun()
}
