package features

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type InputList struct {
	JsonEntries    []binding.String
	CommentEntries []binding.String
	Index          int
	Box            *fyne.Container
	AddButton      *widget.Button
}

func NewInputList() (il *InputList) {
	il = &InputList{
		Box: container.New(layout.NewVBoxLayout()),
	}
	il.setupButton()
	return
}

func (il *InputList) setupButton() {
	il.AddButton = widget.NewButton("+", il.addEntriesLine)
	il.addEntriesLine()
}

func (il *InputList) addEntriesLine() {
	json := binding.NewString()
	comment := binding.NewString()

	il.JsonEntries = append(il.JsonEntries, json)
	il.CommentEntries = append(il.CommentEntries, comment)

	jsonInput := widget.NewMultiLineEntry()
	commentInput := widget.NewMultiLineEntry()

	jsonInput.Bind(il.JsonEntries[il.Index])
	commentInput.Bind(il.CommentEntries[il.Index])

	jsonInput.TextStyle = fyne.TextStyle{
		Monospace: true,
	}

	jsonInput.PlaceHolder = "json"
	commentInput.PlaceHolder = "commentary"

	line := container.New(
		layout.NewGridLayout(2),
		jsonInput, commentInput,
	)

	il.Box.Add(line)
	il.Index += 1
}
