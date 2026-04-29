package interfaces

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo-cli/internal/usecase"
)

type CLIHandler struct {
	usecase *usecase.TaskUseCase
}

func NewCLIHandler(usecase *usecase.TaskUseCase) *CLIHandler {
	return &CLIHandler{usecase: usecase}
}

func (cli *CLIHandler) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1 - Criar tarefa")
		fmt.Println("2 - Listar tarefas")
		fmt.Println("3 - Concluir tarefa")
		fmt.Println("4 - Deletar tarefa")
		fmt.Println("0 - Sair")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Título da tarefa: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			cli.usecase.CreateTask(len(cli.usecase.ListTasks())+1, title)
			fmt.Println("Tarefa criada com sucesso!")
		case "2":
			tasks := cli.usecase.ListTasks()
			for _, task := range tasks {
				status := "⭕"
				if task.Done {
					status = "✅"
				}
				fmt.Printf("%d. [%s] %s\n", task.Id, status, task.Title)
			}

		case "3":
			fmt.Print("ID da tarefa a concluir: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			cli.usecase.CompleteTask(id)

			fmt.Println("Tarefa concluída com sucesso!")

		case "4":
			fmt.Print("ID da tarefa a deletar: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			cli.usecase.DeleteTask(id)

			fmt.Println("Tarefa deletada com sucesso!")

		case "0":
			return
		}
	}
}
