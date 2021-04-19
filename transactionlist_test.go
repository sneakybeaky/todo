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

func SequenceFrom(start int) func(*todo.TransactionList) error {

	s := start
	return func(list *todo.TransactionList) error {
		list.Sequence = func() int {
			s++
			return s
		}
		return nil
	}
}

func TestAddingToTListCreatesTransactionLog(t *testing.T) {

	td := t.TempDir()
	todo1 := todo.Todo{Title: "get this test passing"}

	var originalList *todo.TransactionList
	originalList, err := todo.NewTransactionList(td, SequenceFrom(0))

	if err != nil {
		t.Fatal(err)
	}
	err = originalList.Add(todo1)

	if err != nil {
		t.Fatal(err)
	}

	wantLogFile := path.Join(td, "todo_wal.txt")
	gotData, err := os.ReadFile(wantLogFile)

	if err != nil {
		t.Fatalf("%v", err)
	}

	wantData := []byte("1\t1\tget this test passing\n")
	if !cmp.Equal(wantData, gotData) {
		t.Fatalf(cmp.Diff(string(wantData), string(gotData)))
	}

}
