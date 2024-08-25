package modules

import (
	"github.com/CriticalNoob02/sync-datasus/internal/task"
	"github.com/CriticalNoob02/sync-datasus/pkg/service"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func Manager() ([][]string, *pgxpool.Pool) {
	util.Logger.Info("Manager", "Iniciando", "Carregamento de envs")
	err := godotenv.Load("infra/envs/.env")
	util.Check(err)

	util.Logger.Info("Manager", "Iniciando", "Reader")
	files := task.Reader()

	util.Logger.Info("Manager", "Iniciando", "Spliter")
	newFiles := task.Spliter(files)

	util.Logger.Info("Manager", "Iniciando", "PoolConnection")
	pool := service.GetDbPool()

	return newFiles, pool
}
