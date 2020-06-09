package server

import (
	"net/http"

	"github.com/bornlogic/wiw/server/handlers/generic"
	"github.com/bornlogic/wiw/server/handlers/github"
	"github.com/bornlogic/wiw/workplace"
	"github.com/julienschmidt/httprouter"
)

// server is the main abstraction about server used for start handlers
type server struct {
	router               *httprouter.Router
	workplaceAccessToken string
	formatting           string
}

// NewServer create new server with the given workplace access token and message formatting
func NewServer(workplaceAccessToken, formatting string) *server {
	return &server{
		router:               httprouter.New(),
		workplaceAccessToken: workplaceAccessToken,
		formatting:           formatting,
	}
}

// setupHandlers configure all handlers exposed in server
func (s *server) setupHandlers() {
	s.router.POST("/github/:groupID", github.NewHandler(
		workplace.NewGroupSender(s.workplaceAccessToken, s.formatting),
	).Serve)
	s.router.POST("/:groupID", generic.NewHandler(
		workplace.NewGroupSender(s.workplaceAccessToken, s.formatting),
	).Serve)
}

// Run the server with specified port
func (s *server) Run(port string) error {
	s.setupHandlers()
	return http.ListenAndServe(port, s.router)
}
