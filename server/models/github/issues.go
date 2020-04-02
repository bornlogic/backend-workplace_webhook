package github

import "fmt"

// Issues is the main abstraction about github webhook payload of issues
type Issues struct {
	Action     string
	Issue      Issue
	Repository Repository
	Sender     User
}

func (iss Issues) ToMarkdown() string {
	repo := iss.Repository
	i := iss.Issue
	return fmt.Sprintf(`# Issue %s -> %s
**URL:** %s
**Title:** %s
**User:** %s
Sender: %s
Repository URL: %s`, iss.Action, repo.Name, i.URL, i.Title, i.User, iss.Sender, repo.URL)
}

type Issue struct {
	URL   string `json:"html_url"`
	Title string
	User  User
}

func (i Issue) String() string {
	return fmt.Sprintf(`URL: %s  
Title: %s  
User: %s  `, i.URL, i.Title, i.User)
}
