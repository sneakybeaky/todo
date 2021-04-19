package todo

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

// Sequencer is something that generates monotonically increasing values
type Sequencer func() int

type TransactionList struct {
	dl       *DiskList
	Sequence Sequencer
	LogFile  string
}

type TransactionListOption func(list *TransactionList) error

func NewTransactionList(dir string, opts ...TransactionListOption) (*TransactionList, error) {

	dl, err := NewDiskList(dir)

	if err != nil {
		return nil, err
	}

	tl := &TransactionList{dl: dl, LogFile: path.Join(dir, "todo_wal.txt")}

	for _, opt := range opts {
		if err = opt(tl); err != nil {
			return nil, err
		}

	}

	if err = tl.replay(); err != nil {
		return nil, err
	}

	return tl, nil
}

func (t *TransactionList) Add(todo Todo) error {

	f, err := os.OpenFile(t.LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)

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

// replay restores items from the log file, if it exists
func (t *TransactionList) replay() error {

	f, err := os.OpenFile(path.Join(t.LogFile), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)

	if err != nil {
		return err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		var seq, eventtype int
		var todo Todo
		if _, err := fmt.Sscanf(line, "%d\t%d\t%s", &seq, &eventtype, &todo.Title); err != nil {
			return fmt.Errorf("input parse error: %w", err)
		}

		t.dl.Add(todo)

	}

	return nil
}
