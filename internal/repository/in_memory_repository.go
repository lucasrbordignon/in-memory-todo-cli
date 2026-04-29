package repository

import "todo-cli/internal/domain"

type InMemoryTaskRepository struct {
	tasks []*domain.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: []*domain.Task{},
	}
}

func (r *InMemoryTaskRepository) Create(task *domain.Task) {
	r.tasks = append(r.tasks, task)
}

func (r *InMemoryTaskRepository) FindAll() []*domain.Task {
	return r.tasks
}

func (r *InMemoryTaskRepository) FindById(id int) *domain.Task {
	for _, task := range r.tasks {
		if task.Id == id {
			return task
		}
	}
	return nil
}

func (r *InMemoryTaskRepository) Delete(id int) {
	for i, task := range r.tasks {
		if task.Id == id {
			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)
			return
		}
	}
}
