package task_test

import (
	"testing"
	"todo/pkg/task"
)

func TestAdd(t *testing.T) {
	taskList := task.TaskList{}
	task := task.Task{
		Description: "new task",
		Done:        false,
	}

	taskList.Add(task)

	foundLen := len(taskList.Tasks)
	if foundLen != 1 {
		t.Fatalf("expected taskListo to have length %d, found %d", 1, foundLen)
	}

	if taskList.Tasks[0].Description != task.Description ||
		taskList.Tasks[0].Done != task.Done {
		t.Fatalf("expected taskList.Tasks[0] to be %v, found %v", task, taskList.Tasks[0])
	}
}

func TestGet(t *testing.T) {
	taskList := task.TaskList{}

	task1 := task.Task{
		Description: "new task",
		Done:        false,
	}

	task2 := task.Task{
		Description: "new task 2",
		Done:        false,
	}

	taskId2 := taskList.Add(task2)

	broughtTask, isItThere := taskList.Get(taskId2)

	if !isItThere {
		t.Fatal("expected broughtTask to be found")
	}

	if task2.Description != broughtTask.Description ||
		task2.Done != broughtTask.Done {
		t.Logf("%v", broughtTask.Done)
		t.Fatalf("expected broughtTask to be %v, found %v", task1, taskList.Tasks[0])
	}

	task3 := task.Task{
		Description: "new task 3",
		Done:        false,
	}

	_, isItThere = taskList.Get(task3.Id)

	if isItThere {
		t.Fatal("expected task3 not to be found")
	}

}

func TestGetAll(t *testing.T) {
	taskList := task.TaskList{}

	task1 := task.Task{
		Description: "new task",
		Done:        false,
	}

	task2 := task.Task{
		Description: "new task 2",
		Done:        false,
	}

	taskList.Add(task1)
	taskList.Add(task2)

	listOfAllTasks := taskList.GetAll()

	if len(listOfAllTasks) != len(taskList.Tasks) {
		t.Fatal("returned listOfAllTasks isn't same length as existing taskList")
	}

	if listOfAllTasks[0].Description != taskList.Tasks[0].Description ||
		listOfAllTasks[0].Done != taskList.Tasks[0].Done {
		t.Fatalf("returned a different list than the existant one %v, returned %v", taskList.Tasks, listOfAllTasks)
	}

}
