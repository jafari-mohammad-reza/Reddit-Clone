package src

import (
	//"Reddit-Clone/src/share/database/cache"
	//"Reddit-Clone/src/share/database/db/postgres"
	//"Reddit-Clone/src/share/database/db/postgres/migrations"
	//"github.com/reddit-clone/src/api"
	"Reddit-Clone/src/api"
	//"github.com/reddit-clone/src/share/pkg/queue"
)

//type AppModule struct {
//	UserDomain *userdomain.UserDomain
//}

func InitApp() {
	//initDb(cfg, lg)
	//queue.InitRabbitMq(cfg, lg)
	//defer queue.CloseRabbitConnection(lg)
	//defer queue.CloseRabbitChanel(lg)
	api.InitServer()
	//userdomain.NewUserDomain()
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
