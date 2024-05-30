package quizservice

import (
	"fmt"
	"log"
	"net/http"
	"quiz-app/database"
	"quiz-app/quizmodel"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func UpdateQuiz(c *gin.Context) {
	var quiz quizmodel.Quiz
	err := c.BindJSON(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	objectID, err := primitive.ObjectIDFromHex(quiz.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid ID"})
		return
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":        quiz.QuizName,
		"description": quiz.QuizDescription,
		"questions":   quiz.Questions,
	}}
	fmt.Println("Updating quiz with ID:", quiz.ID)

	if _, err := database.Collection.UpdateOne(c, filter, update); err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func AddQuiz(c *gin.Context) {
	var quiz quizmodel.Quiz
	err := c.BindJSON(&quiz)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if _, err := database.Collection.InsertOne(c, quiz); err != nil {
		//TODO this might be bad
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
