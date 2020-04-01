// Package workplace gives a way to integrate with graph api of facebook
// see https://developers.facebook.com/docs/workplace/reference/graph-api/post/
package workplace

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// GroupSend performs a postage in feed using markdown formatting
// create request with the requirements of facebook developers api
// see: https://developers.facebook.com/docs/groups-api/common-uses#posting-on-a-group
func GroupSend(accessToken, groupID, formatting, message string) (*http.Response, error) {
	const (
		method = "POST"
		URLFmt = "https://graph.facebook.com/%s/feed"
	)
	var (
		url    = fmt.Sprintf(URLFmt, groupID)
		client = &http.Client{}
	)
	req, err := newGroupSendRequest(method, url, accessToken, formatting, message)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func newGroupSendRequest(method, url, accessToken, formatting, message string) (*http.Request, error) {
	return http.NewRequest(method, url, newGroupSendBody(accessToken, formatting, message))
}

func newGroupSendBody(accessToken, formatting, message string) *strings.Reader {
	return strings.NewReader(
		url.Values{
			"access_token": {accessToken},
			"formatting":   {formatting},
			"message":      {message},
		}.Encode(),
	)
}

// GroupSender is the main abstraction about post a message in one specific group
type GroupSender interface {
	Send(groupID, message string) (*http.Response, error)
}

// groupSender implements GroupSender interface
// guards internally the accessToken for send messages in a given group
// guards internally the formatting to send messages
type groupSender struct {
	accessToken string
	formatting  string
}

// Send a given message to a given group with internal accessToken
func (g *groupSender) Send(groupID, message string) (*http.Response, error) {
	return GroupSend(g.accessToken, groupID, g.formatting, message)
}

// NewGroupSender create a new groupSender with the given accessToken
func NewGroupSender(accessToken, formatting string) *groupSender {
	return &groupSender{accessToken, formatting}
}
