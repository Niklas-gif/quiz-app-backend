package main

import (
	"context"
	"log"
	"quiz-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
		database.GetAllQuizzes(c)

	})

	router.POST("/example", func(c *gin.Context) {
		database.InsertExampleQuiz(c)
	})

	router.Run()

	defer database.Client.Disconnect(context.Background())
}
