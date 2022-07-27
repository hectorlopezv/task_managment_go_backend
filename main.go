package main

import (
	"fmt"
)


type Tasks struct{
	ID string `json:"id"`
	TaskName string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date string `json:"date"`
}


var tasks []Tasks

func allTasks(){
	task := Tasks{
		ID: "1",
		TaskName: "Task 1",
		TaskDetail: "Task 1 Detail",
		Date: "2020-01-01",
	}
	tasks = append(tasks, task);
}
func main(){
	fmt.Println("Hello World");
}