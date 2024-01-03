package main

import (
	"github.com/reddit-clone/src/api"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

func main() {
	cfg := config.GetConfig()
	lg := custome_logger.NewLogger(cfg)
	redisCancel, err := cache.InitRedis(cfg, lg)
	defer redisCancel()
	if err != nil {
		lg.Error(custome_logger.Redis, custome_logger.Connect, err.Error(), nil)
	}
	err = postgres.InitPostgres(cfg, lg)
	api.InitServer(cfg)
	lg.Info(custome_logger.General, custome_logger.Startup, "Application started", nil)
}
