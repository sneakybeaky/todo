package todo

import (
	"fmt"
	"os"
	"path"
)

// Sequencer is something that generates monotonically increasing values
type Sequencer func() int

type TransactionList struct {
	dl       *DiskList
	Sequence Sequencer
}

type TransactionListOption func(list *TransactionList) error

func NewTransactionList(dir string, opts ...TransactionListOption) (*TransactionList, error) {

	dl, err := NewDiskList(dir)

	if err != nil {
		return nil, err
	}

	tl := &TransactionList{dl: dl}

	for _, opt := range opts {
		if err = opt(tl); err != nil {
			return nil, err
		}

	}

	return tl, nil
}

func (t *TransactionList) Add(todo Todo) error {

	f, err := os.OpenFile(path.Join(t.dl.dir, "todo_wal.txt"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = fmt.Fprintf( // Write the event to the log
		f,
		"%d\t%d\t%s\n",
		t.Sequence(), EventAdd, todo.Title)

	if err != nil {
		return err
	}

	return t.dl.Add(todo)
}

func (t *TransactionList) Items() []Todo {
	return t.dl.Items()
}
