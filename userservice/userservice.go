package userservice

import (
	"quiz-app/database"
	"quiz-app/usermodel"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var (
	collection = database.DB.Collection("user")
	newUser    = usermodel.User{
		ID:       1,
		Email:    "fake@mail.com",
		Password: "Password1",
	}
)

func CreateAdmin(c *gin.Context) {

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

/*func GetAdmin(c *gin.Context) (usermodel.User) {

	var user usermodel.User
	admin := collection.FindOne(c, bson.D{{Key: "id", Value: newUser.ID}}).Decode(&user)
	if admin != nil {
		println("response%s", admin)
		return admin
	}
}*/
