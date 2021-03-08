package inmemory

import "github.com/sneakybeaky/todo/pkg/todo"

type List struct {
	items []*todo.Todo
}

func NewList() *List {
	return &List{}
}

func (m *List) Add(todo *todo.Todo) {

	if m.contains(todo) {
		return
	}

	m.items = append(m.items, todo)
}

//contains checks if a todo is already in the list
func (m *List) contains(todo *todo.Todo) bool {
	for _, t := range m.items {
		if t == todo {
			return true
		}
	}
	return false
}

func (m *List) Items() []*todo.Todo {
	return m.items
}
