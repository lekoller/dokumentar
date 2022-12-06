package templates

import (
	"encoding/json"
	"os"
	"text/template"
)

type BuilderDTO struct {
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

type Document struct {
	ProjectName   string
	ContainerName string
	ModuleName    string
	Tables        []Table
}

type Table struct {
	Entity     string
	Connection string
	Method     string
	Endpoint   string
	Comment    string
	Rows       []TableRow
}

type TableRow struct {
	Field  string
	Type   string
	Detail string
}

func BuildTemplate(data BuilderDTO) {
	document := Document{
		ProjectName:   data.ProjectName,
		ContainerName: data.ContainerName,
		ModuleName:    data.ModuleName,
	}

	for i, item := range data.List {
		var jsonMap map[string]any
		json.Unmarshal([]byte(data.List[i].JsonEntry), &jsonMap)

		rows := mapToRows(jsonMap)

		if item.Method == "" {
			item.Method = "--"
		}
		document.Tables = append(document.Tables, Table{
			Entity:     item.Entity,
			Connection: item.ConnType,
			Method:     item.Method,
			Endpoint:   item.Endpoint,
			Comment:    item.CommentEntry,
			Rows:       rows,
		})
	}

	tmp_list := []string{
		"templates/head.html",
		"templates/body.html",
		"templates/foot.html",
	}
	t := template.New("body.html")
	// t.Funcs(template.FuncMap{""})
	t = template.Must(t.ParseFiles(tmp_list...))

	f, _ := os.Create("index.html")
	defer f.Close()

	err := t.Execute(f, document)
	if err != nil {
		panic(err)
	}
}
