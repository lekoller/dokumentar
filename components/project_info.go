package components

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type ProjectInfo struct {
	ProjectName   binding.String
	ContainerName binding.String
	ModuleName    binding.String
	Box           *fyne.Container
}

func NewProjectInfo() *ProjectInfo {
	projectName := binding.NewString()
	containerName := binding.NewString()
	moduleName := binding.NewString()

	projectInput := widget.NewEntry()
	containerInput := widget.NewEntry()
	moduleInput := widget.NewEntry()

	projectInput.Bind(projectName)
	containerInput.Bind(containerName)
	moduleInput.Bind(moduleName)

	projectInput.SetPlaceHolder("define project name")
	containerInput.SetPlaceHolder("define software unit name")
	moduleInput.SetPlaceHolder("define where, inside the software")

	box := container.New(
		layout.NewGridLayout(3),
		container.NewPadded(projectInput),
		container.NewPadded(containerInput),
		container.NewPadded(moduleInput),
	)
	return &ProjectInfo{
		ProjectName:   projectName,
		ContainerName: containerName,
		ModuleName:    moduleName,
		Box:           container.NewVBox(box, container.NewPadded(widget.NewSeparator())),
	}
}
