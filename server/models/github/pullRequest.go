package github

import "fmt"

// PullRequests is the abstraction about github webhook payload of pull_request
type PullRequest struct {
	Action     string
	Info       PullRequestInfo `json:"pull_request"`
	Repository Repository
}

func (pr PullRequest) ToMarkdown() string {
	pri := pr.Info
	repo := pr.Repository
	return fmt.Sprintf(`# Pull Request %s -> %s
**URL:** %s
**Title:** %s
User: %s
Repository URL: %s`, pr.Action, repo.Name, pri.URL, pri.Title, pri.User, repo.URL)
}

type PullRequestInfo struct {
	URL   string `json:"html_url"`
	Title string
	User  User
}
