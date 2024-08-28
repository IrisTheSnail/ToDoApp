package cmd

import (
	"reflect"
	"todo/pkg/persistence"
	"todo/pkg/service"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <list> <your task> OR <your task>",
	Short: "adds a new task",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 && len(args) != 1 {
			println(`not enough or too much args, use : 
			add <TaskList> <your new task>
			or 
			add <your new task>
			if you are in the right taskList
			`)
		}

		if len(args) == 1 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String {
			service.Add(currentList, args[0], &persistence.CSVPersistence{})
		}

		if len(args) == 2 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String &&
			reflect.TypeOf(args[1]).Kind() == reflect.String {
			switchTaskList(args[0])
			service.Add(currentList, args[0], &persistence.CSVPersistence{})
		}

	},
}

// todo add <list> <new teak> OR <new task>
// todo switchTaskList <another list>
// todo I_did <list> <your teak> OR <your task>
// todo list <list> (coming soon ?)
// todo listAll
// todo delete <list> <teak>
