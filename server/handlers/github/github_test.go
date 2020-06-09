package github_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"strings"

	"github.com/bornlogic/wiw/server/handlers"
	"github.com/bornlogic/wiw/server/handlers/github"
	"github.com/julienschmidt/httprouter"
)

// githubServeTests includes all cases for github.Serve function
// the cases are organized by `return` order
var githubServeTests = []struct {
	// test name
	name string
	// event returned in response
	event string
	// string body value to be sended to server
	bodyVal string
	// groupID sended by params
	groupID string
	// gsenderType is the type of fake gsender instanced
	// only filled if needs some err
	gsenderType string
	// expected status code
	wantCode int
	// emsg is the expected err message
	wantErr string
	// wantSent is the expectation about message to group
	wantSent bool
}{

	// Invalid request received
	{
		name:     "missing event",
		wantCode: http.StatusExpectationFailed,
		wantErr:  github.EMissingHeaderEvent,
	},
	{
		name:     "missing body",
		event:    "whatever",
		wantCode: http.StatusBadRequest,
		wantErr:  handlers.EMissingData,
	},
	{
		name:     "invalid body",
		event:    "whatever",
		bodyVal:  "badReader",
		wantCode: http.StatusNotAcceptable,
		wantErr:  fmt.Sprintf(handlers.EUnreadableBodyFmt, EReaderErr),
	},

	// Event not mapped
	{
		name:     "unmapped event",
		event:    "unmapped",
		bodyVal:  "whatever",
		wantCode: http.StatusOK,
	},

	// Invalid json body received in mapped events
	{
		name:     "unexpected body from issues",
		event:    "issues",
		bodyVal:  "invalid",
		wantCode: github.StatusResponseNotExpected,
		wantErr: fmt.Sprintf(handlers.EUnexpectedBodyFmt, "issues",
			"invalid character 'i' looking for beginning of value"),
	},
	{
		name:     "unexpected body from pull_request",
		event:    "pull_request",
		bodyVal:  "invalid",
		wantCode: github.StatusResponseNotExpected,
		wantErr: fmt.Sprintf(handlers.EUnexpectedBodyFmt, "pull_request",
			"invalid character 'i' looking for beginning of value"),
	},
	{
		name:     "unexpected body from push",
		event:    "push",
		bodyVal:  "invalid",
		wantCode: github.StatusResponseNotExpected,
		wantErr: fmt.Sprintf(handlers.EUnexpectedBodyFmt, "push",
			"invalid character 'i' looking for beginning of value"),
	},

	// Mapped events
	{
		name:     "sent from issues",
		event:    "issues",
		bodyVal:  `{"action": "whatever"}`,
		wantCode: http.StatusOK,
		wantSent: true,
	},
	{
		name:     "sent from pull_request",
		event:    "pull_request",
		bodyVal:  `{"action": "whatever"}`,
		wantCode: http.StatusOK,
		wantSent: true,
	},
	{
		name:     "not sent from push in unmapped branch",
		event:    "push",
		bodyVal:  `{"ref": "refs/head/whatever"}`,
		wantCode: http.StatusOK,
		wantSent: false,
	},
	{
		name:     "sent from push in master",
		event:    "push",
		bodyVal:  `{"ref": "refs/heads/master"}`,
		wantCode: http.StatusOK,
		wantSent: true,
	},

	// Error returned from send request to workplace
	{
		name:        "error send from issues",
		event:       "issues",
		bodyVal:     `{"action": "whatever"}`,
		gsenderType: "error",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.ERequest, "error"),
	},
	{
		name:        "error sent from pull_request",
		event:       "pull_request",
		bodyVal:     `{"action": "whatever"}`,
		gsenderType: "error",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.ERequest, "error"),
	},
	{
		name:        "error sent from push in master",
		event:       "push",
		bodyVal:     `{"ref": "refs/heads/master"}`,
		gsenderType: "error",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.ERequest, "error"),
	},

	// Invalid status code returned from api
	{
		name:        "invalid status from issues",
		event:       "issues",
		bodyVal:     `{"action": "whatever"}`,
		gsenderType: "invalid_status",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.EUnexpectedStatus, "invalid"),
	},
	{
		name:        "invalid status from pull_request",
		event:       "pull_request",
		bodyVal:     `{"action": "whatever"}`,
		gsenderType: "invalid_status",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.EUnexpectedStatus, "invalid"),
	},
	{
		name:        "invalid status from push",
		event:       "push",
		bodyVal:     `{"ref": "refs/heads/master"}`,
		gsenderType: "invalid_status",
		wantCode:    http.StatusServiceUnavailable,
		wantErr:     fmt.Sprintf(handlers.EUnexpectedStatus, "invalid"),
	},
}

type gsender struct {
	wasCalled bool
	errType   string
}

func (gs *gsender) Send(groupID, message string) (*http.Response, error) {
	gs.wasCalled = true
	switch gs.errType {
	case "invalid_status":
		return &http.Response{StatusCode: -1, Status: "invalid"}, nil
	case "error":
		return nil, fmt.Errorf("error")
	}
	return &http.Response{StatusCode: http.StatusOK}, nil
}

const EReaderErr = "bad reader"

type readerErr struct{}

func (r readerErr) Error() string { return EReaderErr }

type badReader int

func (badReader) Read([]byte) (int, error) { return 0, readerErr{} }

func newFakeGithubReq(event, bodyVal string) *http.Request {

	var req *http.Request
	method, target := "POST", "/github"
	switch bodyVal {
	case "":
		req = httptest.NewRequest(method, target, nil)
	case "badReader":
		req = httptest.NewRequest(method, target, badReader(0))
	default:
		req = httptest.NewRequest(method, target, strings.NewReader(bodyVal))
	}
	if event != "" {
		req.Header.Add(github.HeaderEvent, event)
	}
	return req
}

func newParams(groupID string) httprouter.Params {
	return httprouter.Params{
		httprouter.Param{"groupID", groupID},
	}
}

// TestGithubServe run all tests to test Serve internally
func TestGithubServe(t *testing.T) {
	for _, tc := range githubServeTests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := newFakeGithubReq(tc.event, tc.bodyVal)
			ps := newParams(tc.groupID)
			gs := &gsender{
				errType: tc.gsenderType,
			}
			handler := github.NewHandler(gs)
			handler.Serve(w, r, ps)
			if tc.wantCode != w.Code {
				t.Errorf("want %d, got %d", tc.wantCode, w.Code)
			}
			b, _ := ioutil.ReadAll(w.Body)
			if gotErr := string(b); tc.wantErr != gotErr {
				t.Errorf("want %q, got %q", tc.wantErr, gotErr)
			}
			if tc.wantSent && !gs.wasCalled {
				t.Errorf("Send method not called")
			}
		})
	}
}
