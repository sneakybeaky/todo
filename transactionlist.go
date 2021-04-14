package todo

import (
	"os"
	"path"
)

type TransactionList struct {
	dl *DiskList
}

func NewTransactionList(dir string) (*TransactionList, error) {

	dl, err := NewDiskList(dir)

	if err != nil {
		return nil, err
	}

	return &TransactionList{dl: dl}, nil
}

func (t *TransactionList) Add(todo Todo) error {

	f, err := os.OpenFile(path.Join(t.dl.dir, "todo_wal.txt"), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	return t.dl.Add(todo)
}

func (t *TransactionList) Items() []Todo {
	return t.dl.Items()
}
