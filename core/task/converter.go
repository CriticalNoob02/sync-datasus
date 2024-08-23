package task

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync-datasus/core/config"
	"sync-datasus/core/util"
)

// Funcao responsavel por converter um arquivo DBC para o formato DBF;
func Converter(filename string) (string, error) {
	filenameDbf := strings.Replace(filename, ".dbc", ".dbf", 1)
	pathDbc := fmt.Sprintf("%s/%s", config.GetDownloadPath(), filename)
	pathDbf := fmt.Sprintf("%s/%s", config.GetExtractPath(), filenameDbf)

	command := exec.Command("core/service/blast/blast-dbf", pathDbc, pathDbf)
	err := command.Run()
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return "", err
	}

	err = os.Remove(pathDbc)
	if err != nil {
		util.Logger.Error("Opss", "err", err.Error())
		return "", err
	}

	return filenameDbf, nil
}
