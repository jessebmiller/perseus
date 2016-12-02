package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	DEPS = make(map[string]interface{})
	DEPS["store"] = MapStore{make(map[string][]string)}
	http.HandleFunc("/", rootHandler)
	fmt.Println("Percius server starting up on 0.0.0.0:2120.")
	log.Fatal(http.ListenAndServe("0.0.0.0:2120", nil))
}
