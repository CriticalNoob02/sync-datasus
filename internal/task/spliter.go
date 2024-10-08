package task

import "github.com/CriticalNoob02/sync-datasus/internal/config"

// Funcao responsavel por realizar a separacao da lista em batchs;
func Spliter(list []string) [][]string {
	listLen := len(list)
	batch := config.GetBatchLimit()
	var splitedList = [][]string{}

	count := 0
	for count < listLen {
		end := count + batch
		if end > listLen {
			end = listLen
		}

		batchList := list[count:end]
		splitedList = append(splitedList, batchList)

		count += batch
	}
	return splitedList
}
