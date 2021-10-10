package cmd

import (
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

type AddCmdArg struct {
	Subject string
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Add(rootArg.File, addCmdArg.Subject); err != nil {
			log.Panic(err)
		}
	},
}

var addCmdArg = AddCmdArg{}

func init() {
	addCmd.Flags().StringVarP(&addCmdArg.Subject, "subject", "s", "", "todo subject")
	addCmd.MarkFlagRequired("subject")

	rootCmd.AddCommand(addCmd)
}
