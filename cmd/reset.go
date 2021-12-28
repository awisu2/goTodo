package cmd

import (
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

type ResetCmdArg struct {
	ID int
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if err := todo.Reset(rootArg.File, resetCmdArg.ID); err != nil {
			log.Panic(err)
		}
	},
}

var resetCmdArg = ResetCmdArg{}

func init() {
	resetCmd.Flags().IntVarP(&resetCmdArg.ID, "id", "", 0, "todo subject")
	resetCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(resetCmd)
}
