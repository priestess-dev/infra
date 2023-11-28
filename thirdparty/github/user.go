package github

type User struct {
	ID           int    `json:"id"`            // id
	Login        string `json:"login"`         // login name
	DisplayLogin string `json:"display_login"` // display login name
	GravatarID   string `json:"gravatar_id"`   // gravatar id
	URL          string `json:"url"`           // url
	AvatarURL    string `json:"avatar_url"`    // avatar url
}

// GetUserResponse is the response of get authenticated user
// ref: https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#get-the-authenticated-user
type GetUserResponse struct {
	ID          int    `json:"id"`           // id
	Login       string `json:"login"`        // login name
	Name        string `json:"name"`         // full name
	Email       string `json:"email"`        // email
	PublicRepos int    `json:"public_repos"` // public repos
	AvatarURL   string `json:"avatar_url"`   // avatar url
	URL         string `json:"url"`          // url
}

type ListUsersRequest struct {
	Since   int `json:"since"`    // A user ID. Only return users with an ID greater than this ID
	PerPage int `json:"per_page"` // results per page (max 100), default 30
}

// ListUsersResponse is the response of list users
// ref: https://docs.github.com/en/rest/users/users?apiVersion=2022-11-28#update-the-authenticated-user
type ListUsersResponse []GetUserResponse
