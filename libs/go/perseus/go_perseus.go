package perseus

import (
	"bytes"
	"fmt"
	"github.com/jessebmiller/cfg"
	"net/http"
	"net/url"
)

// Send sends a message (async) to the configured namespace
func Send(message string) {
	go send(message)
}

// Sendf sends a formatted message
func Sendf(format string, a ...interface{}) {
	go send(fmt.Sprintf(format, a...))
}

func send(message string) {
	if cfg.Get("PERSEUS_SEND", "true") == "false" {
		// someonee turned perseus off
		return
	}
	host := cfg.Get("PERSEUS_HOST", "http://perseus")
	ns := cfg.Get("PERSEUS_NS", "/default")
	vals := url.Values{"message": {message}}
	http.Post(
		host+ns,
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(vals.Encode()),
	)
}
