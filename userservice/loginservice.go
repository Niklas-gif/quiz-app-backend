package userservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"quiz-app/database"
	"quiz-app/usermodel"
	"quiz-app/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginService(w http.ResponseWriter, r *http.Request, c *gin.Context) {
	w.Header().Set("Content-Type", "application/json")

	//TODO this is just for testing!
	var user usermodel.User
	var _ = database.DB.Collection("user").FindOne(c, bson.D{{Key: "id", Value: "1"}}).Decode(&user)
	json.NewDecoder(r.Body).Decode(&user)
	fmt.Printf("The user request value %v", user)

	if user.Email == "fake@mail.com" && user.Password == "Password1" {
		tokenString, err := utils.CreateToken(user.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, tokenString)
		return
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid credentials")
	}
}
