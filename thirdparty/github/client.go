package github

import (
	ih "github.com/priestess-dev/infra/http"
	"golang.org/x/oauth2"
	"net/http"
	"strings"
)

type EndpointEnum string

const (
	EndpointGetUser        EndpointEnum = "https://api.github.com/user"
	EndpointListUsers      EndpointEnum = "https://api.github.com/users"
	EndpointListUserEvents EndpointEnum = "https://api.github.com/users/:user/events/public"
)

func (e EndpointEnum) String() string {
	return string(e)
}

type Client interface {
	UserEndpoints
	EventEndPoints
}

type UserEndpoints interface {
	GetUser() (GetUserResponse, error)
	ListUsers(ListUsersRequest) (ListUsersResponse, error)
}

type EventEndPoints interface {
	ListPublicEventsForUser(GetUserEventsRequest) (GetUserEventsResponse, error)
}

type client struct {
	client *http.Client
	token  *oauth2.Token
}

func (c *client) makeAuthRequest(url, method string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return &http.Request{}, nil
	}
	req.Header.Add("Authorization", "token "+c.token.AccessToken)
	return req, nil
}

func (c *client) GetUser() (GetUserResponse, error) {
	handler := ih.EndpointHandler[ih.EmptyRequest, GetUserResponse](c.client, EndpointGetUser.String(), http.MethodGet, nil, "")
	return handler(ih.EmptyRequest{})
}

func (c *client) ListUsers(req ListUsersRequest) (ListUsersResponse, error) {
	handler := ih.EndpointHandler[ListUsersRequest, ListUsersResponse](c.client, EndpointListUsers.String(), http.MethodGet, nil, "")
	return handler(req)
}

func (c *client) ListPublicEventsForUser(req GetUserEventsRequest) (GetUserEventsResponse, error) {
	if req.PerPage == 0 {
		req.PerPage = 30
	}
	if req.Page == 0 {
		req.Page = 1
	}
	handler := ih.EndpointHandler[GetUserEventsRequest, GetUserEventsResponse](
		c.client,
		strings.ReplaceAll(EndpointListUserEvents.String(), ":user", req.Username),
		http.MethodGet,
		nil,
		"",
	)
	return handler(req)
}

func (c *client) GetToken() *oauth2.Token {
	return c.token
}

func (c *client) SetToken(token *oauth2.Token) {
	c.token = token
}

func NewClient(httpClient *http.Client, token *oauth2.Token) Client {
	return &client{
		client: httpClient,
		token:  token,
	}
}
