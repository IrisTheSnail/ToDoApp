package task

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type TaskList struct {
	Name  string
	Tasks []*Task
}

type Task struct {
	Id          string
	Description string
	Done        bool
}

func (taskList *TaskList) Add(task *Task) {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%v", task.Description, time.Now())))
	task.Id = string(hasher.Sum(nil))[:8]
	taskList.Tasks = append(taskList.Tasks, task)
}

func (taskList *TaskList) Get(id string) (*Task, bool) {
	for _, task := range taskList.Tasks {
		if task.Id == id {
			return task, true
		}
	}

	return &Task{}, false
}

func (taskList *TaskList) GetAll() []*Task {
	return taskList.Tasks
}

func (taskList *TaskList) Delete(id string) error {
	for index, task := range taskList.Tasks {
		if task.Id == id {
			taskList.Tasks = append(taskList.Tasks[:index], taskList.Tasks[index+1:]...)
			return nil
		}
	}

	return fmt.Errorf("couldn't find task with id=%s", id)
}

func (taskList *TaskList) Update(id string, updatedTask *Task) error {
	for index, task := range taskList.Tasks {
		if task.Id == id {
			taskList.Tasks[index] = updatedTask
			return nil
		}
	}

	return fmt.Errorf("couldn't find task with id=%s", id)
}

// bro but who accesses task by thier id i think we ll need search in taskss
