package usermodel

type User struct {
	ID       uint   `json: "id" bson: "_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
