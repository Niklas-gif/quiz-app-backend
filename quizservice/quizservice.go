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

	_, err := database.Collection.InsertOne(c, quizmodel.SampleQuiz)
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

func GetAnswer(c *gin.Context) {
	question := GetQuestion(c)
	index := c.Param("answerindex")
	filter := bson.D{{Key: "answers", Value: "answers"}}

	var answer quizmodel.Answer
	err := database.Collection.FindOne(context.Background(), filter).Decode(&question)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find Question"})
		return
	}

	i, err := strconv.Atoi(index)
	if err != nil || i < 0 || i >= len(question.Answers) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid answer index"})
		return
	}

	answer = question.Answers[i]
	c.IndentedJSON(http.StatusOK, gin.H{"answer": answer})
}

func AddAnswer(c *gin.Context) {

}

func DeleteAnswer() {

}
