package todo_test

import (
	"os"
	"path"
	"testing"

	"todo"

	"github.com/google/go-cmp/cmp"
)

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

func TestRestoreFromTransactionLog(t *testing.T) {
	const input = "testdata/singleadd.txt"

	td := t.TempDir()
	list, err := todo.NewTransactionList(td, TransactionLog(input))
	if err != nil {
		t.Fatal(err)
	}

	got := list.Items()
	want := []todo.Todo{{Title: "blahblahblah"}}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TransactionLog(path string) func(*todo.TransactionList) error {

	return func(list *todo.TransactionList) error {
		list.LogFile = path
		return nil
	}
}

func SequenceFrom(start int) func(*todo.TransactionList) error {

	return func(list *todo.TransactionList) error {
		list.Sequence = func() int {
			start++
			return start
		}
		return nil
	}
}