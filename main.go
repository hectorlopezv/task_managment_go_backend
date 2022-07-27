package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       "2020-01-01",
	}
	tasks = append(tasks, task)
	fmt.Println("your task are ", tasks)
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       "2020-01-02",
	}
	tasks = append(tasks, task2)
}
func homePage(http.ResponseWriter, *http.Request) {
	fmt.Println("Home page Handler")

}
func getTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page Handler")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}
func getTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page Handler")

	taskId := mux.Vars(r)["id"]
	flag := false

	for _, task := range tasks {
		if task.ID == taskId {
			json.NewEncoder(w).Encode(task)
			flag = true
			break
		}
	}

	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}
func createTask(http.ResponseWriter, *http.Request) {
	fmt.Println("Home page Handler")

}
func deleteTask(http.ResponseWriter, *http.Request) {
	fmt.Println("Home page Handler")

}
func updateTask(http.ResponseWriter, *http.Request) {
	fmt.Println("Home page Handler")

}
func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("GET")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("UPDATE")
	log.Fatal(http.ListenAndServe(":8082", router))
}
func main() {
	allTasks()
	handleRoutes()
	fmt.Println("Hello Fluter task Manager Crud")
}
