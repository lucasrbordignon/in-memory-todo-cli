package domain

type Task struct {
	Id    int
	Title string
	Done  bool
}

func (t *Task) markAsDone() {
	t.Done = true
}
