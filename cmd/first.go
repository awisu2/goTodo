package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

type firstCmdArg struct {
	todo.FirstOption
}

var firstCmd = &cobra.Command{
	Use:   "first",
	Short: "get todo",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := todo.First(rootArg.File, &_firstCmdArg.FirstOption)
		if err != nil {
			log.Panic(err)
		}

		if todo == nil {
			return
		}

		_json, err := json.Marshal(todo)
		if err != nil {
			log.Panic()
		}
		fmt.Println(string(_json))
	},
}

var _firstCmdArg = firstCmdArg{}

func init() {
	firstCmd.Flags().BoolVarP(&_firstCmdArg.NoEnd, "noend", "", false, "if it exists without end todo")

	rootCmd.AddCommand(firstCmd)
}
