package authentication

import "github.com/reddit-clone/src/share/config"
type AuthenticationService struct {
	cfg *config.Config
}

func NewAuthenticationService(cfg *config.Config) *AuthenticationService {
	return &AuthenticationService{
		cfg,
	}
}