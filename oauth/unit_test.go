//go:build !travis

package oauth

import (
	ih "github.com/priestess-dev/infra/v1/http"
	"net/http"
	"os"
	"testing"
)

func getClientConfig() *Config {
	return &Config{
		RedirectURL:  "http://127.0.0.1:8000/callback",
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		Scopes:       []string{"user:email"},
	}
}

func TestGithubHandler(t *testing.T) {
	// create server
	server := ih.NewServer(
		"127.0.0.1",
		8000,
	)
	// create github handler
	config := getClientConfig()
	t.Logf("config: %+v", config)
	githubHandler := NewGithubHandler(config)
	// add routes
	server.AddRoutes(
		ih.EndpointConfig{
			Path:    "/login",
			Method:  http.MethodGet,
			Handler: githubHandler.GinLoginHandler,
		},
		ih.EndpointConfig{
			Path:    "/callback",
			Method:  http.MethodGet,
			Handler: githubHandler.GinCallbackHandler,
		},
	)
	err := server.Start()
	if err != nil {
		t.Fatalf("server start error: %s", err.Error())
	}
}
