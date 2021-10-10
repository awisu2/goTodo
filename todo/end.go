package todo

import (
	"errors"
	"fmt"
	"log"

	"github.com/awisu2/goTodo/status"
)

func End(file string, id int) error {
	// 既存データがあれば読み込み
	var todos Todos
	if err := read(file, &todos); err != nil {
		return err
	}

	// ステータス更新
	todo := todos.Get(id)
	if todo == nil {
		text := fmt.Sprintf("nothing todo. id: %d", id)
		return errors.New(text)
	}
	todo.Status = status.END

	// 保存
	log.Println(todos, todo)
	if err := todos.Save(file); err != nil {
		return err
	}

	return nil
}
