package src

import (
	//"Reddit-Clone/src/share/database/cache"
	//"Reddit-Clone/src/share/database/db/postgres"
	//"Reddit-Clone/src/share/database/db/postgres/migrations"
	//"github.com/reddit-clone/src/api"
	"Reddit-Clone/src/api"
	userdomain "Reddit-Clone/src/domains/user-domain"
	"Reddit-Clone/src/share/config"
	//"github.com/reddit-clone/src/share/database/cache"
	//"github.com/reddit-clone/src/share/database/db/postgres"
	//"github.com/reddit-clone/src/share/database/db/postgres/migrations"
	//"github.com/reddit-clone/src/share/pkg/custome_logger"
	"Reddit-Clone/src/share/pkg/custome_logger"
	//"github.com/reddit-clone/src/share/pkg/queue"
)

type AppModule struct {
	UserDomain *userdomain.UserDomain
}

func InitApp(cfg *config.Config, lg custome_logger.Logger) {
	//initDb(cfg, lg)
	//queue.InitRabbitMq(cfg, lg)
	//defer queue.CloseRabbitConnection(lg)
	//defer queue.CloseRabbitChanel(lg)
	api.InitServer(cfg)
	userdomain.NewUserDomain()
}

//func initDb(cfg *config.Config, lg custome_logger.Logger) {
//	redisCancel, err := cache.InitRedis(cfg, lg)
//	defer redisCancel()
//	if err != nil {
//		lg.Error(custome_logger.Redis, custome_logger.Connect, err.Error(), nil)
//		panic(err)
//	}
//	err = postgres.InitPostgres(cfg, lg)
//	if err != nil {
//		lg.Error(custome_logger.Postgres, custome_logger.Connect, err.Error(), nil)
//		panic(err)
//	}
//	err = migrations.SeedData()
//	if err != nil {
//		lg.Error(custome_logger.Postgres, custome_logger.Seed, err.Error(), nil)
//		panic(err)
//	}
//}
