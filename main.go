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
	//database.InitMongoDB()
	router := gin.Default()

	router.Use(middleware.ConfigureCORS)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/admin", func(c *gin.Context) {
		userservice.CreateAdmin(c)
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

	router.POST("/add", func(c *gin.Context) {
		quizservice.AddQuiz(c)
	})

	/*TODO: router.POST("/quiz/:quizName/question/:questionIndex/answer", func(c *gin.Context) {
		quizservice.AddAnswer(c)
	})*/

	router.DELETE("/quizzes/:name/:index", func(c *gin.Context) {

	})

	router.Run()

	defer database.Client.Disconnect(context.Background())
}
