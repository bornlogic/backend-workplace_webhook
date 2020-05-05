package workplace

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// tcasesSendToGroup shares cases for creation of body based on given args
var tcasesSendToGroup = []struct{ name, accessToken, formatting, message, want string }{
	{
		name:        "empty inputs",
		accessToken: "",
		formatting:  "",
		message:     "",
		want:        "access_token=&formatting=&message=",
	}, {
		name:        "missing access_token",
		accessToken: "",
		formatting:  "MARKDOWN",
		message:     "#HelloWorld",
		want: ("access_token=" +
			"&formatting=MARKDOWN" +
			"&message=%23HelloWorld"),
	}, {
		name:        "missing formatting",
		accessToken: "YXdkIGF3ZCBhZCBhd2QgYSBhZ",
		formatting:  "",
		message:     "#HelloWorld",
		want: ("access_token=YXdkIGF3ZCBhZCBhd2QgYSBhZ" +
			"&formatting=" +
			"&message=%23HelloWorld"),
	}, {
		name:        "missing message",
		accessToken: "YXdkIGF3ZCBhZCBhd2QgYSBhZ",
		formatting:  "PLAINTEXT",
		message:     "",
		want: ("access_token=YXdkIGF3ZCBhZCBhd2QgYSBhZ" +
			"&formatting=PLAINTEXT" +
			"&message="),
	}, {
		name:        "all filled inputs",
		accessToken: "YXdkIGF3ZCBhZCBhd2QgYSBhZ",
		formatting:  "MARKDOWN",
		message:     "#HelloWorld",
		want: ("access_token=YXdkIGF3ZCBhZCBhd2QgYSBhZ" +
			"&formatting=MARKDOWN" +
			"&message=%23HelloWorld"),
	},
}

// TestSendToGroupBody tests if sended args return the expected body
func TestSendToGroupBody(t *testing.T) {
	for _, tc := range tcasesSendToGroup {
		t.Run(tc.name, func(t *testing.T) {
			got := newSendToGroupBody(tc.accessToken, tc.formatting, tc.message)
			mustStrEqReader(t, tc.want, got)
		})
	}
}

// mustStrEqReader verify if given string is equal to given reader
// read string from reader and test equality
func mustStrEqReader(t *testing.T, s string, r io.Reader) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("read all: %s", err)
	}
	if ss := string(b); s != ss {
		t.Errorf("%s != %s", s, ss)
	}
}

// TestSendToGroupRequest verify if the given request created could be done in one server
func TestSendToGroupRequest(t *testing.T) {
	var methods = [4]string{"POST", "GET", "PUT", "OPTIONS"}
	len_methods := len(methods)
	rand.Seed(time.Now().UnixNano())
	for _, tc := range tcasesSendToGroup {
		t.Run(tc.name, func(t *testing.T) {
			method := methods[rand.Intn(len_methods)]
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != method {
					t.Errorf("method: want %s got %s", method, r.Method)
				}
				mustStrEqReader(t, tc.want, r.Body)
			}))
			defer ts.Close()
			req, err := newSendToGroupRequest(method, ts.URL, tc.accessToken, tc.formatting, tc.message)
			if err != nil {
				t.Errorf("new group post request: %s", err)
			}
			c := &http.Client{}
			if _, err = c.Do(req); err != nil {
				t.Errorf("do group post request: %s", err)
			}
		})
	}
}
