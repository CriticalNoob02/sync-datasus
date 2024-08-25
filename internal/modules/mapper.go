package modules

import "os"

func Mapper(types string) {
	switch types {
	case "RAAS":
		os.Setenv("MODULE_TYPE", "PS")
		os.Setenv("MODULE_REMOTE_DIR", "dissemin/publicos/SIASUS/200801_/Dados")
		os.Setenv("MODULE_TABLE_NAME", "tb_fat_importacoes_raas")
		os.Setenv("MODULE_TABLE_SCHEMA", "public")
		os.Setenv("QUERY_LIMIT", "1000")
	}
}
