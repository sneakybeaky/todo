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
	dir := tempDir(t)

	list, err := todo.NewDiskList(dir)

	if err != nil {
		t.Fatal(err)
	}

	todo1 := todo.Todo{Title: "get this test passing"}
	want := []todo.Todo{todo1}
	list.Add(todo1)
	got := list.Items()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestCloseAndOpenDiskList(t *testing.T) {

	dir := tempDir(t)

	todo1 := todo.Todo{Title: "get this test passing"}
	{
		list, err := todo.NewDiskList(dir)
		if err != nil {
			t.Fatal(err)
		}
		list.Add(todo1)
	}

	want := []todo.Todo{todo1}

	list, err := todo.NewDiskList(dir) // should load list from disk
	if err != nil {
		t.Fatal(err)
	}

	got := list.Items()
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
