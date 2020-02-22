package driver

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/log/log15adapter"
	"github.com/jackc/pgx/v4/pgxpool"
	log "gopkg.in/inconshreveable/log15.v2"
)

var db *pgxpool.Pool

// ConnectDB for connection db
func ConnectDB() *pgxpool.Pool {
	logger := log15adapter.NewLogger(log.New("module", "pgx"))
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("APP_DB_URL"))

	if err != nil {
		log.Crit("Unable to parse APP_DB_URL", "error", err)
		os.Exit(1)
	}

	poolConfig.ConnConfig.Logger = logger

	db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Crit("Unable to create connection pool", "error", err)
		os.Exit(1)
	}

	return db
}
