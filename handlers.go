package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ns returns the namespace this request is in
func ns(r *http.Request) string {
	return r.URL.Path
}

// rootHandler routes requests based on their method
func rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, get(r, DEPS["store"].(Store)))
	case "POST":
		length, err := post(r, DEPS["store"].(Store))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, length)
	default:
		code := http.StatusMethodNotAllowed
		http.Error(w, http.StatusText(code), code)
	}
}

// get handles get requests given some context state
func get(r *http.Request, s Store) string {
	return strings.Join(s.get(ns(r)), "\n")
}

// postHandler handles post requests
func post(r *http.Request, s Store) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}
	messages, present := r.Form["message"]
	fmt.Println(r.URL.Query())
	if !present {
		return "", errors.New("Missing message")
	}
	s.add(ns(r), messages[0])
	return strconv.Itoa(len(messages[0])), nil
}
