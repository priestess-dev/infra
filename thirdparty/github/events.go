package github

import "time"

type Repository struct {
	ID          int       `json:"id"`                  // id
	Name        string    `json:"name"`                // name
	FullName    string    `json:"full_name,omitempty"` // full name
	Url         string    `json:"url"`                 // url
	Private     bool      `json:"private,omitempty"`   // private
	Owner       User      `json:"owner,omitempty"`     // owner
	HtmlUrl     string    `json:"html_url,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	PushedAt    time.Time `json:"pushed_at,omitempty"`
}

type WatchEvent struct {
	Action string `json:"action"` // action (e.g. started)
}

type Commit struct {
	Sha    string `json:"sha"` // sha
	Author struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}
type PullRequestHead struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	Sha   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}
type PullRequest struct {
	Url                string    `json:"url"`
	Id                 int       `json:"id"`
	NodeId             string    `json:"node_id"`
	HtmlUrl            string    `json:"html_url"`
	DiffUrl            string    `json:"diff_url"`
	PatchUrl           string    `json:"patch_url"`
	IssueUrl           string    `json:"issue_url"`
	Number             int       `json:"number"`
	State              string    `json:"state"`
	Locked             bool      `json:"locked"`
	Title              string    `json:"title"`
	User               User      `json:"user"`
	Body               string    `json:"body"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	ClosedAt           time.Time `json:"closed_at"`
	MergedAt           time.Time `json:"merged_at"`
	MergeCommitSha     string    `json:"merge_commit_sha"`
	Assignee           User      `json:"assignee"`
	Assignees          []User    `json:"assignees"`
	RequestedReviewers []User    `json:"requested_reviewers"`
	//RequestedTeams     []interface{} `json:"requested_teams"`
	//Labels            []interface{} `json:"labels"`
	//Milestone         interface{} `json:"milestone"`
	Draft             bool            `json:"draft"`
	CommitsUrl        string          `json:"commits_url"`
	ReviewCommentsUrl string          `json:"review_comments_url"`
	ReviewCommentUrl  string          `json:"review_comment_url"`
	CommentsUrl       string          `json:"comments_url"`
	StatusesUrl       string          `json:"statuses_url"`
	Head              PullRequestHead `json:"head"`
}

type PushEvent struct {
	RepositoryId int      `json:"repository_id"` // repository id
	PushId       int64    `json:"push_id"`       // push id
	Size         int      `json:"size"`          // size
	DistinctSize int      `json:"distinct_size"` // distinct size
	Ref          string   `json:"ref"`           // ref
	Head         string   `json:"head"`          // head
	Before       string   `json:"before"`        // before
	Commits      []Commit `json:"commits"`       // commits
}

type CreateEvent struct {
	Ref          string `json:"ref"`           // ref
	RefType      string `json:"ref_type"`      // ref type
	MasterBranch string `json:"master_branch"` // master branch
	Description  string `json:"description"`   // description
	PusherType   string `json:"pusher_type"`   // pusher type
}

type ReleaseEvent struct {
	Action  string `json:"action"`
	Release struct {
		Url             string    `json:"url"`
		AssetsUrl       string    `json:"assets_url"`
		UploadUrl       string    `json:"upload_url"`
		HtmlUrl         string    `json:"html_url"`
		Id              int       `json:"id"`
		Author          User      `json:"author"`
		NodeId          string    `json:"node_id"`
		TagName         string    `json:"tag_name"`
		TargetCommitish string    `json:"target_commitish"` // target commitish (e.g. master)
		Name            string    `json:"name"`
		Draft           bool      `json:"draft"`
		Prerelease      bool      `json:"prerelease"`
		CreatedAt       time.Time `json:"created_at"`
		PublishedAt     time.Time `json:"published_at"`
		//Assets                          []interface{} `json:"assets"`
		TarballUrl                      string `json:"tarball_url"`
		ZipballUrl                      string `json:"zipball_url"`
		Body                            string `json:"body"`
		ShortDescriptionHtml            string `json:"short_description_html"`
		IsShortDescriptionHtmlTruncated bool   `json:"is_short_description_html_truncated"`
	} `json:"release"`
}

type PullRequestEvent struct {
	Action string `json:"action"`
	Number int    `json:"number"`
}

type PublicEvent struct {
}

type EventPayload interface {
	WatchEvent | PushEvent | CreateEvent | PublicEvent | ReleaseEvent
}

type Event[T EventPayload] struct {
	ID        string     `json:"id"`         // id
	Type      string     `json:"type"`       // type of event (e.g. PushEvent)
	Repo      Repository `json:"repo"`       // repository
	Actor     User       `json:"actor"`      // actor
	Payload   T          `json:"payload"`    // payload
	Public    bool       `json:"public"`     // public
	CreatedAt time.Time  `json:"created_at"` // created at
}
