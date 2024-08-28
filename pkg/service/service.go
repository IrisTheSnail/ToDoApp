package service

import (
	"todo/pkg/persistence"
	"todo/pkg/task"
)

func Add(currentList task.TaskList, TaskDescription string, persistence persistence.Persistence) {
	var taskToAdd = task.Task{Id: "", Description: TaskDescription, Done: false}

	currentList.Tasks, _ = persistence.Load()
	currentList.Add(taskToAdd)
	persistence.Save(currentList.Tasks)
}

func I_did(currentList task.TaskList, TaskDescription string, persistence persistence.Persistence) error {
	currentList.Tasks, _ = persistence.Load()
	searchedTask := currentList.SearchByDescription(TaskDescription)
	currentList.MarkDone(&searchedTask)
	err := persistence.Save(currentList.Tasks)
	return err
}

func ListAll(currentList task.TaskList, persistence persistence.Persistence) []task.Task {
	currentList.Tasks, _ = persistence.Load()
	AllItems := currentList.GetAll()
	return AllItems
} //for evaluation

func Delete(currentList task.TaskList, TaskDescription string, persistence persistence.Persistence) error {
	currentList.Tasks, _ = persistence.Load()
	searchedTask := currentList.SearchByDescription(TaskDescription)
	currentList.Delete(searchedTask.Id)
	err := persistence.Save(currentList.Tasks)
	return err
}
