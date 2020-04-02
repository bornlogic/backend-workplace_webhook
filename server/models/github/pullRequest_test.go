package github_test

import (
	"testing"

	"github.com/bornlogic/wiw/server/models/github"
)

var tcasesPullRequest = []struct {
	name string
	pr   github.PullRequest
	want string
}{
	{
		name: "empty fields",
		pr:   github.PullRequest{},
		want: `# Pull Request  -> 
**URL:** 
**Title:** 
User: 
Repository URL: `,
	}, {
		name: "all fields",
		pr: github.PullRequest{
			Action: "action",
			Info: github.PullRequestInfo{
				User: github.User{
					Login: "login",
				},
			},
			Repository: github.Repository{
				Name: "repo name",
				URL:  "raw_url",
			},
		},
		want: `# Pull Request action -> repo name
**URL:** 
**Title:** 
User: login
Repository URL: raw_url`,
	},
}

func TestPullRequestToMarkdown(t *testing.T) {
	for _, tc := range tcasesPullRequest {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.pr.ToMarkdown(); got != tc.want {
				t.Fatalf("%s != %s", got, tc.want)
			}
		})
	}
}
