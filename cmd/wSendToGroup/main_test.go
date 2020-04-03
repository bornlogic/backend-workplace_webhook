// +build integration

package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

const (
	WAT  = "WORKPLACE_ACCESS_TOKEN"
	WGIT = "WORKPLACE_GROUP_ID_TEST"
)

var cmdPath = "github.com/bornlogic/wiw/cmd/wSendToGroup"

var accessTokenTest = os.Getenv(WAT)
var groupIdTest = os.Getenv(WGIT)

// TestSendToGroup tests if workflow to send a message to an group is worked
// For run this test you need to assign the envs WORKPLACE_ACCESS_TOKEN and WORKPLACE_GROUP_ID_TEST
func TestSendToGroup(t *testing.T) {
	checkRequirements(t)
	runSendToGroup(t, accessTokenTest, groupIdTest, "PLAINTEXT", "workplace group send integration test")
}

func runSendToGroup(t *testing.T, accessToken, groupId, formatting, message string) {
	cmd := exec.Command(
		"go", "run", cmdPath,
		"--verbose",
		"--access-token", accessToken,
		"--group-id", groupId,
		"--formatting", formatting,
		"--message", message,
	)
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr

	err := cmd.Run()
	t.Logf("stdout: %s", out.String())
	t.Logf("stderr: %s", outErr.String())

	if err != nil {
		t.Fatal(err)
	}
}

func checkRequirements(t *testing.T) {
	const errMsgFmt = "please set env %s for can run the test"
	if accessTokenTest == "" {
		t.Fatalf(errMsgFmt, WAT)
	}
	if groupIdTest == "" {
		t.Fatalf(errMsgFmt, WGIT)
	}
}
