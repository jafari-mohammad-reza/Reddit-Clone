package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/share/config"
)
type AuthenticationController struct {
	cfg *config.Config
	service *AuthenticationService
}

func NewAuthenticationController(cfg *config.Config , service *AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		cfg,
		service,
	}
}

func(c *AuthenticationController) Login(ctx *gin.Context) {
	ctx.JSON(200 , gin.H{"msg": "login route"})
}