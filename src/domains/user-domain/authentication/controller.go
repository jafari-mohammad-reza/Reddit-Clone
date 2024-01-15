package authentication

import (
	"net/http"

	"github.com/gin-gonic/gin"
	api "github.com/reddit-clone/src/api/helper"
	authDto "github.com/reddit-clone/src/domains/user-domain/authentication/dto"
	authentication "github.com/reddit-clone/src/domains/user-domain/authentication/services"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/services"
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
	var body  authDto.LoginDto
	if err:=ctx.BindJSON(&body);err!=nil{
		status := http.StatusBadRequest 
		ctx.JSON(status  , api.GenerateErrorResponse(err, "/auth/login", &status))
		return 
	}
	token , err := c.service.Login(body)
	if err != nil {
		status := http.StatusBadRequest
		ctx.JSON(status  , api.GenerateErrorResponse(err, "/auth/login", &status))
		return 
	}
	
	ctx.JSON(http.StatusOK, api.GenerateSuccessResponse[*services.JwtAuthToken](token, nil, nil))
}

func(c *AuthenticationController) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})
}

func(c *AuthenticationController) GithubLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})

}

func(c *AuthenticationController) GithubVerify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})
}


func(c *AuthenticationController) GoogleLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})

}

func(c *AuthenticationController) GoogleVerify(ctx *gin.Context) {
	ctx.JSON(http.StatusOK , gin.H{})
}

