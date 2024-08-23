package database

import "fmt"

func GetCollunsQuery(tableName string, schema string) string {
	return fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_schema = '%s' and table_name = '%s';", schema, tableName)
}
