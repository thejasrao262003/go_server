package repositories

import (
	"context"
	"time"

	"go_server/internal/database"
	"go_server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoTaskRepository struct {
	dbName         string
	collectionName string
}

func NewMongoTaskRepository(dbName, collectionName string) *MongoTaskRepository {
	return &MongoTaskRepository{
		dbName:         dbName,
		collectionName: collectionName,
	}
}

func (r *MongoTaskRepository) Create(task models.Task) (models.Task, error){
	collection := database.Client.Database(r.dbName).Collection(r.collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := collection.InsertOne(ctx, task)
	return task, err
}

func (r *MongoTaskRepository) GetAll() ([]models.Task, error){
	collection := database.Client.Database(r.dbName).Collection(r.collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err!=nil{
		return nil, err
	}

	defer cursor.Close(ctx)

	var tasks []models.Task

	for cursor.Next(ctx){
		var task models.Task
		
		err := cursor.Decode(&task)
		if err!=nil{
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, err
}

func (r *MongoTaskRepository) Update(id primitive.ObjectID, updateData bson.M) error {
	collection := database.Client.Database(r.dbName).Collection(r.collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	update := bson.M{
		"$set": updateData,
	}

	_, err := collection.UpdateOne(ctx, bson.M{
		"_id": id,
	}, update)

	return err
}

func (r *MongoTaskRepository) Delete(id primitive.ObjectID) error {
	collection := database.Client.Database(r.dbName).Collection(r.collectionName)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}