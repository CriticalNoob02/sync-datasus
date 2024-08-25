package database

import (
	"fmt"
	"strings"
)

func GetInsertQuery(table string, schema string, columns []string, values [][]string) string {
	columnsString := fmt.Sprintf(`"%s"`, strings.Join(columns, `", "`))

	query := fmt.Sprintf("INSERT INTO %s.%s (%s) VALUES ", schema, table, columnsString)

	var valuesParts []string
	for _, row := range values {
		for i, val := range row {
			row[i] = fmt.Sprintf("'%s'", val)
		}
		valuesParts = append(valuesParts, fmt.Sprintf("(%s)", strings.Join(row, ", ")))
	}

	query += strings.Join(valuesParts, ", ")
	return query
}
