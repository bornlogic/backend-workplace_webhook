package github_test

import (
	"testing"

	"github.com/bornlogic/wiw/server/models/github"
)

var tcasesPush = []struct {
	name string
	push github.Push
	want string
}{
	{
		name: "all empty",
		push: github.Push{},
		want: `#  Push -> 
**Pusher:** 
**URL:** 
Repository URL: 
## **[WARNING] MISSING PULL REQUEST**`,
	}, {
		name: "all fields with PR",
		push: github.Push{
			Pusher: github.Pusher{
				Name: "pusher",
			},
			Ref:     "ref",
			Compare: "compare",
			Repository: github.Repository{
				Name: "repo name",
				URL:  "raw_url",
			},
			HeadCommit: github.HeadCommit{
				Message: "message commit here",
				URL:     "raw_url",
				Committer: github.Committer{
					Username: github.GithubUser,
				},
			},
		},
		want: `# repo name Push -> ref
**Pusher:** pusher
**URL:** compare
Repository URL: raw_url
## message commit here`,
	}, {
		name: "all fields without PR",
		push: github.Push{
			Pusher: github.Pusher{
				Name: "pusher",
			},
			Ref:     "ref",
			Compare: "compare",
			Repository: github.Repository{
				Name: "repo name",
				URL:  "raw_url",
			},
			HeadCommit: github.HeadCommit{
				Message: "message commit here",
				URL:     "raw_url",
				Committer: github.Committer{
					Username: "another user",
				},
			},
		},
		want: `# repo name Push -> ref
**Pusher:** pusher
**URL:** compare
Repository URL: raw_url
## **[WARNING] MISSING PULL REQUEST**`,
	},
}

func TestPushToMarkdown(t *testing.T) {
	for _, tc := range tcasesPush {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.push.ToMarkdown(); got != tc.want {
				t.Fatalf("%s != %s", got, tc.want)
			}
		})
	}
}
