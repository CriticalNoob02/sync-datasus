package modules

import (
	"sync-datasus/core/service"
	"sync-datasus/core/task"
	"sync-datasus/core/util"

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
