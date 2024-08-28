package service_test

import (
	"testing"
	mocks "todo/mocks/todo/pkg/persistence"
	"todo/pkg/service"
	"todo/pkg/task"
)

func AddTest(t *testing.T) {
	// , currentList task.TaskList, TaskDescription string
	currentListMock := task.TaskList{Name: "todo",
		Tasks: []task.Task{
			{Id: "2345", Description: "pick up junk", Done: false},
			{Id: "3469", Description: "eat", Done: false},
		},
	}

	var TaskDescriptionMock = "go to school"

	var persistence = mocks.NewMockPersistence(t)

	service.Add(currentListMock, TaskDescriptionMock, persistence)
}
