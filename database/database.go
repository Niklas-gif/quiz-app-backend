package database

import (
	"context"
	"log"
	"quiz-app/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func InitMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	db = client.Database("quiz_app")
	collection = db.Collection("quiz_collection")
}

func InsertExampleQuiz() {
	// Sample quiz data
	sampleQuiz := models.Quiz{
		QuizDescription: "Sample Quiz",
		Questions: []models.Question{{
			Description:      "Was ist die Hauptstadt von Deutschland?",
			IsMultipleChoice: true,
			Answers: []models.Answer{
				{Description: "Paris", IsCorrect: false},
				{Description: "London", IsCorrect: false},
				{Description: "Berlin", IsCorrect: true},
			},
		}, {
			Description:      "Was ist 2 + 2?",
			IsMultipleChoice: false,
			Answers: []models.Answer{
				{Description: "4", IsCorrect: true},
				{Description: "42", IsCorrect: false},
				{Description: "Banana", IsCorrect: false},
			},
		},
		},
	}
	// Inserting sample quiz data into MongoDB
	_, err := collection.InsertOne(context.Background(), sampleQuiz)
	if err != nil {
		log.Fatal("Failed to insert sample data into MongoDB:", err)
	}
}

//Quiz operations

func GetAllQuizzes() {

}

func GetQuiz() {

}

func DeleteQuiz() {

}

func AddQuiz() {

}

func AddQuestion() {

}

func DeleteQuestion() {

}

func AddAnswer() {

}

func DeleteAnswer() {

}
