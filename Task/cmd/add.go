package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This adds tasks to the list",

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \" %s \" in to the list \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
