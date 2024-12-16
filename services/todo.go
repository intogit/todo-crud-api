package services

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	Task      string    `json:"task,omitempty" bson:"task, omitempty"`
	Status    string    `json:"status,omitempty" bson:"status, omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at, omitempty"`
}

var client *mongo.Client

func New(mongo *mongo.Client) Todo {
	client = mongo
	return Todo{}
}

func returnCollectionPointer(db string, coll string) *mongo.Collection {
	return client.Database(db).Collection(coll)
}
func InsertTodo(newTodo Todo) error {
	coll := returnCollectionPointer(os.Getenv("DB_NAME"), "todos")
	newTodoObject := Todo{
		Task:      newTodo.Task,
		Status:    newTodo.Status,
		CreatedAt: time.Now(),
	}
	_, err := coll.InsertOne(context.TODO(), newTodoObject)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func GetAllTodo() ([]Todo, error) {
	coll := returnCollectionPointer(os.Getenv("DB_NAME"), "todos")

	filter := bson.D{}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var results []Todo
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	return results, err
}

func GetTodoById(id string) (Todo, error) {
	coll := returnCollectionPointer(os.Getenv("DB_NAME"), "todos")

	mongoID, err := primitive.ObjectIDFromHex(id)
	log.Println(id)
	log.Println(mongoID)
	if err != nil {
		return Todo{}, err
	}

	filter := bson.D{{"_id", mongoID}}
	var todo Todo
	err = coll.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		log.Fatal(err)
		return Todo{}, err
	}
	return todo, nil
}
