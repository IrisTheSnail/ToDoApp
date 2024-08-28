package cmd

import (
	"fmt"
	"reflect"
	"todo/pkg/persistence"
	"todo/pkg/service"

	"github.com/spf13/cobra"
)

var I_didCmd = &cobra.Command{
	Use:   "I_did <list> <your task> OR <your task>",
	Short: "Marks a task as done",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 && len(args) != 1 {
			println(`not enough or too much args, use : 
			I_did <list> <your task>
			or 
			I_did <your task>
			if you are in the right taskList
			`)
		}

		if len(args) == 1 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String {

			service.I_did(currentList, args[0], &persistence.CSVPersistence{})
			fmt.Println("Niiice you did it! You finished the task : " + args[0])
		}

		if len(args) == 2 &&
			reflect.TypeOf(args[0]).Kind() == reflect.String &&
			reflect.TypeOf(args[1]).Kind() == reflect.String {
			switchTaskList(args[0])
			service.I_did(currentList, args[0], &persistence.CSVPersistence{})
			fmt.Println("Niiice you did it! You finished the task : " + args[0])
		}
	},
}
