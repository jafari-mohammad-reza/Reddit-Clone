package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	dto "github.com/reddit-clone/src/domains/user-domain/authentication/dto"
	"github.com/reddit-clone/src/share/config"
)

type GithubAuthService struct {
	cfg *config.Config
}

func NewGithubAuthService(cfg *config.Config) *GithubAuthService {
	return &GithubAuthService{
		cfg: cfg,
	}
}
func (service *GithubAuthService) GetAuthorizationUrl() string {
	return fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s", service.cfg.ApiKey.GithubClientId, service.cfg.Server.ApiUrl+"/api/v1/auth/github/callback")
}

func (service *GithubAuthService) Callback(code string) (*dto.GithubAccessTokenResponse, error) {
	reqBody := map[string]string{
		"client_id":     service.cfg.ApiKey.GithubClientId,
		"client_secret": service.cfg.ApiKey.GithubClientSecret,
		"code":          code,
		"redirect_uri":  service.cfg.Server.ApiUrl + "/api/v1/auth/github/callback",
	}

	reqJSON, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(reqJSON))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get access token: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ghResp dto.GithubAccessTokenResponse
	err = json.Unmarshal(respBody, &ghResp)
	if err != nil {
		return nil, err
	}

	return &ghResp, nil
}

func (service *GithubAuthService) Authorize(accessToken string) (*string, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "token "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to authorize: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := string(respBody)
	return &data, nil
}