package server_test

import (
	"testing"

	"github.com/bornlogic/wiw/server"
)

// ExpectedServer is the interface with expected functions for the new server
type ExpectedServer interface {
	Run(port string) error
}

func TestNewServer(t *testing.T) {
	// should be possible create a new server with expected functions
	// it fails on compilation time if don't implements
	var _ ExpectedServer = server.NewServer("accessToken", "formatting")
}
