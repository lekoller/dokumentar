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
	Entity       string
	Connection   string
	Method       string
	Endpoint     string
	Comment      []CommentWord
	Rows         []TableRow
	SubTables    []Table
	HasSubTables bool
}

type CommentWord struct {
	Word      map[int]string
	Highlight bool
}

type TableRow struct {
	Field        string
	Type         string
	Detail       string
	SubTable     *Table
	NotPrimitive bool
}

func BuildTemplate(dto BuilderDTO) {
	document := Document{
		ProjectName:   dto.ProjectName,
		ContainerName: dto.ContainerName,
		ModuleName:    dto.ModuleName,
	}

	for i, item := range dto.List {
		var tables []Table
		var jsonMap map[string]any
		json.Unmarshal([]byte(dto.List[i].JsonEntry), &jsonMap)

		rows, fields := mapToRows(jsonMap)

		for _, row := range rows {
			if row.SubTable != nil {
				var tbs []Table

				for _, ro := range row.SubTable.Rows {
					if ro.SubTable != nil {
						var ts []Table

						for _, r := range ro.SubTable.Rows {
							if r.SubTable != nil {
								ts = append(ts, *r.SubTable)
							}
						}
						ro.SubTable.SubTables = ts
						ro.SubTable.HasSubTables = len(ts) > 0

						tbs = append(tbs, *ro.SubTable)
					}
				}
				row.SubTable.SubTables = tbs
				row.SubTable.HasSubTables = len(tbs) > 0

				tables = append(tables, *row.SubTable)
			}
		}
		document.Tables = append(document.Tables, Table{
			Entity:       item.Entity,
			Connection:   item.ConnType,
			Method:       item.Method,
			Endpoint:     item.Endpoint,
			Comment:      highlightFields(fields, item.CommentEntry),
			Rows:         rows,
			SubTables:    tables,
			HasSubTables: len(tables) > 0,
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
