package generic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bornlogic/wiw/server/handlers"
	"github.com/bornlogic/wiw/server/models/generic"
	"github.com/bornlogic/wiw/workplace"
	"github.com/julienschmidt/httprouter"
)

// NewHandler create a new generic handler with a internal group sender for Serve usage
func NewHandler(gsender workplace.GroupSender) *handler {
	return &handler{gsender}
}

// handler is the abstraction about generic handler
type handler struct {
	gsender workplace.GroupSender
}

// Serve receives a generic message to send to a given groupID
func (h *handler) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, handlers.EMissingData)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintf(w, handlers.EUnreadableBodyFmt, err)
		return
	}
	defer r.Body.Close()

	var (
		message string
		groupID = ps.ByName("groupID")
	)

	var payload generic.Payload
	if err := json.Unmarshal(b, &payload); err != nil {
		w.WriteHeader(handlers.StatusResponseNotExpected)
		fmt.Fprintf(w, handlers.EUnexpectedBodyFmt, "issues", err)
		return
	}
	message = payload.ToMarkdown()

	if message == "" {
		fmt.Fprintf(w, handlers.EUnexpectedBodyFmt, "generic", handlers.EUnexpectedEmptyMessage)
		return
	}

	resp, err := h.gsender.Send(groupID, message)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, handlers.ERequest, err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, handlers.EUnexpectedStatus, resp.Status)
		return
	}
}
