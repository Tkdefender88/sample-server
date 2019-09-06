package main

import (
	"fmt"
	"log"
	"net/http"
	"sample-server/server"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/book", server.ServeRest)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
