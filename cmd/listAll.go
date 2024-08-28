package cmd

import (
	"fmt"
	"todo/pkg/persistence"
	"todo/pkg/service"

	"github.com/spf13/cobra"
)

// listAllCmd represents the listAll command
var listAllCmd = &cobra.Command{
	Use:   "listAll",
	Short: "Lists all tasks",

	Run: func(cmd *cobra.Command, args []string) {
		list := service.ListAll(currentList, &persistence.CSVPersistence{})
		for index := range list {
			fmt.Printf("%v \n", list[index])
		}
	},
}
