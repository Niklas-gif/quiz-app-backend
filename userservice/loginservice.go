package userservice

import (
	"fmt"
	"net/http"
	"quiz-app/database"
	"quiz-app/usermodel"
	"quiz-app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {

	var user usermodel.User
	err := c.BindJSON(&user)

	if err != nil {
		fmt.Print(err)
		return
	}
	//TODO
	var admin usermodel.User
	err = database.DB.Collection("user").FindOne(c, bson.D{{Key: "email", Value: "fake@mail.com"}}).Decode(&admin)

	if err != nil {
		fmt.Printf("The user request value %v", admin)
		return
	}

	if user.Email == admin.Email {
		tokenString, err := utils.CreateToken(admin.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{})
	}
}
