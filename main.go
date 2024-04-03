package main

import (
	"context"
	"fmt"
	"log"
	"quiz-app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	initMongoDB()
	//var answer = models.Answer{Description: "Test", IsCorrect: true}
	fmt.Printf("Hello, go!\n")

	//Example
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func initMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	db = client.Database("your_database_name")
	collection = db.Collection("your_collection_name")
	insertTest()
	fmt.Println(client)
}

/*func insertTest() {
	var _, err = collection.InsertOne(context.Background(), bson.M{"message": "This is a sample message from MongoDB"})
	if err != nil {
		log.Fatal("Failed to insert sample data into MongoDB:", err)
	}
}*/

func insertTest() {
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
				{Description: "Banana", IsCorrect: true},
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
