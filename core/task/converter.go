package task

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync-datasus/core/config"
)

func Converter(filename string) (bool, error) {
	filenameDbf := strings.Replace(filename, ".dbc", ".dbf", 1)
	pathDbc := fmt.Sprintf("%s/%s", config.GetDownloadPath(), filename)
	pathDbf := fmt.Sprintf("%s/%s", config.GetExtractPath(), filenameDbf)

	command := exec.Command("core/service/blast/blast-dbf", pathDbc, pathDbf)
	err := command.Run()
	if err != nil {
		return false, errors.New("<fg=magenta> Converter</> - " + err.Error())
	}

	err = os.Remove(pathDbc)
	if err != nil {
		return false, errors.New("<fg=magenta> Converter</> - " + err.Error())
	}

	return true, nil
}
