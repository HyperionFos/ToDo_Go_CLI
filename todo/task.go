package todo

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Name        string
	Done        bool
	CreatedAt   time.Time
	IsTemporary bool
	Deadline    *time.Time
}

func NewTask(id int, name string, isTemporary bool) *Task {
	return &Task{
		ID:          id,
		Name:        name,
		Done:        false,
		CreatedAt:   time.Now(),
		IsTemporary: isTemporary,
		Deadline:    nil,
	}
}

func (t *Task) SetDeadline(deadline time.Time) {
	if t.IsTemporary {
		t.Deadline = &deadline
	}
}

func (t *Task) Complete() {
	t.Done = true
}

func (t *Task) String() string {
	status := "[ ]"
	if t.Done {
		status = "[x]"
	}
	return fmt.Sprintf("%s #%d: %s", status, t.ID, t.Name)
}
