package perseus

import (
	"bytes"
	"github.com/jessebmiller/cfg"
	"net/http"
	"net/url"
)

// Send sends a message (async) to the configured namespace
func Send(message string) {
	go send(message)
}

func send(message string) {
	if cfg.Get("PERSEUS_SEND", "true") == "false" {
		// someonee turned perseus off
		return
	}
	host := cfg.Get("PERSEUS_URL", "http://perseus")
	ns := cfg.Get("PERSEUS_NS", "/default")
	vals := url.Values{"message": {message}}
	resp, _ := http.Post(
		host+ns,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(vals.Encode()),
	)
	resp.Body.Close()
}
