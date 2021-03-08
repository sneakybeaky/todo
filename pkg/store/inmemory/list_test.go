package inmemory_test

import (
	"github.com/sneakybeaky/todo/pkg/store/inmemory"
	"github.com/sneakybeaky/todo/pkg/todo"
	"testing"
)

func TestAdd(t *testing.T) {

	// given a todo item
	want := todo.Todo{Title: "get this test passing"}

	// when I add it to a new list
	l := inmemory.NewList()
	l.Add(&want)

	// then the list should have just that todo
	f := l.Items()

	if len(f) != 1 {
		t.Errorf("The list should have 1 item, not %d", len(f))
	}

	got := f[0]

	if got != &want {
		t.Errorf("wanted %+v, got %+v", want, *got)
	}

}
