package quizservice

import (
	"net/http"
	"quiz-app/database"
	"quiz-app/quizmodel"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func init() {
	collection = database.Collection
}

func InsertExampleQuiz(c *gin.Context) error {

	_, err := collection.InsertOne(c, quizmodel.SampleQuiz)
	return err
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

func InsertQuiz(c *gin.Context) {
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
