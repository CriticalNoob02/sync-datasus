package main

import (
	"sync-datasus/core/config"
	"sync-datasus/core/service"
	"sync-datasus/core/task"
	"sync-datasus/core/util"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("infra/envs/.env")
	util.Check(err)
	util.Logger.Info("Variaveis carregadas...")

	files := task.Reader()
	util.Logger.Info("Reader finalizado...")

	newFiles := task.Spliter(files)
	util.Logger.Info("Spliter finalizado...")

	// Workers

	conn := service.FtpLogin("anonymous", "anonymous", config.GetFtpUrl())

	// Permissao para o blast: chmod +x core/service/blast/blast-dbf

	util.Logger.Info("Iniciando Download e Converter...")
	for _, file := range newFiles[0] {
		var localFile string = (config.GetDownloadPath() + "/" + file)
		var remoteFile string = (config.GetModuleRemoteDir() + "/" + file)

		service.FtpDownloadFile(conn, localFile, remoteFile)

		_, err = task.Converter(file)
		util.Check(err)
	}
	util.Logger.Info("Download e Converter finalizados com sucesso...")

}
