package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/reddit-clone/src/share/config"
	"time"
)

type JwtService struct {
	cfg *config.Config
}

func NewJwtService(cfg *config.Config) *JwtService {
	return &JwtService{
		cfg,
	}
}

type JwtToken struct {
	Token string
	Exp   int64
}

type JwtAuthToken struct {
	AccessToken     string
	AccessTokenExp  int64
	RefreshToken    string
	RefreshTokenExp int64
}
type JwtAuthTokenPayload struct {
	Id    int
	Roles []string
}

func (s *JwtService) GenerateAuthToken(payload JwtAuthTokenPayload) (*JwtAuthToken, error) {
	mp := jwt.MapClaims{}
	mp["Id"] = payload.Id
	mp["Roles"] = payload.Roles
	// access token expiration time
	accExp := time.Now().Add(time.Hour).Unix()
	accToken, err := s.GenerateToken(mp, accExp)
	if err != nil {
		return nil, err
	}
	// refresh token expiration time
	refExp := time.Now().Add(time.Hour).Unix()
	refToken, err := s.GenerateToken(mp, refExp)
	if err != nil {
		return nil, err
	}
	authToken := JwtAuthToken{
		AccessToken:     accToken.Token,
		AccessTokenExp:  accToken.Exp,
		RefreshToken:    refToken.Token,
		RefreshTokenExp: refToken.Exp,
	}
	return &authToken, nil
}
func (s *JwtService) GenerateToken(payload jwt.MapClaims, exp int64) (*JwtToken, error) {
	secret := s.cfg.JWT.Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}
	return &JwtToken{
		Token: tokenString,
		Exp:   exp,
	}, err
}
