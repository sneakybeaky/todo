package todo_test

import (
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

func TestAppendItemToList(t *testing.T) {

	list := mustDiskList(t, t.TempDir())

	todo1 := todo.Todo{Title: "get this test passing"}
	want := []todo.Todo{todo1}
	_ = list.Add(todo1)
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestListPersists(t *testing.T) {

	td := t.TempDir()
	todo1 := todo.Todo{Title: "get this test passing"}

	{
		list := mustDiskList(t, td)
		_ = list.Add(todo1)
	}

	list := mustDiskList(t, td) // should load list from disk

	got := list.Items()
	want := []todo.Todo{todo1}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func mustDiskList(t *testing.T, dir string) *todo.DiskList {

	t.Helper()
	list, err := todo.NewDiskList(dir)

	if err != nil {
		t.Fatal(err)
	}

	return list
}
