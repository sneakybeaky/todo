package todo_test

import (
	"testing"
	"todo"
)

func TestTodo(t *testing.T) {
	_ = todo.Todo{
		Title: "get started",
	}
}
