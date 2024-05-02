package quizservice

import (
	"context"
	"net/http"
	"quiz-app/database"
	"quiz-app/quizmodel"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetQuestion(c *gin.Context) *quizmodel.Question {
	name := c.Param("name")
	index := c.Param("questionindex")
	filter := bson.D{{Key: "name", Value: name}}

	var quiz quizmodel.Quiz
	err := database.Collection.FindOne(context.Background(), filter).Decode(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find quiz"})
		return nil
	}

	i, err := strconv.Atoi(index)
	if err != nil || i < 0 || i >= len(quiz.Questions) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question index"})
		return nil
	}

	question := quiz.Questions[i]
	c.IndentedJSON(http.StatusOK, gin.H{"question": question})
	return &question
}

func AddQuestion(c *gin.Context) {

}

func DeleteQuestion() {

}
