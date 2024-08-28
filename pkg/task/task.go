package task

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

type TaskList struct {
	Name  string
	Tasks []Task
}

type Task struct {
	Id          string
	Description string
	Done        bool
}

func (taskList *TaskList) Add(task Task) string {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s-%v", task.Description, time.Now())))
	task.Id = base64.StdEncoding.EncodeToString(hasher.Sum(nil)[:8])
	taskList.Tasks = append(taskList.Tasks, task)
	return task.Id
}

func (taskList *TaskList) Get(id string) (Task, bool) {
	for _, task := range taskList.Tasks {
		if task.Id == id {
			return task, true
		}
	}

	return Task{}, false
}

func (taskList *TaskList) GetAll() []Task {
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

func (taskList *TaskList) Update(id string, updatedTask Task) error {
	for index, task := range taskList.Tasks {
		if task.Id == id {
			taskList.Tasks[index] = updatedTask
			return nil
		}
	}

	return fmt.Errorf("couldn't find task with id=%s", id)
}

func (taskList *TaskList) MarkDone(FinishedTask *Task) {
	FinishedTask.Done = true
	taskList.Update(FinishedTask.Id, *FinishedTask)

}

// bro but who accesses task by thier id i think we ll need search in taskss
func (taskList *TaskList) SearchByDescription(SearchedTaskDescription string) Task {
	for _, task := range taskList.Tasks {
		if task.Description == SearchedTaskDescription {
			return task
		}
	}
	return Task{} //is this good?
}
