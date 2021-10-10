package cmd

import (
	"fmt"
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

type ListCmdArg struct {
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		json, err := todo.ListJson(rootArg.File)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(json))
	},
}

var listCmdArg = ListCmdArg{}

func init() {
	rootCmd.AddCommand(listCmd)
}
