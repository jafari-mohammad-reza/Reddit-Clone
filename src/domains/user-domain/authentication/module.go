package authentication

import (
	"github.com/gin-gonic/gin"
	authentication "github.com/reddit-clone/src/domains/user-domain/authentication/services"
	"github.com/reddit-clone/src/share/config"
)

type AuthenticationModule struct {
	cfg        *config.Config
	service    *authentication.AuthenticationService
	controller *AuthenticationController
}

func initRoutes(r *gin.RouterGroup, c *AuthenticationController) {
	authGroup := r.Group("/auth")
	authGroup.POST("/login", c.Login)
	authGroup.POST("/register", c.Register)
	githubGroup := authGroup.Group("/github")
	githubGroup.POST("/login", c.GithubLogin)
	githubGroup.POST("/verify", c.GithubVerify)
	googleGroup := authGroup.Group("/google")
	googleGroup.POST("/login", c.GoogleLogin)
	googleGroup.POST("/verify", c.GoogleVerify)
}
func NewAuthenticationModule(r *gin.RouterGroup) *AuthenticationModule {
	cfg := config.GetConfig()
	service := authentication.NewAuthenticationService(cfg)
	controller := NewAuthenticationController(cfg, service)
	initRoutes(r, controller)
	return &AuthenticationModule{
		cfg,
		service,
		controller,
	}
}
