package inmemory

import "github.com/sneakybeaky/todo/pkg/todo"

type List struct {
	items []*todo.Todo
}

func NewList() *List {
	return &List{}
}

func (m *List) Add(todo *todo.Todo) {

	// Check to see if we already have the todo
	for _, t := range m.items {
		if t == todo {
			return
		}
	}

	m.items = append(m.items, todo)
}

func (m *List) Items() []*todo.Todo {
	return m.items
}
