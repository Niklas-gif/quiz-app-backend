package quizservice

import (
	"context"
	"log"
	"net/http"
	"quiz-app/database"
	"quiz-app/quizmodel"
	"strconv"

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
	c.IndentedJSON(http.StatusOK, quizzes)
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

	c.IndentedJSON(http.StatusOK, quizzes)
}

func DeleteQuiz() {

}

func AddQuiz(c *gin.Context) {
	var quiz quizmodel.Quiz
	err := c.BindJSON(&quiz)
	if err != nil {
		return
	}
	database.Collection.InsertOne(c, quiz)
}

func GetQuestion(c *gin.Context) {
	name := c.Param("name")
	index := c.Param("index")
	filter := bson.D{{Key: "name", Value: name}}

	var quiz quizmodel.Quiz
	err := database.Collection.FindOne(context.Background(), filter).Decode(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find quiz"})
		return
	}

	i, err := strconv.Atoi(index)
	if err != nil || i < 0 || i >= len(quiz.Questions) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question index"})
		return
	}

	question := quiz.Questions[i]
	c.IndentedJSON(http.StatusOK, gin.H{"question": question})
}

func AddQuestion(c *gin.Context) {

}

func DeleteQuestion() {

}

func AddAnswer(c *gin.Context) {
	//TODO
}

func DeleteAnswer() {

}
