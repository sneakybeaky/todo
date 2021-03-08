package todo

type Todo struct {
	Title string
}

type List interface {

	// Add adds the todo to the list
	Add(todo *Todo)

	//Items returns all the todos in the list
	Items() []*Todo
}
