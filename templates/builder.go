package templates

import (
	"encoding/json"
	"log"
	"os"
	"text/template"

	"github.com/gogap/config"
	"github.com/gogap/go-pandoc/pandoc"
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
	Comment      string
	Rows         []TableRow
	SubTables    []Table
	HasSubTables bool
}

type TableRow struct {
	Field        string
	Type         string
	Detail       string
	SubTable     *Table
	NotPrimitive bool
}

func BuildTemplate(data BuilderDTO) {
	document := Document{
		ProjectName:   data.ProjectName,
		ContainerName: data.ContainerName,
		ModuleName:    data.ModuleName,
	}

	for i, item := range data.List {
		var tables []Table
		var jsonMap map[string]any
		json.Unmarshal([]byte(data.List[i].JsonEntry), &jsonMap)

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

	// wkhtmltopdf.

	// pdfGen, err := wkhtmltopdf.NewPDFGenerator()
	// if err != nil {
	// 	log.Println("creating generator:", err.Error())
	// }
	// pdfGen.AddPage(wkhtmltopdf.NewPageReader(f))
	// err = pdfGen.Create()
	// if err != nil {
	// 	log.Println("creating document:", err.Error())
	// }
	// pdfGen.WriteFile("./docs.pdf")

	// doc := wkhtmltopdf.NewDocument()
	// pg := wkhtmltopdf.NewPage("index.html")
	// doc.AddPages(pg)

	// err = doc.WriteToFile("docs.pdf")
	// if err != nil {
	// 	log.Println("creating document:", err.Error())
	// }
	// m := pdf.NewMaroto(consts.Portrait, consts.A4)
	// m.SetPageMargins(20, 10, 20)

	// err = m.OutputFileAndClose("docs_" + document.ProjectName + ".pdf")
	// if err != nil {
	// 	log.Println("creating document:", err.Error())
	// }
	conf := config.NewConfig()
	log.Println(conf.Configuration)

	pdoc, err := pandoc.New(conf)
	if err != nil {
		log.Println(err.Error())
	}
	convData, err := pdoc.Convert(
		pandoc.FetcherOptions{
			Name:   "data",
			Params: []byte((`{"url": "./index.html"`)),
		}, pandoc.ConvertOptions{
			From:      "html",
			To:        "pdf",
			DataDir:   "./",
			PDFEngine: "weasyprint",
		},
	)

	log.Println(convData)
	f2, err := os.Create("docs.pdf")
	_, err = f2.Write(convData)
	if err != nil {
		log.Println(err.Error())
	}
}
