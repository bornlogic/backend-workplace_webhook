// sendToGroup is a cli for usage worksplace SendToGroup function
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/bornlogic/wiw/workplace"
)

// if access token isn't passed by flag, use value of env `WORKPLACE_ACCESS_TOKEN`
var DefaultAccessToken = os.Getenv("WORKPLACE_ACCESS_TOKEN")

// args for function for send to SendToGroup
var accessToken, groupID, formatting, message string

// print verbose mode
var isVerbose bool

func main() {
	flag.Parse()

	if err := checkArgs(accessToken, groupID, formatting, message); err != nil {
		log.Fatal(err)
	}

	verbose := func(s string) {
		if isVerbose {
			log.Println(s)
		}
	}

	verbose("do group send")
	resp, err := workplace.SendToGroup(accessToken, groupID, formatting, message)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("invalid status returned: %s", resp.Status)
	}

	verbose("read response")
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("read response: %s", err)
	}
	defer resp.Body.Close()

	verbose(fmt.Sprintf("response: %s", string(b)))
}

func checkArgs(accessToken, groupID, formatting, message string) error {
	if accessToken == "" {
		return fmt.Errorf("missing accessToken")
	}
	if groupID == "" {
		return fmt.Errorf("missing groupID")
	}
	if formatting == "" {
		return fmt.Errorf("missing formatting")
	}
	if message == "" {
		return fmt.Errorf("missing message")
	}
	return nil
}

func init() {
	const shortHandSuffix = "(shorthand)"

	const usageAccessToken = "access token used to connect with workplace api, if empty it will use the env `WORKPLACE_ACCESS_TOKEN`"
	flag.StringVar(&accessToken, "access-token", DefaultAccessToken, usageAccessToken)
	flag.StringVar(&accessToken, "t", DefaultAccessToken, usageAccessToken+shortHandSuffix)

	const usageGroupID = "group id of group for send the message"
	flag.StringVar(&groupID, "group-id", "", usageGroupID)
	flag.StringVar(&groupID, "g", "", usageGroupID+shortHandSuffix)

	const usageMessage = "message to send in given group"
	flag.StringVar(&message, "message", "", usageMessage)
	flag.StringVar(&message, "m", "", usageMessage+shortHandSuffix)

	const (
		usageFormatting   = "formatting of message, eg: PLAINTEXT, MARKDOWN"
		defaultFormatting = "PLAINTEXT"
	)
	flag.StringVar(&formatting, "formatting", defaultFormatting, usageFormatting)
	flag.StringVar(&formatting, "f", defaultFormatting, usageFormatting+shortHandSuffix)

	const usageVerbose = "prints feedback of operations"
	flag.BoolVar(&isVerbose, "verbose", false, usageVerbose)
	flag.BoolVar(&isVerbose, "v", false, usageVerbose+shortHandSuffix)

}
