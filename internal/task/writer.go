package task

import (
	"context"
	"os"
	"strings"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/internal/database"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/LindsayBradford/go-dbf/godbf"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Funcao responsavel por ler os dados do arquivo DBF e escrever no banco;
func Writer(filepath string, columns []string, conn *pgxpool.Conn) error {
	var count int = 0
	var recordRow [][]string
	var upperCaseCol []string

	for _, str := range columns {
		upperCaseCol = append(upperCaseCol, strings.ToUpper(str))
	}

	dbfTable, err := godbf.NewFromFile(filepath, "utf-8")
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return err
	}
	recordCount := dbfTable.NumberOfRecords()

	cursor, err := conn.Begin(context.Background())
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return err
	}
	defer func() {
		if err != nil {
			cursor.Rollback(context.Background())
		}
	}()

	for i := 0; i < recordCount; i++ {
		var recordCol []string
		count += 1
		cols := dbfTable.FieldNames()

		for _, colName := range cols {
			if contains(upperCaseCol, colName) {
				e := indexOf(colName, cols)
				recordCol = append(recordCol, dbfTable.FieldValue(i, e))
			}
		}
		recordRow = append(recordRow, recordCol)
		if count == config.GetQueryLimit() || i >= recordCount-1 {
			query := database.GetInsertQuery(config.GetModuleTableName(), config.GetModuleSchemaName(), columns, recordRow)
			util.Logger.Debug(query)

			_, err = cursor.Exec(context.Background(), query)
			if err != nil {
				util.Logger.Error("Opss", "err", err.Error())
				return err
			}
			count = 0
			recordRow = nil
		}
	}

	err = cursor.Commit(context.Background())
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return err
	}

	os.Remove(filepath)
	return nil
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func indexOf(str string, arr []string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}
