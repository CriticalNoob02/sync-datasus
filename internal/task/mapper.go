package task

import (
	"context"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/internal/database"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Funcao responsavel por retornar as colunas da tabela;
func Mapper(conn *pgxpool.Conn) ([]string, error) {
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

	colQuery := database.GetCollunsQuery(config.GetModuleTableName(), config.GetModuleSchemaName())

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
