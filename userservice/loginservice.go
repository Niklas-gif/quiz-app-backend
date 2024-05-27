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

	//TODO verfiy ADMIN!
	var user usermodel.User
	var _ = database.DB.Collection("user").FindOne(c, bson.D{{Key: "id", Value: "1"}}).Decode(&user)
	fmt.Printf("The user request value %v", user)

	tokenString, err := utils.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
