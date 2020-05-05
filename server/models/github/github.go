package github

type Repository struct {
	Name string
	URL  string `json:"html_url"`
}

type User struct {
	Login string
}

func (u User) String() string {
	return u.Login
}
