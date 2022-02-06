package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type task struct {
	ID      uuid.UUID `json:"ID"`
	Name    string    `json:"Name"`
	Content string    `json:"Content"`
}

type AllTasks []task

var tasks = AllTasks{
	{
		ID:      uuid.New(),
		Name:    "First Task",
		Content: "First Content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to Go API")
	if err != nil {
		return
	}
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = uuid.New()
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := uuid.Parse(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
		return
	}

	for _, t := range tasks {
		if t.ID == taskId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}
