package modules

import (
	"context"
	"sync-datasus/core/config"
	"sync-datasus/core/service"
	"sync-datasus/core/task"
	"sync-datasus/core/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Worker(files []string, pool *pgxpool.Pool) {
	conn := service.FtpLogin("anonymous", "anonymous", config.GetFtpUrl())

	// Permissao para o blast: chmod +x core/service/blast/blast-dbf

	for _, file := range files {
		dbConn, err := pool.Acquire(context.Background())
		util.Check(err)

		var localFile string = (config.GetDownloadPath() + "/" + file)
		var remoteFile string = (config.GetModuleRemoteDir() + "/" + file)

		util.Logger.Info("Worker", "Iniciando", "Download e conversao", "Arquivo", file)
		service.FtpDownloadFile(conn, localFile, remoteFile)
		fileConverted, err := task.Converter(file)
		util.Check(err)

		var fileConvertedPath string = (config.GetExtractPath() + "/" + fileConverted)

		util.Logger.Info("Worker", "Iniciando", "Insert no banco")
		columns, err := task.Mapper(dbConn, "tb_fat_importacoes_raas", "public")

		util.Check(err)
		err = task.Writer(fileConvertedPath, "tb_fat_importacoes_raas", "public", columns, dbConn)
		util.Check(err)
		dbConn.Release()
	}
}
