package todo

import (
	"encoding/json"

	"github.com/awisu2/goTodo/status"
	"github.com/awisu2/goUtils/files"
)

type TodoStatus string

type Todo struct {
	ID      int           `json:"id"`
	Subject string        `json:"subject"` // main content
	Status  status.Status `json:"status"`
	Note    string        `json:"note"` // any/sub text
}

type Todos struct {
	Todos []Todo `json:"todos"`
}

func (t *Todos) CreateNextId() int {
	maxId := 0
	for _, todo := range t.Todos {
		if maxId < todo.ID {
			maxId = todo.ID
		}
	}
	return maxId + 1
}

func (t *Todos) Get(id int) *Todo {
	for i := range t.Todos {
		if t.Todos[i].ID == id {
			return &t.Todos[i]
		}
	}
	return nil
}

func (t *Todos) Save(file string) error {
	json, err := t.Json()
	if err != nil {
		return err
	}

	if err := files.SaveWithMkdirAll(json, file, 0777); err != nil {
		return err
	}
	return nil
}

func (t *Todos) Add(todo Todo) {
	t.Todos = append(t.Todos, todo)
}

func (t *Todos) Json() ([]byte, error) {
	return json.Marshal(t)
}

func read(file string, todos *Todos) error {
	data, err := files.Read(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return err
	}

	return nil
}
