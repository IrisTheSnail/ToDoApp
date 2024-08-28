package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

// switchListCmd represents the switchList command
var switchListCmd = &cobra.Command{
	Use:   "switchList <name of exiting list of tasks>",
	Short: "Switches to a TaskList",
	Run: func(cmd *cobra.Command, args []string) {
		if reflect.TypeOf(args[0]).Kind() == reflect.String {
			switchTaskList(args[0])
		}
	},
}

func switchTaskList(TargetList string) {
	currentList.Name = TargetList
	fmt.Printf("switched to %v", currentList)
}
