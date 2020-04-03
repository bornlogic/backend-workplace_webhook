package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bornlogic/wiw/server/models/github"
	"github.com/bornlogic/wiw/workplace"
	"github.com/julienschmidt/httprouter"
)

// HeaderEvent is returned from github to says what even occours
// see: https://developer.github.com/webhooks
const HeaderEvent = "X-Github-Event"

// Custom status code for response not expected
const StatusResponseNotExpected = 516

// Errors in string format useful to handle in Serve function
const (
	EMissingHeaderEvent = "missing " + HeaderEvent + " in header from response"
	EMissingData        = "missing data in body"
	EUnreadableBodyFmt  = "can't read body: %s"
	EUnexpectedBodyFmt  = "can't unmarshal %s from body: %s"
	ERequest            = "Error returned on request to api: %s"
	EUnexpectedStatus   = "unexpected status returned from api: %s"
)

// handler is the abstraction about github handler
type handler struct {
	gsender workplace.GroupSender
}

// NewHandler create a new github handler with a internal group sender for Serve usage
func NewHandler(gsender workplace.GroupSender) *handler {
	return &handler{gsender}
}

// Serve receives response from api and group param for internally send the message to group
// chose the message is based on the event sended
// events issues and pullRequest are mapped
// push events only in branch `master`is mapped
func (h *handler) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	events, ok := r.Header[HeaderEvent]
	if !ok {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, EMissingHeaderEvent)
		return
	}
	event := events[0]

	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, EMissingData)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, EUnreadableBodyFmt, err)
		return
	}
	defer r.Body.Close()

	var (
		message string
		groupID = ps.ByName("groupID")
	)

	switch event {
	case "issues":
		var issues github.Issues
		if err := json.Unmarshal(b, &issues); err != nil {
			w.WriteHeader(StatusResponseNotExpected)
			fmt.Fprintf(w, EUnexpectedBodyFmt, "issues", err)
			return
		}
		message = issues.ToMarkdown()

	case "pull_request":
		var pullRequest github.PullRequest
		if err := json.Unmarshal(b, &pullRequest); err != nil {
			w.WriteHeader(StatusResponseNotExpected)
			fmt.Fprintf(w, EUnexpectedBodyFmt, "pull_request", err)
			return
		}
		message = pullRequest.ToMarkdown()

	case "push":
		const refMaster = "refs/heads/master"
		var push github.Push
		if err := json.Unmarshal(b, &push); err != nil {
			w.WriteHeader(StatusResponseNotExpected)
			fmt.Fprintf(w, EUnexpectedBodyFmt, "push", err)
			return
		}
		if push.Ref == refMaster {
			message = push.ToMarkdown()
		}
	}

	if message == "" {
		return
	}

	resp, err := h.gsender.Send(groupID, message)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, ERequest, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, EUnexpectedStatus, resp.Status)
		return
	}
}
