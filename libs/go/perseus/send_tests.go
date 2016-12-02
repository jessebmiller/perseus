package perseus

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// TestSend assures that send is sending messages to namespaces correctly
func TestSend(t *testing.T) {

	tsets := []struct {
		in         string
		inNs       string
		outMessage string
		outNs      string
	}{
		{"message", "namespace", "message", "namespace"},
		{"a", "b", "a", "b"},
	}

	for _, tset := range tsets {
		// start a server that checks that the request is correct
		requestCount := 0
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				requestCount++
				if r.URL.Path != tset.outNs {
					t.Errorf("requested ns %v, want %v", r.URL.Path, tset.outNs)
				}
				sentMessage := r.URL.Query()["message"][0]
				if sentMessage != tset.outMessage {
					t.Errorf(
						"requested message %v, want %v",
						sentMessage,
						tset.outMessage,
					)
				}
			},
		))
		defer server.Close()

		// configure the environment to use that server
		os.Setenv("PERSEUS_HOST", server.URL)
		send(tset.in)

		// confirm that a call happened
		if requestCount != 1 {
			t.Errorf("expected 1 call, got %v", requestCount)
		}
	}
}
