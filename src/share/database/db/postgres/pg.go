package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

var pgClient *sql.DB

func InitPostgres(cfg *config.Config, lg custome_logger.Logger) error {
	var err error
	port, _ := strconv.Atoi(cfg.Postgres.Port)
	url := fmt.Sprintf(
		"sslmode=%s host=%s port=%d user=%s password=%s dbname=%s",
		cfg.Postgres.SSLMode,
		cfg.Postgres.Host,
		port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DbName,
	)
	pgClient, err = sql.Open("postgres", url)
	if err != nil {
		return err
	}
	// Set maximum number of open connections to the database.
	pgClient.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	// Set maximum number of idle connections in the pool.
	pgClient.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	// Set maximum amount of time a connection may be reused.
	pgClient.SetConnMaxLifetime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)
	// Set the maximum amount of time a connection may be idle.
	pgClient.SetConnMaxIdleTime(time.Duration(cfg.Postgres.ConnMaxLifetime) * time.Minute)
	// Set a timeout for establishing new connections.
	pgClient.SetConnMaxIdleTime(time.Duration(cfg.Postgres.MaxIdleConnsTime) * time.Minute)
	// Ping the database to verify connection is established.
	err = pgClient.Ping()
	if err != nil {
		return err
	}
	lg.Info(custome_logger.Postgres, custome_logger.Connect, "Postgres connected", nil)
	return nil
}

func GetPostgres() *sql.DB {
	return pgClient
}

func ClosePostgres(lg custome_logger.Logger) error {
	lg.Info(custome_logger.Postgres, custome_logger.Close, "", nil)
	return pgClient.Close()
}
