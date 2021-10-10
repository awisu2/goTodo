package todo

import (
	"github.com/awisu2/goTodo/status"
	"github.com/awisu2/goUtils/files"
)

func Add(file string, subject string) error {
	// 既存データがあれば読み込み
	var todos Todos
	if err := files.IsExists(file); err != nil {
		todos = Todos{
			Todos: []Todo{},
		}
	} else {
		if err := read(file, &todos); err != nil {
			return err
		}
	}

	// 新規作成
	todo := Todo{
		ID:      todos.CreateNextId(),
		Subject: subject,
		Status:  status.NON,
	}
	todos.Add(todo)

	// 保存
	if err := todos.Save(file); err != nil {
		return err
	}

	return nil
}
