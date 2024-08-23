package task

import (
	"context"
	"sync-datasus/core/database"
	"sync-datasus/core/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Funcao responsavel por retornar as colunas da tabela;
func Mapper(conn *pgxpool.Conn, tableName string, schema string) ([]string, error) {
	cursor, err := conn.Begin(context.Background())
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return nil, err
	}
	defer func() {
		if err != nil {
			cursor.Rollback(context.Background())
		}
	}()

	colQuery := database.GetCollunsQuery(tableName, schema)

	rows, err := cursor.Query(context.Background(), colQuery)
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return nil, err
	}

	var colList []string
	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			util.Logger.Error("Opss", "err", err.Error())
			return nil, err
		}
		colList = append(colList, value)
	}
	return colList, nil
}
