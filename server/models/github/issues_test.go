package github_test

import (
	"testing"

	"github.com/bornlogic/wiw/server/models/github"
)

var tcasesIssues = []struct {
	name   string
	issues github.Issues
	want   string
}{
	{
		name:   "all empty",
		issues: github.Issues{},
		want: `# Issue  -> 
**URL:** 
**Title:** 
**User:** 
Sender: 
Repository URL: `,
	}, {
		name: "all fields",
		issues: github.Issues{
			Action: "action",
			Issue: github.Issue{
				URL:   "raw_url",
				Title: "title",
				User: github.User{
					Login: "login",
				},
			},
			Repository: github.Repository{
				Name: "repo name",
				URL:  "raw_url",
			},
			Sender: github.User{
				Login: "login",
			},
		},
		want: `# Issue action -> repo name
**URL:** raw_url
**Title:** title
**User:** login
Sender: login
Repository URL: raw_url`,
	},
}

func TestIssuesToMarkdown(t *testing.T) {
	for _, tc := range tcasesIssues {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.issues.ToMarkdown(); got != tc.want {
				t.Fatalf("%s != %s", got, tc.want)
			}
		})
	}
}
