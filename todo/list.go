package todo

type TodoList struct {
	tasks  []*Task
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		tasks:  nil,
		nextID: 1,
	}
}

func (l *TodoList) Add(name string, isTemporary bool) {
	task := NewTask(l.nextID, name, isTemporary)
	l.tasks = append(l.tasks, task)
	l.nextID++
}
