package user_domain

import (
	"Reddit-Clone/src/domains/user-domain/authentication"
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
