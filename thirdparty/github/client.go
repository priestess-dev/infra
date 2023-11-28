package github

import (
	ih "github.com/priestess-dev/infra/http"
	"golang.org/x/oauth2"
	"net/http"
)

type EndpointEnum string

const (
	EndpointGetUser    EndpointEnum = "https://api.github.com/user"
	EndpointListUsers  EndpointEnum = "ListUsers"
	EndpointCreateUser EndpointEnum = "CreateUser"
)

func (e EndpointEnum) String() string {
	return string(e)
}

type Endpoints interface {
	GetUser() (GetUserResponse, error)
	ListUsers(ListUsersRequest) (ListUsersResponse, error)
}

type Client struct {
	client *http.Client
	token  *oauth2.Token
}

func (c *Client) makeAuthRequest(url, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &http.Request{}, nil
	}
	req.Header.Add("Authorization", "token "+c.token.AccessToken)
	return req, nil
}

func (c *Client) GetUser() (GetUserResponse, error) {
	handler := ih.EndpointHandler[ih.EmptyRequest, GetUserResponse](c.client, EndpointGetUser.String(), http.MethodGet, nil, "")
	return handler(ih.EmptyRequest{})
}

func (c *Client) ListUsers(req ListUsersRequest) (ListUsersResponse, error) {
	handler := ih.EndpointHandler[ListUsersRequest, ListUsersResponse](c.client, EndpointListUsers.String(), http.MethodGet, nil, "")
	return handler(req)
}

func (c *Client) GetToken() *oauth2.Token {
	return c.token
}

func (c *Client) SetToken(token *oauth2.Token) {
	c.token = token
}

func NewGithubClient(client *http.Client, token *oauth2.Token) *Client {
	return &Client{
		client: client,
		token:  token,
	}
}
