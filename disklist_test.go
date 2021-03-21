package todo_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

func TestAddToEmptyDiskList(t *testing.T) {

	list := mustDiskList(t, tempDir(t))

	todo1 := todo.Todo{Title: "get this test passing"}
	want := []todo.Todo{todo1}
	_ = list.Add(todo1)
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCloseAndOpenDiskList(t *testing.T) {

	td := tempDir(t)
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

func tempDir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "todotest")
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Log(fmt.Errorf("unable to delete temp dir %w", err))
		}
	})

	return dir
}

func mustDiskList(t *testing.T, dir string) *todo.DiskList {

	list, err := todo.NewDiskList(dir)

	if err != nil {
		t.Fatal(err)
	}

	return list
}
