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

func (dl *DiskList) store() error {
	contents, err := json.MarshalIndent(dl.Items(), "", " ")

	if err != nil {
		return fmt.Errorf("unable to prepare list to store to disk : %v", err)
	}
	if err = ioutil.WriteFile(path.Join(dl.dir, "todo.json"), contents, 0644); err != nil {
		return fmt.Errorf("unable to store list to disk : %v", err)
	}

	return nil
}

func (dl *DiskList) restore() error {
	if _, err := os.Stat(path.Join(dl.dir, "todo.json")); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("unable to restore list from disk : %v", err)
	}
	return dl.read()
}

func (dl *DiskList) read() error {
	content, err := ioutil.ReadFile(path.Join(dl.dir, "todo.json"))
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
