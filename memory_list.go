package todo

type MemoryList struct {
	items []Todo
}

func NewMemoryList() *MemoryList {
	return &MemoryList{}
}

func (m *MemoryList) Add(todo Todo) error {

	if m.contains(todo) {
		return nil
	}

	m.items = append(m.items, todo)

	return nil
}

//contains checks if a todo is already in the list
func (m *MemoryList) contains(todo Todo) bool {
	for _, t := range m.items {
		if t == todo {
			return true
		}
	}
	return false
}

func (m *MemoryList) Items() []Todo {
	return m.items
}
