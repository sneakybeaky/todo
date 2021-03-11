package todo_test

import (
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

func TestAddToEmptyList(t *testing.T) {
	list := todo.NewMemoryList()
	todo1 := todo.Todo{ Title: "get this test passing" }
	want := []todo.Todo{todo1}
	list.Add(todo1)
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddDuplicateItemIgnored(t *testing.T) {
	list := todo.NewMemoryList()
	todo1 := todo.Todo{ Title: "get this test passing" }
	list.Add(todo1)
	list.Add(todo1)
	want := []todo.Todo{
		todo1,
	}
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestAddDifferentItemIsOkay(t *testing.T) {
	list := todo.NewMemoryList()
	todo1 := todo.Todo{ Title: "get this test passing" }
	list.Add(todo1)
	todo2 := todo.Todo{ Title: "really get this test passing" }
	list.Add(todo2)
	// should be ignored
	want := []todo.Todo{
		todo1,
		todo2,
	}
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
