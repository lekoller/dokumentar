package templates

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
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
		row := TableRow{Field: field, Type: valueType, Detail: detail}

		_, ok = value.(map[string]any)
		if ok {
			valueType = ""
			fieldSliced := strings.Split(field, "_")

			for _, frag := range fieldSliced {
				valueType += strings.Title(frag)
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
		// log.Println(value, ok)
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

func highlightFields(fields []string, commentary string) (comment string) {
	for _, field := range fields {
		var com string
		comSliced := strings.Split(commentary, field)

		if len(comSliced) > 1 {
			for i, frag := range comSliced {
				com += frag
				if i < (len(comSliced) - 1) {
					com += `<span class="text-sky-600">` + field + `</span>`
				}
			}
			commentary = com
		}
		comment = commentary
	}
	return
}
