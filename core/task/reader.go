package task

import (
	"sync-datasus/core/config"
	"sync-datasus/core/service"
	"sync-datasus/core/util"

	"github.com/charmbracelet/log"
)

// Funcao responsavel por efetuar a leitura e filtragem da lista
// de arquivos na conexao FTP;
func Reader() []string {
	conn := service.FtpLogin("anonymous", "anonymous", config.GetFtpUrl())

	log.Debug("Login ftp realizado com sucesso...")

	filesName := service.FtpList(conn, config.GetModuleRemoteDir())
	log.Debug("Listagem agrupada com sucesso...")

	var filteredList []string
	for _, file := range filesName {
		mockDate := util.DataFilterStruct{
			MonthPosition: [2]int{4, 6},
			YearPosition:  [2]int{6, 8},
		}

		if util.FilterString(config.GetModuleType(), file, "start") {
			validation, err := util.FilterDate(config.GetModuleLimitDate(), file, mockDate)
			util.Check(err)
			if validation {
				filteredList = append(filteredList, file)
			}
		}
	}
	log.Debug("Listagem filtrada com sucesso...")
	return filteredList
}
