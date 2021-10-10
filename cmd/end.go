package cmd

import (
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

type EndCmdArg struct {
	ID int
}

var endCmd = &cobra.Command{
	Use:   "end",
	Short: "end todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.End(rootArg.File, endCmdArg.ID); err != nil {
			log.Panic(err)
		}
	},
}

var endCmdArg = EndCmdArg{}

func init() {
	endCmd.Flags().IntVarP(&endCmdArg.ID, "id", "", 0, "todo subject")
	endCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(endCmd)
}
