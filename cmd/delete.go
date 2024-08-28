package cmd

import (
	"reflect"
	"todo/pkg/persistence"
	"todo/pkg/service"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <list> <your task> OR <your task>",
	Short: "deletes a task",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 && len(args) != 1 {
			println(`not enough or too much args, use : 
			delete <TaskList> <task>
			or 
			delete <task>
			if you are in the right taskList
			`)
		}

		if len(args) == 1 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String {
			service.Delete(currentList, args[0], &persistence.CSVPersistence{})
		}

		if len(args) == 2 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String &&
			reflect.TypeOf(args[1]).Kind() == reflect.String {
			switchTaskList(args[0])
			service.Delete(currentList, args[0], &persistence.CSVPersistence{})
		}
	},
}
