package services

import (
	"errors"
	"time"

	"go_server/internal/models"
	"go_server/internal/repositories"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskService struct {
	repo repositories.TaskRepository
}


func NewTaskService(repo repositories.TaskRepository) *TaskService {
	return &TaskService {
		repo: repo,
	}
}

func (s *TaskService) CreateTask(title string) (models.Task, error){
	
	if title == "" {
		return models.Task{}, errors.New("Title is empty")
	}

	task := models.Task{
		ID: primitive.NewObjectID(),
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	return s.repo.Create(task)
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) UpdateTask(id string, updateData bson.M) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task id")
	}

	return s.repo.Update(objectID, updateData)
}

func (s *TaskService) DeleteTask(id string) error {

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task id")
	}

	return s.repo.Delete(objectID)
}
