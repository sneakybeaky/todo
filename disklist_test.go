package todo_test

import (
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

func TestAddToEmptyDiskList(t *testing.T) {
	list := todo.NewDiskList()
	todo1 := todo.Todo{ Title: "get this test passing" }
	want := []todo.Todo{todo1}
	list.Add(todo1)
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCloseAndOpenDiskList(t *testing.T) {
	t.FailNow()
}