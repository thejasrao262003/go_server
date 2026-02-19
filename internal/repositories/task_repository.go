package repositories

import (
	"go_server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskRepository interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	Update(id primitive.ObjectID, updateData bson.M) error
	Delete(id primitive.ObjectID) error
}