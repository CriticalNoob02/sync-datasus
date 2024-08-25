package util

import (
	"os"

	"github.com/charmbracelet/log"
)

func MakeDir(path string) bool {
	err := os.MkdirAll(path, 0755)
	exist := os.IsExist(err)
	if exist {
		log.Info("directory already exists!")
	}
	return true
}
