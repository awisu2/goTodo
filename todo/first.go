package todo

import (
	"github.com/awisu2/goTodo/status"
)

type FirstOption struct {
	NoEnd bool
}

func First(file string, opt *FirstOption) (*Todo, error) {
	noEnd := false
	if opt != nil {
		noEnd = opt.NoEnd
	}

	// 既存データがあれば読み込み
	var todos Todos
	if err := read(file, &todos); err != nil {
		return nil, err
	}

	// ステータス更新
	if noEnd {
		for _, todo := range todos.Todos {
			if todo.Status != status.END {
				return &todo, nil
			}
		}
	} else {
		if len(todos.Todos) > 0 {
			return &todos.Todos[0], nil
		}
	}
	return nil, nil
}
