package userservice

import (
	"quiz-app/database"
	"quiz-app/usermodel"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var ()

func CreateAdmin(c *gin.Context) {
	collection := database.DB.Collection("user")

	var newUser = usermodel.User{
		ID:       1,
		Email:    "fake@mail.com",
		Password: "Password1",
	}

	filter := bson.D{{Key: "id", Value: newUser.ID}}

	var user usermodel.User
	response := collection.FindOne(c, filter).Decode(&user)
	if response != nil {
		println("response%s", response)

	}
	if user.ID == newUser.ID {
		return
	}
	collection.InsertOne(c, newUser)

}
