package server

import (
	"encoding/json"
	"log"
	"net/http"
)

//Book ...
type Book struct {
	BookID int    `json:"bookId"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

//ServeRest serves a simple rest response
func ServeRest(w http.ResponseWriter, r *http.Request) {
	book := Book{
		BookID: 1,
		Title:  "Essentials of Software Engineering",
		Author: "Frank Tsui",
	}

	bytes, err := json.Marshal(book)
	if err != nil {
		log.Println("Error marshalling json: ", err)
		return
	}

	w.Header().Set("Content-Type", "applicaion/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
