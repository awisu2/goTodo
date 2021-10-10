package todo

func List(file string, todos *Todos) error {
	// 既存データがあれば読み込み
	if err := read(file, todos); err != nil {
		return err
	}

	return nil
}

func ListJson(file string) ([]byte, error) {
	var todos Todos
	List(file, &todos)

	return todos.Json()
}
