package services

import (
	"mongo-todo-app/mocks/repository"
	"mongo-todo-app/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var mockRepo *repository.MockTodoRepository
var service TodoService

var FakeData = []models.Todo{
	{Id: primitive.NewObjectID(), Title: "Mock Title1", Content: "Mock Content1"},
	{Id: primitive.NewObjectID(), Title: "Mock Title2", Content: "Mock Content2"},
	{Id: primitive.NewObjectID(), Title: "Mock Title3", Content: "Mock Content3"},
}

func setup(t *testing.T) func() {
	ct := gomock.NewController(t)
	defer ct.Finish()

	mockRepo = repository.NewMockTodoRepository(ct)
	service = NewTodoService(mockRepo)

	return func() {
		service = nil
		defer ct.Finish()
	}
}

func TestDefaultTodoService_TodoGetAll(t *testing.T) {
	td := setup(t)
	defer td()

	mockRepo.EXPECT().GetAll().Return(FakeData, nil)

	result, err := service.TodoGetAll()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, result)
}
