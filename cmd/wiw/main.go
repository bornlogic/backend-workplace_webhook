package main

import (
	"flag"
	"log"
	"os"

	"github.com/bornlogic/wiw/server"
)

var (
	workplaceAccessToken = os.Getenv("WORKPLACE_ACCESS_TOKEN")
)

// Formatting used internally is only markdown
const formatting = "MARKDOWN"

// Post to up the server passed as flag, default is :3000
var Port string

func main() {
	flag.Parse()
	svr := server.NewServer(workplaceAccessToken, formatting)
	log.Fatal(svr.Run(Port))
}

func init() {
	const (
		usagePort   = "port used for serve"
		defaultPort = ":3000"
	)
	flag.StringVar(&Port, "port", defaultPort, usagePort)
	flag.StringVar(&Port, "p", defaultPort, usagePort+"(shorthand)")
}
