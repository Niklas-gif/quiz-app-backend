package database

import (
	"context"
	"log"
	"quiz-app/quizmodel"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
)

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	DB = client.Database("quiz_app")
	Collection = DB.Collection("quiz_collection")

	for _, quiz := range quizmodel.ExampleQuizzes {
		filter := bson.M{"name": quiz.QuizName}
		var existingQuiz quizmodel.Quiz
		err := Collection.FindOne(context.Background(), filter).Decode(&existingQuiz)
		if err == mongo.ErrNoDocuments {
			_, err := Collection.InsertOne(context.Background(), quiz)
			if err != nil {
				log.Fatalf("Failed to insert data: %v", err)
			}
			log.Printf("Inserted quiz: %s\n", quiz.QuizName)
		} else if err != nil {
			log.Fatalf("Failed to check existing data: %v", err)
		} else {
			log.Printf("Quiz already exists: %s\n", quiz.QuizName)
		}
	}

	log.Println("Initial data inserted successfully")

	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"name": 1,
		},
		Options: options.Index().SetUnique(true),
	}

	_, err = Collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}
}
