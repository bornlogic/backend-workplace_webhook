package main_test

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

var cmdPath = "github.com/bornlogic/wiw/cmd/workplace/groupSend"
var accessTokenTest = os.Getenv("WORKPLACE_ACCESS_TOKEN")
var groupIdTest = os.Getenv("WORKPLACE_GROUP_ID_TEST")

// TestGroupSend tests if workflow to send a message to an group is worked
// For run this test you need to assign the envs VARS WORKPLACE_ACCESS_TOKEN and WORKPLACE_GROUP_ID_TEST
func TestGroupSend(t *testing.T) {
	runGroupSend(t, accessTokenTest, groupIdTest, "PLAINTEXT", "workplace group send integration test")
}

func runGroupSend(t *testing.T, accessToken, groupId, formatting, message string) {
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
		t.Fatalf(err)
	}
}
