package todo

type DiskList struct{

}

func NewDiskList() List {
	return &DiskList{}
}

func (dl *DiskList) Add(t Todo) {
}

func (dl *DiskList) Items() []Todo {
	return []Todo{}
}