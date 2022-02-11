package note

import (
	"fmt"
	"log"

	"github.com/awisu2/goTodo/todo"
	"github.com/spf13/cobra"
)

var params = struct {
	Id   int
	Note string
}{}

var Cmd = &cobra.Command{
	Use:   "note",
	Short: "set any note(text)",
	Args:  cobra.ArbitraryArgs, // 引数設定(ArbitraryArgs: なんでもOK)
	// argsには下記で設定しているflag引数("-x arg")以外の値がセットされる
	Run: func(cmd *cobra.Command, args []string) {
		file, err := cmd.PersistentFlags().GetString("file")
		if err != nil {
			log.Fatalln(err)
		}

		todo.Note(file, params.Id, params.Note)
	},
}

func initConfig() {
	// viper(config値/ファイルに強いmodule)などを実行
}

// go 標準関数: 実行時に処理
func init() {
	// 引数が解析されたあと、Runの実行前に実行されるイベント設定
	cobra.OnInitialize(initConfig)

	// 引数設定
	flags, _ := Cmd.Flags(), Cmd.PersistentFlags()

	flags.IntVarP(&params.Id, "id", "i", 0, "todo id")
	flags.StringVarP(&params.Note, "note", "", "", "todo note")

	// 必須指定
	requireds := []string{"id", "note"}
	persistentRequireds := []string{}
	for _, required := range requireds {
		Cmd.MarkFlagRequired(required)
	}
	for _, required := range persistentRequireds {
		Cmd.MarkPersistentFlagRequired(required)
	}

	// 配下コマンドの追加
	Cmd.AddCommand(subCmd)
}

// 参考実装では、cmd.Execute()で実行できるようにしている
func Execute() error {
	return Cmd.Execute()
}

// サブコマンド: ごちゃつくのでほかファイルに(可能な限り他パッケージで)
var subCmd = &cobra.Command{
	Use: "sub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("")
	},
}
