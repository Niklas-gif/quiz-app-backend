package database

import (
	"context"
	"log"
	"net/http"
	"quiz-app/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
)

func InitMongoDB() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	DB = client.Database("quiz_app")
	Collection = DB.Collection("quiz_collection")
}

func InsertExampleQuiz(c *gin.Context) {
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
	_, err := Collection.InsertOne(c, sampleQuiz)
	if err != nil {
		log.Fatal("Failed to insert sample data into MongoDB:", err)
	}
}

//Quiz operations

func GetAllQuizzes(c *gin.Context) {
	cursor, err := Collection.Find(c, bson.D{{}})
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
