package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       "2020-01-02",
	}
	tasks = append(tasks, task2)
	fmt.Println("your task are ", tasks)

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

	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Home page Handler")
	var task Tasks

	json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(len(tasks) + 1)
	task.Date = time.Now().Format("01-02-2006")
	tasks = append(tasks, task)
	print(tasks)
	json.NewEncoder(w).Encode(tasks)

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page Handler")
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	flag := false
	for index, item := range tasks {
		if item.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			print(tasks)
			json.NewEncoder(w).Encode(tasks)
			flag = true
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page Handler")
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	json.NewDecoder(r.Body).Decode(&task)
	task.Date = time.Now().Format("01-02-2006")
	task.ID = mux.Vars(r)["id"]

	flag := false
	for index, item := range tasks {
		if item.ID == task.ID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			tasks = append(tasks, task)
			print(tasks)
			json.NewEncoder(w).Encode(tasks)
			flag = true
			return
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}

}
func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	handleRoutes()
	fmt.Println("Hello Fluter task Manager Crud")
}
