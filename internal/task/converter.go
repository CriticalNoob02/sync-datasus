package task

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"
)

// Funcao responsavel por converter um arquivo DBC para o formato DBF;
func Converter(filename string) (string, error) {
	filenameDbf := strings.Replace(filename, ".dbc", ".dbf", 1)
	pathDbc := fmt.Sprintf("%s/%s", config.GetDownloadPath(), filename)
	pathDbf := fmt.Sprintf("%s/%s", config.GetExtractPath(), filenameDbf)

	command := exec.Command("pkg/service/blast/blast-dbf", pathDbc, pathDbf)
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
