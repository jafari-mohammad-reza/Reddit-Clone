package main

import (
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)

func main() {
	cfg := config.GetConfig()
	lg := custome_logger.NewLogger(cfg)
	redisCancel, err := cache.InitRedis(cfg)
	defer redisCancel()
	if err != nil {
		lg.Error(custome_logger.Redis, custome_logger.Connect, err.Error(), nil)
	}
	mongoCancel, err := db.InitMongo(cfg)
	if err != nil {
		lg.Error(custome_logger.Mongo, custome_logger.Connect, err.Error(), nil)
	}
	defer mongoCancel()
	lg.Info(custome_logger.General, custome_logger.Startup, "Application started", nil)

}
