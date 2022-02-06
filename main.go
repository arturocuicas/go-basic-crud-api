package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type AllTasks []task

var tasks = AllTasks{
	{
		ID:      1,
		Name:    "First Task",
		Content: "First Content",
	},
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to Go API")
	if err != nil {
		return
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", IndexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
	// hola
}
