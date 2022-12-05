package templates

import (
	"html/template"
	"os"
)

type BuilderData struct {
	ProjectName   string
	ContainerName string
	ModuleName    string
	List          []BuilderListItem
}

type BuilderListItem struct {
	Entity       string
	ConnType     string
	Method       string
	Endpoint     string
	JsonEntry    string
	CommentEntry string
}

func BuildTemplate(data BuilderData) {
	tmp_list := []string{
		"header.html",
		"body.html",
	}
	t := template.New("body.html")
	// t.Funcs(template.FuncMap{""})
	t = template.Must(t.ParseFiles(tmp_list...))

	f, _ := os.Create("index.html")
	defer f.Close()

	err := t.Execute(f)
}
