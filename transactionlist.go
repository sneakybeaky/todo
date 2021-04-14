package todo

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
	return t.dl.Add(todo)
}

func (t *TransactionList) Items() []Todo {
	return t.dl.Items()
}
