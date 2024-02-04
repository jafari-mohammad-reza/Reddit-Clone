package user_domain

import (
	"github.com/gin-gonic/gin"
	"github.com/reddit-clone/src/domains/user-domain/authentication"
)

type UserDomain struct {
	authModule *authentication.AuthenticationModule
}

func NewUserDomain(r *gin.RouterGroup) *UserDomain {
	authModule := authentication.NewAuthenticationModule(r)

	return &UserDomain{
		authModule,
	}
}
