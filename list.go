package todo

type List struct {
	items []Todo
}

func NewList() *List {
	return &List{}
}

func (m *List) Add(todo Todo) {

	if m.contains(todo) {
		return
	}

	m.items = append(m.items, todo)
}

//contains checks if a todo is already in the list
func (m *List) contains(todo Todo) bool {
	for _, t := range m.items {
		if t == todo {
			return true
		}
	}
	return false
}

func (m *List) Items() []Todo {
	return m.items
}
