package authentication

import (
	"Reddit-Clone/src/api"
	"Reddit-Clone/src/share/config"
)

type AuthenticationModule struct {
	cfg        *config.Config
	service    *AuthenticationService
	controller *AuthenticationController
}

func initRoutes(c *AuthenticationController) {
	router := api.GetApiRoute()
	authGroup := router.Group("/auth")
	authGroup.GET("/login", c.Login)
}
func NewAuthentionModule() *AuthenticationModule {
	cfg := config.GetConfig()
	service := NewAuthenticationService(cfg)
	controller := NewAuthenticationController(cfg, service)
	initRoutes(controller)
	return &AuthenticationModule{
		cfg,
		service,
		controller,
	}
}
