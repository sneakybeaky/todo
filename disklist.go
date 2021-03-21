package todo

type DiskList struct {
	list List
}

func NewDiskList() List {
	return &DiskList{list: NewMemoryList()}
}

func (dl *DiskList) Add(t Todo) {
	dl.list.Add(t)
}

func (dl *DiskList) Items() []Todo {
	return dl.list.Items()
}
