package quizservice

import (
	"log"
	"net/http"
	"quiz-app/database"
	"quiz-app/quizmodel"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertExampleQuiz(c *gin.Context) {
	// Sample quiz data
	sampleQuiz := quizmodel.Quiz{
		QuizName:        "Quiz2",
		QuizDescription: "Sample Quiz",
		Questions: []quizmodel.Question{{
			Description:      "Was ist die Hauptstadt von Deutschland?",
			IsMultipleChoice: true,
			Answers: []quizmodel.Answer{
				{Description: "Paris", IsCorrect: false},
				{Description: "London", IsCorrect: false},
				{Description: "Berlin", IsCorrect: true},
			},
		}, {
			Description:      "Was ist 2 + 2?",
			IsMultipleChoice: false,
			Answers: []quizmodel.Answer{
				{Description: "4", IsCorrect: true},
				{Description: "42", IsCorrect: false},
				{Description: "Banana", IsCorrect: false},
			},
		},
		},
	}
	// Inserting sample quiz data into MongoDB
	_, err := database.Collection.InsertOne(c, sampleQuiz)
	if err != nil {
		log.Fatal("Failed to insert sample data into MongoDB:", err)
	}
}

//Quiz operations

func GetAllQuizzes(c *gin.Context) {
	cursor, err := database.Collection.Find(c, bson.D{{}})
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

// TODO: Ignore upper and lower case!!!
func GetQuizByName(c *gin.Context) {
	name := c.Param("name")

	cursor, err := database.Collection.Find(c, bson.M{"name": name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var quizzes []bson.M
	if err := cursor.All(c, &quizzes); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quizzes)
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
