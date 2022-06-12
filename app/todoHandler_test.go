package app

import (
	services "mongo-todo-app/mocks/service"
	"mongo-todo-app/models"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var td TodoHandler
var mockService *services.MockTodoService

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	mockService = services.NewMockTodoService(ct)
	td = TodoHandler{Service: mockService}

	return func() {
		defer ct.Finish()
	}
}

func TestTodoHandler_GetAllTodo(t *testing.T) {

	tdh := setup(t)
	defer tdh()

	router := fiber.New()
	router.Get("/api/todos", td.GetAllTodo)

	var FakeDataForHandler = []models.Todo{
		{Id: primitive.NewObjectID(), Title: "Mock Title1", Content: "Mock Content1"},
		{Id: primitive.NewObjectID(), Title: "Mock Title2", Content: "Mock Content2"},
		{Id: primitive.NewObjectID(), Title: "Mock Title3", Content: "Mock Content3"},
	}

	mockService.EXPECT().TodoGetAll().Return(FakeDataForHandler, nil)

	req := httptest.NewRequest("GET", "/api/todos", nil)

	res, _ := router.Test(req, 1)

	assert.Equal(t, 200, res.StatusCode)
}
