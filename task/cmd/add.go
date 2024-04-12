package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/uguremirmustafa/task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")

		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Failed to add task", err.Error())
			return
		}

		fmt.Printf("Added \"%s\" to your task list\n", task)

	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
