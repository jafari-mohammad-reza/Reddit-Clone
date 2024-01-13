package authentication

import (
	authDto "github.com/reddit-clone/src/domains/user-domain/authentication/dto"
	"github.com/reddit-clone/src/share/config"
	"github.com/reddit-clone/src/share/services"
)
type AuthenticationService struct {
	cfg *config.Config
}

func NewAuthenticationService(cfg *config.Config) *AuthenticationService {
	return &AuthenticationService{
		cfg,
	}
}

func (s *AuthenticationService) Login(dto authDto.LoginDto)  (*services.JwtAuthToken, error) {
	return nil , nil
}