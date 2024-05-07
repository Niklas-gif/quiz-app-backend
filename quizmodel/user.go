package quizmodel

type User struct {
	email    string `json:"email" bson:"email"`
	password string `json:"password" bson:"password"`
}
