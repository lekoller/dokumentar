package components

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type ItemData struct {
	Id           int
	Entity       binding.String
	ConnType     *string
	Method       *string
	Endpoint     binding.String
	JsonEntry    binding.String
	CommentEntry binding.String
	Deleted      bool
}

func (i *ItemData) delete() {
	i.Deleted = true
}

type InputList struct {
	Items     []*ItemData
	Index     int
	Box       *fyne.Container
	AddButton *widget.Button
}

func NewInputList() (il *InputList) {
	il = &InputList{
		Box: container.New(layout.NewVBoxLayout()),
	}
	il.setupButton()
	il.addEntriesLine()
	return
}

func (il *InputList) Clear() {
	il.Items = []*ItemData{}
	il.Index = 0
	il.Box.RemoveAll()
	il.setupButton()
	il.addEntriesLine()
}

func (il *InputList) Remove(id int, row *fyne.Container) {
	for _, item := range il.Items {
		if item.Id == id {
			item.delete()
		}
	}
	il.Box.Remove(row)
}

func (il *InputList) setupButton() {
	il.AddButton = widget.NewButton("add item", il.addEntriesLine)
	il.AddButton.Icon = theme.ContentAddIcon()
}

func (il *InputList) addEntriesLine() {
	var row *fyne.Container
	var connType string
	var method string
	entity := binding.NewString()
	endpoint := binding.NewString()
	json := binding.NewString()
	comment := binding.NewString()

	httpMethodInput := widget.NewSelect([]string{"GET", "POST", "PUT", "PATCH", "DELETE"}, func(s string) {
		method = s
	})
	httpMethodInput.Disable()
	connectionTypeInput := widget.NewSelect([]string{"HTTP", "gRPC", "Topic (Queue)"}, func(s string) {
		if s == "HTTP" {
			httpMethodInput.Enable()
		} else {
			httpMethodInput.Disable()
		}
		connType = s
	})
	entityInput := widget.NewEntry()
	endpointInput := widget.NewEntry()
	jsonInput := widget.NewMultiLineEntry()
	commentInput := widget.NewMultiLineEntry()

	data := &ItemData{
		Id:           il.Index + 1,
		Entity:       entity,
		ConnType:     &connType,
		Method:       &method,
		Endpoint:     endpoint,
		JsonEntry:    json,
		CommentEntry: comment,
	}
	deleteButton := widget.NewButton("", func() {
		il.Remove(data.Id, row)
	})
	deleteButton.Icon = theme.CancelIcon()

	il.Items = append(il.Items, data)

	entityInput.Bind(il.Items[il.Index].Entity)
	endpointInput.Bind(il.Items[il.Index].Endpoint)
	jsonInput.Bind(il.Items[il.Index].JsonEntry)
	commentInput.Bind(il.Items[il.Index].CommentEntry)

	jsonInput.TextStyle = fyne.TextStyle{
		Monospace: true,
	}
	connectionTypeInput.PlaceHolder = "Connection Type"
	httpMethodInput.PlaceHolder = "HTTP Method"
	entityInput.SetPlaceHolder("Entity Name")
	endpointInput.SetPlaceHolder("/detail/endpoint")
	jsonInput.SetPlaceHolder("json")
	commentInput.SetPlaceHolder("commentary")

	headGrid := container.New(
		layout.NewGridLayout(3),
		container.NewPadded(canvas.NewText("  Item #"+strconv.Itoa(data.Id), color.RGBA{R: 104, G: 112, B: 132})),
		container.NewPadded(entityInput),
		container.NewPadded(container.NewHBox(layout.NewSpacer(), deleteButton)),
	)
	singleLinesGrid := container.New(
		layout.NewGridLayout(3),
		container.NewPadded(connectionTypeInput),
		container.NewPadded(httpMethodInput),
		container.NewPadded(endpointInput),
	)
	multiLinesGrid := container.New(
		layout.NewGridLayout(2),
		container.NewPadded(jsonInput),
		container.NewPadded(commentInput),
	)

	row = container.New(
		layout.NewVBoxLayout(),
		headGrid,
		singleLinesGrid,
		multiLinesGrid,
		container.NewPadded(widget.NewSeparator()),
	)

	il.Box.Add(row)
	il.Index += 1
}
