package app

import (
	"mongo-todo-app/models"
	"mongo-todo-app/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.TodoInsert(todo)

	if err != nil || !result.Status {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {

	todos, err := h.Service.TodoGetAll()

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
	if len(todos) == 0 {
		return c.Status(http.StatusOK).JSON(fiber.Map{"todos": c.JSON(todos), "status": "List is empty"})
	}

	return c.Status(http.StatusOK).JSON(todos)
}

func (h TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	query := c.Params("id")
	id, _ := primitive.ObjectIDFromHex(query)
	result, err := h.Service.TodoDelete(id)
	if err != nil || !result.Status {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": c.JSON(result.Status), "error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}
