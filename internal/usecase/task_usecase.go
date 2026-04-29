package usecase

import (
	"todo-cli/internal/domain"
	"todo-cli/internal/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) CreateTask(id int, title string) {
	task := &domain.Task{
		Id:    id,
		Title: title,
		Done:  false,
	}
	uc.repo.Create(task)
}

func (uc *TaskUseCase) ListTasks() []*domain.Task {
	return uc.repo.FindAll()
}

func (uc *TaskUseCase) CompleteTask(id int) {
	task := uc.repo.FindById(id)
	if task != nil {
		task.Done = true
	}
}

func (uc *TaskUseCase) DeleteTask(id int) {
	uc.repo.Delete(id)
}
