# TODO CLI (Go)

Aplicação simples de linha de comando para gerenciamento de tarefas, desenvolvida em Go com foco em conceitos de arquitetura limpa e POO adaptada ao Go.

## 🚀 Funcionalidades

- Criar tarefa
- Listar tarefas
- Marcar como concluída
- Deletar tarefa

## 🧠 Conceitos aplicados

- Structs como entidades
- Interfaces (Repository Pattern)
- Injeção de dependência
- Separação por camadas (Domain, UseCase, Repository, Interface)

## 📁 Estrutura

cmd/
internal/
├── domain
├── usecase
├── repository
└── interface

## ▶️ Como rodar

```bash
go run ./cmd
```

## 🛠️ Build
```bash
go build -o todo
```
./todo