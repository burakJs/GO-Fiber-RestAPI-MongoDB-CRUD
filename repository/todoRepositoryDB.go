package repository

import (
	"context"
	"errors"
	"log"
	"mongo-todo-app/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockgen -destination=../mocks/repository/mockTodoRepository.go -package=repository mongo-todo-app/repository TodoRepository

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()
	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		_ = errors.New("Insert failed err: " + err.Error())
		return false, err
	}
	return true, nil
}

func (t TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (t TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		log.Fatalln(err)
		return false, err
	}

	return true, nil
}

func NewTodoRepositoryDB(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
