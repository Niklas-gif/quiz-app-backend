package usermodel

type User struct {
	ID       int    `json: "id" bson: "id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
