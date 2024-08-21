package util

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var Logger = getLogger()

func getLogger() *log.Logger {
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
	})

	return logger
}
