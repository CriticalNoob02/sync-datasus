package main

import (
	"sync"

	"github.com/CriticalNoob02/sync-datasus/internal/modules"
)

func main() {
	var wg sync.WaitGroup

	modules.Mapper("RAAS")
	newFiles, pool := modules.Manager()

	for _, listFile := range newFiles {
		wg.Add(1)
		go func() {
			modules.Worker(listFile, pool)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
