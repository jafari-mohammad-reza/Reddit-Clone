package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	authentication "github.com/reddit-clone/src/domains/user-domain/authentication/services"
	"github.com/reddit-clone/src/share/config"
)
type AuthenticationController struct {
	cfg *config.Config
	service *authentication.AuthenticationService
}

func NewAuthenticationController(cfg *config.Config , service *authentication.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		cfg,
		service,
	}
}

func(c *AuthenticationController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})
}

func(c *AuthenticationController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})
}