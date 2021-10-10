package main

import (
	"github.com/awisu2/goTodo/cmd"
)

func main() {
	// INFO: ここでエラーハンドリングすると、余計なログ出力になるのでキャッチしない
	// cobra の requireエラーは *errors.errorString だったソース覗いても専用のErrorを用意しないみたいなので判別不可
	cmd.Execute()
}
