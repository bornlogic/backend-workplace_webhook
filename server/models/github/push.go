package github

import "fmt"

const GithubUser = "web-flow"

type Push struct {
	Pusher     Pusher
	Ref        string
	Compare    string
	Repository Repository
	HeadCommit HeadCommit `json:"head_commit"`
}

func (p Push) ToMarkdown() string {
	hc := p.HeadCommit
	repo := p.Repository
	var msg string
	if hc.Committer.String() != GithubUser {
		msg += "## **[WARNING] MISSING PULL REQUEST**"
	} else {
		msg += fmt.Sprintf(`## %s`, hc.Message)
	}
	return fmt.Sprintf(`# %s Push -> %s
**Pusher:** %s
**URL:** %s
Repository URL: %s
%s`, repo.Name, p.Ref, p.Pusher, p.Compare, repo.URL, msg)
}

type Pusher struct {
	Name string
}

func (p Pusher) String() string {
	return p.Name
}

type HeadCommit struct {
	Message   string
	URL       string
	Committer Committer
}

type Committer struct {
	Username string
}

func (c Committer) String() string {
	return c.Username
}
