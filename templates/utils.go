package templates

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func mapToRows(jsonMap map[string]any) (rows []TableRow, fields []string) {
	for field, value := range jsonMap {
		var subTable Table
		var valueType string
		var detail string

		fields = append(fields, field)

		floatVal, ok := value.(float64)
		if ok {
			valueType = "number"
			detail = "(float)"
		}
		strVal := strconv.FormatFloat(floatVal, 'f', -1, 64)
		_, err := strconv.Atoi(strVal)
		if err == nil {
			valueType = "number"
			detail = "(int)"
		}
		str, ok := value.(string)
		if ok {
			detail = ""
			valueType = "string"

			_, err := uuid.Parse(str)
			if err == nil {
				detail = "(uuid)"
			}
		}
		_, ok = value.(bool)
		if ok {
			detail = ""
			valueType = "boolean"
		}

		row := TableRow{Field: field, Type: valueType, Detail: detail}

		_, ok = value.(map[string]any)
		if ok {
			valueType = ""
			fieldSliced := strings.Split(field, "_")

			for _, frag := range fieldSliced {
				// valueType += strings.Title(frag)
				valueType += cases.Title(language.Und).String(frag)
			}
			r, f := mapToRows(value.(map[string]any))
			subTable.Rows = r
			subTable.Entity = valueType
			row.SubTable = &subTable
			row.Type = valueType
			row.Detail = ""
			row.NotPrimitive = true
			fields = append(fields, f...)
		}

		aList, ok := value.([]any)
		if ok {
			first := aList[0]

			floatVal, ok := first.(float64)
			if ok {
				valueType = "[]number"
				row.Detail = "(float)"
			}
			inStrVal := strconv.FormatFloat(floatVal, 'f', -1, 64)
			_, err := strconv.Atoi(inStrVal)
			if err == nil {
				valueType = "[]number"
				row.Detail = "(int)"
			}
			str, ok := first.(string)
			if ok {
				row.Detail = ""
				valueType = "[]string"

				_, err := uuid.Parse(str)
				if err == nil {
					detail += "(uuid)"
				}
			}

			_, ok = first.(map[string]any)
			if ok {
				valueType = ""
				fieldSliced := strings.Split(field, "_")

				for _, frag := range fieldSliced {
					valueType += strings.Title(frag)
				}
				r, f := mapToRows(first.(map[string]any))
				subTable.Rows = r
				subTable.Entity = valueType
				row.SubTable = &subTable
				row.Detail = "List"
				row.NotPrimitive = true
				fields = append(fields, f...)
			}
			row.Type = valueType
		}

		rows = append(rows, row)
	}
	return
}

func highlightFields(fields []string, commentary string) (cws []CommentWord) {
	words := strings.Split(commentary, " ")

	for _, word := range words {
		var w *map[int]string
		hl, field := doesWordContainsAField(word, fields)

		if hl {
			wSlc := strings.Split(word, field)

			if len(wSlc) > 0 {
				if hasAlpha(wSlc[0]) {
					hl = false
				} else {
					switch len(wSlc) {
					case 1:
						// w[0] = ""
						// w[1] = field
						// w[2] = wSlc[0]
						w = &map[int]string{
							0: "",
							1: field,
							2: wSlc[0],
						}
					case 2:
						if hasAlpha(wSlc[1]) {
							hl = false
						} else {
							// w[0] = wSlc[0]
							// w[1] = field
							// w[2] = wSlc[1]
							w = &map[int]string{
								0: wSlc[0],
								1: field,
								2: wSlc[1],
							}
						}
					}
				}
			} else {
				w = &map[int]string{
					0: word,
					1: "",
					2: "",
				}
			}
		}

		if !hl {
			w = &map[int]string{
				0: word,
				1: "",
				2: "",
			}
		}

		cws = append(cws, CommentWord{
			Word:      *w,
			Highlight: hl,
		})
	}
	// for _, field := range fields {
	// 	var com string
	// 	comSliced := strings.Split(commentary, field)
	// 	if len(comSliced) > 1 {
	// 		for i, frag := range comSliced {
	// 			com += frag
	// 			if i < (len(comSliced) - 1) {
	// 				com += `<span class="text-sky-600">` + field + `</span>`
	// 			}
	// 		}
	// 		commentary = com
	// 	}
	// 	comment = commentary
	// }

	// comment = commentary
	return
}

func hasAlpha(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func doesWordContainsAField(word string, fields []string) (bool, string) {
	for _, field := range fields {
		if strings.Contains(word, field) {
			return true, field
		}
	}
	return false, ""
}
