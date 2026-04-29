package main

import (
	"todo-cli/internal/interfaces"
	"todo-cli/internal/repository"
	"todo-cli/internal/usecase"
)

func main() {
	repo := repository.NewInMemoryTaskRepository()
	usecase := usecase.NewTaskUseCase(repo)
	cli := interfaces.NewCLIHandler(usecase)

	cli.Start()
}
