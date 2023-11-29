package oauth

type ProviderEnum string

const (
	Github ProviderEnum = "github"
	Google ProviderEnum = "google"
)

type Config struct {
	ClientID     string   `json:"client_id"`     // client id
	ClientSecret string   `json:"client_secret"` // client secret
	RedirectURL  string   `json:"redirect_url"`  // redirect url
	Scopes       []string `json:"scopes"`        // scopes
}
