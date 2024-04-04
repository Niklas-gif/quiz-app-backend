package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func initMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	db = client.Database("quiz_app")
	collection = db.Collection("quiz_collection")
}

//Quizz operations

func getAllQuizzes() {

}

func getQuiz() {

}

func deleteQuiz() {

}

func addQuiz() {

}

func addQuestion() {

}

func deleteQuestion() {

}

func addAnswer() {

}

func deleteAnswer() {

}
