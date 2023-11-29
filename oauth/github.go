package oauth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/priestess-dev/infra/utils/random"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"io"
	"net/http"
)

type GithubHandler struct {
	oauthConfig *oauth2.Config
	Token       *oauth2.Token
	State       string
}

func (h *GithubHandler) GinLoginHandler(c *gin.Context) {
	authUrl := h.oauthConfig.AuthCodeURL(h.State)
	c.Redirect(http.StatusTemporaryRedirect, authUrl)
}

func (h *GithubHandler) GinCallbackHandler(c *gin.Context) {
	// handle state
	state := c.Query("state")
	fmt.Printf("\nstate: %s\n", state)
	if state != h.State {
		c.String(http.StatusBadRequest, "invalid state")
		return
	}

	code := c.Query("code")
	token, err := h.oauthConfig.Exchange(c, code)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	h.Token = token
	c.JSONP(http.StatusOK, token)
}

func (h *GithubHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	authUrl := h.oauthConfig.AuthCodeURL(h.State)
	http.Redirect(w, r, authUrl, http.StatusTemporaryRedirect)
}

func (h *GithubHandler) HandleCallback(w http.ResponseWriter, r *http.Request) {
	// handle state
	state := r.FormValue("state")
	if state != h.State {
		http.Error(w, "invalid state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := h.oauthConfig.Exchange(r.Context(), code)
	if err != nil {
		// todo: log error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.Token = token
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(fmt.Sprintf(`{"access_token": "%s"}`, token.AccessToken)))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *GithubHandler) GetUserInfo(token *oauth2.Token) (string, error) {
	// Must specify access token via Authorization header (https://docs.github.com/en/developers/apps/authorizing-oauth-apps#2-users-are-redirected-back-to-your-site-by-github)
	// Authorization: bearer OAUTH-TOKEN
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", token.AccessToken))
	response, err := http.DefaultClient.Do(req)
	fmt.Printf("%v", response)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed closing response body: %s", err.Error())
		}
	}(response.Body)
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func createGithubOAuthConfig(oauthConfig *Config) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  oauthConfig.RedirectURL,
		ClientID:     oauthConfig.ClientID,
		ClientSecret: oauthConfig.ClientSecret,
		Scopes:       oauthConfig.Scopes,
		Endpoint:     github.Endpoint,
	}
}

func NewGithubHandler(oauthConfig *Config) *GithubHandler {
	handler := &GithubHandler{
		oauthConfig: createGithubOAuthConfig(oauthConfig),
	}
	handler.State = random.RandString(16)
	return handler
}
