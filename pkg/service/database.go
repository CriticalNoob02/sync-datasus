package service

import (
	"context"
	"sync"
	"time"

	"github.com/CriticalNoob02/sync-datasus/internal/config"
	"github.com/CriticalNoob02/sync-datasus/pkg/util"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	poolInstance *pgxpool.Pool
	once         sync.Once
)

func getConfig() *pgxpool.Config {
	const (
		defaultMinConns          = int32(0)
		defaultMaxConnLifetime   = time.Hour * 2
		defaultMaxConnIdleTime   = time.Hour
		defaultHealthCheckPeriod = time.Minute
	)

	dbConfig, err := pgxpool.ParseConfig(config.GetDatabaseUrl())
	if err != nil {
		util.Logger.Fatal("Opss", "err", err.Error())
	}

	dbConfig.MaxConns = int32(config.GetNumWorks())
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod

	return dbConfig
}

func GetDbPool() *pgxpool.Pool {
	once.Do(func() {
		dbConfig := getConfig()

		var err error
		poolInstance, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
		if err != nil {
			util.Logger.Fatal("Opss", "err", err.Error())
		}

	})
	return poolInstance
}
