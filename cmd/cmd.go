package cmd

import (
	"github.com/awisu2/goTodo/cmd/note"
	"github.com/spf13/cobra"
)

type RootArg struct {
	File string
}

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "json file base todo command",
	Long:  "json file base todo command",
	Args:  cobra.NoArgs,
}

var rootArg = &RootArg{}

func init() {
	cobra.OnInitialize(onInitialize)

	rootCmd.PersistentFlags().StringVarP(&rootArg.File, "file", "f", "", "target json file")
	rootCmd.MarkPersistentFlagRequired("file")

	rootCmd.AddCommand(note.Cmd)
}

func onInitialize() {
	// anything ok
}

func Execute() error {
	return rootCmd.Execute()
}
