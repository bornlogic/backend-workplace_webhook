package generic

import "fmt"

type Payload struct {
	Title   string
	Message string
}

func (p Payload) ToMarkdown() string {
	if p.Title == "" && p.Message == "" {
		return ""
	}
	return fmt.Sprintf(`# %s
%s`, p.Title, p.Message)
}
