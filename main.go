package main

import (
	"fmt"
	"todo/pkg/task"
)

func main() {
	// who? ME, IRISTHESNAIL !!!
	// insert in the list
	// change from the list
	// mark as  doone
	// delete from list
	taskList := &task.TaskList{}
	task := &task.Task{
		Description: "bruh",
		Done:        false,
	}
	taskList.Add(task)

	got, _ := taskList.Get(task.Id)

	fmt.Printf("task = %v", got.Description)

}
