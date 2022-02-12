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
	UUID    uuid.UUID `json:"uuid"`
	Name    string    `json:"name"`
	Content string    `json:"content"`
}

type AllTasks []task

var tasks = AllTasks{
	{
		UUID:    uuid.New(),
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

func getTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskUUId, err := uuid.Parse(vars["uuid"])

	if err != nil {
		fmt.Fprintf(w, "Invalid UUID")
		return
	}

	for _, t := range tasks {
		if t.UUID == taskUUId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
		}
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Insert a valid Task")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.UUID = uuid.New()
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskUUId, err := uuid.Parse(vars["uuid"])

	if err != nil {
		fmt.Fprintf(w, "Invalid UUID")
		return
	}

	for i, t := range tasks {
		if taskUUId == t.UUID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v hasbeen remove successfully", taskUUId)

		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskUUId, err := uuid.Parse(vars["uuid"])
	var updateTask task

	if err != nil {
		fmt.Fprintf(w, "Invalid UUID")
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}

	json.Unmarshal(reqBody, &updateTask)

	for i, t := range tasks {
		if t.UUID == taskUUId {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updateTask.UUID = taskUUId
			tasks = append(tasks, updateTask)

			fmt.Fprintf(w, "The task with UUID %v has been updated successfully", taskUUId)
		}
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{uuid}", getTask).Methods("GET")
	router.HandleFunc("/tasks/{uuid}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{uuid}", updateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))
}
