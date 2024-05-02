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

func GetAnswer(c *gin.Context) {
	name := c.Param("name")
	indexquestion := c.Param("questionindex")
	indexanswer := c.Param("answerindex")
	filter := bson.D{{Key: "name", Value: name}}

	var quiz quizmodel.Quiz
	err := database.Collection.FindOne(context.Background(), filter).Decode(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find quiz"})
		return
	}

	i, err := strconv.Atoi(indexquestion)
	if err != nil || i < 0 || i >= len(quiz.Questions) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question index"})
		return
	}

	question := quiz.Questions[i]

	j, err := strconv.Atoi(indexanswer)
	if err != nil || j < 0 || j >= len(question.Answers) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer index"})
		return
	}

	answer := question.Answers[j]
	c.IndentedJSON(http.StatusOK, gin.H{"answer": answer})
}

func AddAnswer(c *gin.Context) {

}

func DeleteAnswer() {

}
