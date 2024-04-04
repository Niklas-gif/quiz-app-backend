package main

import (
	"context"
	"log"
	"net/http"
	"quiz-app/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
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

	//Example
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/all", func(c *gin.Context) {
		cursor, err := collection.Find(c, bson.D{{}})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var quizzes []bson.M
		if err = cursor.All(c, &quizzes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, quizzes)

	})

	router.Run()

	defer client.Disconnect(context.Background())
}

func initMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	db = client.Database("quiz_app")
	collection = db.Collection("quiz_collection")
}

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
