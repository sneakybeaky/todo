package todo

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

// Sequencer is something that generates monotonically increasing values
type Sequencer func() int

const LogFilename = "todo_wal.txt"

type TransactionList struct {
	list     List
	Sequence Sequencer
	LogFile  string
}

type TransactionListOption func(list *TransactionList) error

func NewTransactionList(dir string, opts ...TransactionListOption) (*TransactionList, error) {

	list := NewMemoryList()

	tl := &TransactionList{list: list, LogFile: path.Join(dir, LogFilename)}

	for _, opt := range opts {
		if err := opt(tl); err != nil {
			return nil, err
		}
	}

	if err := tl.replay(); err != nil {
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

	return t.list.Add(todo)
}

func (t *TransactionList) Items() []Todo {
	return t.list.Items()
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

		t.list.Add(todo)

	}

	return nil
}

func (t *TransactionList) Snapshot() error {
	return os.Truncate(t.LogFile, 0)
}
