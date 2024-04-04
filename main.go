package main

import (
	"context"
	"log"
	"net/http"
	"quiz-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.InitMongoDB()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/all", func(c *gin.Context) {
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

	})

	router.POST("/example", func(c *gin.Context) {
		database.InsertExampleQuiz()
	})

	router.Run()

	defer database.Client.Disconnect(context.Background())
}
