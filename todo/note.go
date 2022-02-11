package todo

import (
	"errors"
	"fmt"
)

func Note(file string, id int, note string) error {
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
	todo.Note = note

	// 保存
	if err := todos.Save(file); err != nil {
		return err
	}

	return nil
}
