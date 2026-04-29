package repository

import "todo-cli/internal/domain"

type TaskRepository interface {
	Create(task *domain.Task)
	FindAll() []*domain.Task
	FindById(id int) *domain.Task
	Delete(id int)
}
