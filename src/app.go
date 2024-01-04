package src

import (
	"github.com/reddit-clone/src/api"
	user_domain "github.com/reddit-clone/src/domains/user-domain"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/database/cache"
	"github.com/reddit-clone/src/share/database/db/postgres"
	"github.com/reddit-clone/src/share/pkg/custome_logger"
)
type AppModule struct {
	UserDoamin *user_domain.UserDomain
}
func InitApp(cfg *config.Config , lg custome_logger.Logger){
	redisCancel, err := cache.InitRedis(cfg, lg)
	defer redisCancel()
	if err != nil {
		lg.Error(custome_logger.Redis, custome_logger.Connect, err.Error(), nil)
	}
	err = postgres.InitPostgres(cfg, lg)
	if err != nil {
		lg.Error(custome_logger.Postgres, custome_logger.Connect, err.Error(), nil)
	}
	api.InitServer(cfg)
	user_domain.NewUserDomain()
}