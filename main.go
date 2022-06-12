package main

import (
	"mongo-todo-app/app"
	"mongo-todo-app/configs"
	"mongo-todo-app/repository"
	"mongo-todo-app/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appRoute := fiber.New()
	configs.ConnectDB()

	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDB := repository.NewTodoRepositoryDB(dbClient)

	todoHandler := app.TodoHandler{Service: services.NewTodoService(TodoRepositoryDB)}

	appRoute.Post("/api/todo", todoHandler.CreateTodo)
	appRoute.Get("/api/todos", todoHandler.GetAllTodo)
	appRoute.Delete("/api/todo/:id", todoHandler.DeleteTodo)
	appRoute.Listen("localhost:8080")
}
