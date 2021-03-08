package inmemory

import "github.com/sneakybeaky/todo/pkg/todo"

type List struct {
	items []*todo.Todo
}

func NewList() *List {
	return &List{}
}

func (m *List) Add(todo *todo.Todo) {
	m.items = append(m.items, todo)
}

func (m *List) Items() []*todo.Todo {
	return m.items
}
