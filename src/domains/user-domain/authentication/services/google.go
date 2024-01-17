package authentication

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/reddit-clone/src/share/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)
type GoogleAuthService struct {
	cfg *config.Config
}

type GoogleUserInfo struct {
	ID            string 
	Email         string 
	VerifiedEmail bool   
	Name          string 
	GivenName     string 
	FamilyName    string 
	Link          string 
	Picture       string 
}

func NewGoogleAuthService(cfg *config.Config) *GoogleAuthService {
	return &GoogleAuthService{
		cfg: cfg,
	}
}

func (service *GoogleAuthService) GetGoogleOauthConfig() oauth2.Config {
	return oauth2.Config{
		RedirectURL:  service.cfg.Credentials.CallBackUrl,
		ClientID:     service.cfg.Credentials.ClientID,
		ClientSecret: service.cfg.Credentials.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

func (service *GoogleAuthService) GetLoginUrl() (*string, error) {
	URL, err := url.Parse("https://accounts.google.com/o/oauth2/auth")
	if err != nil {
		return nil, err
	}
	parameters := url.Values{}
	parameters.Add("client_id", os.Getenv("GOOGLE_CLIENT_ID"))
	scops := []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"}
	parameters.Add("scope", strings.Join(scops, " "))
	parameters.Add("redirect_uri", fmt.Sprintf("%s/%s", service.cfg.Server.ApiUrl, "api/v1/auth/google/callback"))
	parameters.Add("response_type", "code")
	parameters.Add("state", "")
	URL.RawQuery = parameters.Encode()
	loginUrl := URL.String()
	return &loginUrl, nil
}

func (service *GoogleAuthService) GetUserEmail(code string) (*string, error) {
	// Use the authorization code that is pushed to the redirect URL.
	googleOauthConfig := service.GetGoogleOauthConfig()
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + url.QueryEscape(token.AccessToken))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(response.Body)

	contents, err := io.ReadAll(response.Body)
	var userInfo GoogleUserInfo
	err = json.Unmarshal(contents, &userInfo)
	if err != nil {
		return nil, err
	}
	return &userInfo.Email, nil
}