package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

type DiskList struct {
	list List
	dir  string
}

func NewDiskList(dir string) (*DiskList, error) {

	dl := &DiskList{dir: dir, list: NewMemoryList()}
	if err := dl.restore(); err != nil {
		return nil, err
	}
	return dl, nil
}

func (dl *DiskList) Add(t Todo) {
	dl.list.Add(t)
	_ = dl.store()
}

func (dl *DiskList) Items() []Todo {
	return dl.list.Items()
}

func (dl *DiskList) filename() string {
	return path.Join(dl.dir, "todo.json")
}

// store writes the contents of the list to disk
func (dl *DiskList) store() error {
	contents, err := json.MarshalIndent(dl.Items(), "", " ")

	if err != nil {
		return fmt.Errorf("unable to prepare list to store to disk : %v", err)
	}
	if err = ioutil.WriteFile(dl.filename(), contents, 0644); err != nil {
		return fmt.Errorf("unable to store list to disk : %v", err)
	}

	return nil
}

// restore reads any previously stored list from disk
func (dl *DiskList) restore() error {
	if _, err := os.Stat(dl.filename()); err != nil {
		if os.IsNotExist(err) {
			return nil // assume this means this is a brand new list - no stored file
		}
		return fmt.Errorf("unable to restore list from disk : %v", err)
	}
	return dl.read()
}

func (dl *DiskList) read() error {
	content, err := ioutil.ReadFile(dl.filename())
	if err != nil {
		return fmt.Errorf("unable to read list from disk : %v", err)
	}

	todos := make([]Todo, 0)
	err = json.Unmarshal(content, &todos)
	if err != nil {
		return fmt.Errorf("unable to unmarshal list : %v", err)
	}

	for _, todo := range todos {
		dl.list.Add(todo)
	}

	return nil
}
