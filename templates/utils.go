package templates

import (
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func mapToRows(jsonMap map[string]any) (rows []TableRow) {
	for field, value := range jsonMap {
		var subTable Table
		var valueType string
		var detail string

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
			subTable.Entity = valueType
			subTable.Rows = mapToRows(value.(map[string]any))
			row.SubTable = &subTable
			row.Type = valueType
			row.Detail = ""
			row.NotPrimitive = true
		}
		rows = append(rows, row)
	}
	return
}
