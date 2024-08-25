package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetBatchLimit() int {
	limit, _ := strconv.Atoi(os.Getenv("BATCH_LIMIT"))
	return limit
}

func GetQueryLimit() int {
	limit, _ := strconv.Atoi(os.Getenv("QUERY_LIMIT"))
	return limit
}

func GetNumWorks() int {
	works, _ := strconv.Atoi(os.Getenv("NUM_WORKS"))
	return works
}

func GetFtpUrl() string {
	return os.Getenv("FTP_SERVER_URL")
}

func GetModuleType() string {
	return os.Getenv("MODULE_TYPE")
}

func GetModuleRemoteDir() string {
	return os.Getenv("MODULE_REMOTE_DIR")
}

func GetModuleLimitDate() string {
	return os.Getenv("MODULE_LIMIT_DATE")
}

func GetModuleTableName() string {
	return os.Getenv("MODULE_TABLE_NAME")
}

func GetModuleSchemaName() string {
	return os.Getenv("MODULE_TABLE_SCHEMA")
}

func GetDownloadPath() string {
	return "storage/download"
}

func GetExtractPath() string {
	return "storage/extract"
}

func GetDatabaseUrl() string {
	user := os.Getenv("DB_USERNAME")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, name)
	return url
}
