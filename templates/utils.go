package templates

import (
	"strconv"

	"github.com/google/uuid"
)

func mapToRows(jsonMap map[string]any) (rows []TableRow) {
	for field, value := range jsonMap {
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
		rows = append(rows, TableRow{Field: field, Type: valueType, Detail: detail})
	}
	return
}
