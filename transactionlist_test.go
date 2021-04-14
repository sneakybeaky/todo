package todo_test

import (
	"os"
	"path"
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

func TestTListPersists(t *testing.T) {

	td := t.TempDir()
	todo1 := todo.Todo{Title: "get this test passing"}

	var originalList *todo.TransactionList
	originalList, err := todo.NewTransactionList(td)

	if err != nil {
		t.Fatal(err)
	}
	err = originalList.Add(todo1)

	if err != nil {
		t.Fatal(err)
	}

	var afterList *todo.TransactionList
	afterList, err = todo.NewTransactionList(td)

	if err != nil {
		t.Fatal(err)
	}
	got := afterList.Items()

	want := []todo.Todo{todo1}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}

func TestAddingToTListCreatesTransactionLog(t *testing.T) {

	td := t.TempDir()
	todo1 := todo.Todo{Title: "get this test passing"}

	var originalList *todo.TransactionList
	originalList, err := todo.NewTransactionList(td)

	if err != nil {
		t.Fatal(err)
	}
	err = originalList.Add(todo1)

	if err != nil {
		t.Fatal(err)
	}

	wantLogFile := path.Join(td, "todo_wal.txt")
	_, err = os.Stat(wantLogFile)

	if err != nil {
		t.Fatalf("%v", err)
	}

	t.Cleanup(func() {
		err := os.Remove(wantLogFile)
		if err != nil {
			t.Fatalf("Unable to clean up transaction log file %v", err)
		}
	})

}
