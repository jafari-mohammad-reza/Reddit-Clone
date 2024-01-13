package user_domain

import (
	authentication "github.com/reddit-clone/src/domains/user-domain/authentication"
)

type UserDomain struct {
	authModule *authentication.AuthenticationModule
}

func NewUserDomain() *UserDomain {
	authModule := authentication.NewAuthentionModule()
	return &UserDomain{
		authModule,
	}
}
