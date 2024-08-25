package modules

import (
	"context"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/internal/task"
	"github.com/CriticalNoob02/sync-datasus/pkg/service"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Worker(files []string, pool *pgxpool.Pool) {
	conn := service.FtpLogin("anonymous", "anonymous", config.GetFtpUrl())

	// TODO Permissao para o blast: chmod +x sync-datasus/pkg/service/blast/blast-dbf

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

		util.Logger.Debug("Worker", "Iniciando", "Insert no banco")
		columns, err := task.Mapper(dbConn)

		util.Check(err)
		err = task.Writer(fileConvertedPath, columns, dbConn)
		util.Check(err)

		dbConn.Release()
	}
}
