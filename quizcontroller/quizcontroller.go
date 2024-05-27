package quizcontroller

import (
	"quiz-app/quizservice"

	"github.com/gin-gonic/gin"
)

//TODO define endpoints

func InsertExampleQuiz(c *gin.Context) {
	quizservice.InsertExampleQuiz(c)
}

func InsertQuiz(c *gin.Context) {

}
