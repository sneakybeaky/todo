package todo

type Todo struct {
	Title string
}

type List interface {
	Add(Todo)
	Items() []Todo
}
