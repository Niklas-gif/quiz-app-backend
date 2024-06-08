package main

import (
	"context"
	"log"
	"quiz-app/database"
	"quiz-app/middleware"
	"quiz-app/quizservice"
	"quiz-app/userservice"

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
	router := gin.Default()
	//router.Run("0.0.0.0:3030")

	router.Use(middleware.ConfigureCORS)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		userservice.Login(c)
	})

	router.GET("/quizzes", func(c *gin.Context) {
		quizservice.GetAllQuizzes(c)

	})

	router.GET("/quizzes/:name", func(c *gin.Context) {
		quizservice.GetQuizByName(c)
	})

	router.GET("/quizzes/:name/:questionindex", func(c *gin.Context) {
		quizservice.GetQuestion(c)
	})

	router.GET("/quizzes/:name/:questionindex/:answerindex", func(c *gin.Context) {
		quizservice.GetAnswer(c)
	})

	router.PUT("/quizzes/:name", func(c *gin.Context) {
		quizservice.AddQuestion(c)
	})

	router.POST("/example", func(c *gin.Context) {
		quizservice.InsertExampleQuiz(c)
	})

	router.PUT("/add", middleware.AuthenticationMiddleware(), func(c *gin.Context) {
		quizservice.AddQuiz(c)
	})

	router.PUT("/update", middleware.AuthenticationMiddleware(), func(c *gin.Context) {
		quizservice.UpdateQuiz(c)
	})

	router.DELETE("/delete", middleware.AuthenticationMiddleware(), func(c *gin.Context) {
		quizservice.DeleteQuiz(c)
	})

	router.Run("0.0.0.0:3030")

	defer database.Client.Disconnect(context.Background())
}
