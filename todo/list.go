package todo

import "fmt"

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

func (l *TodoList) Complete(id int) {
	for _, task := range l.tasks {
		if task.ID == id {
			task.Complete()
			return
		}
	}
}

func (l *TodoList) Print() {
	if len(l.tasks) == 0 {
		fmt.Println("В списке нет задач")
		return
	}
	for _, task := range l.tasks {
		fmt.Println(task)
	}
}

func (l *TodoList) DeleteTask(id int) {
	for i, task := range l.tasks {
		if task.ID == id {
			l.tasks = append(l.tasks[:i], l.tasks[i+1:]...)
			for j, task := range l.tasks {
				task.ID = j + 1
			}
			l.nextID = len(l.tasks) + 1
			return
		}
	}
}
