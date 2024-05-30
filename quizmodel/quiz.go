package quizmodel

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	ID              primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	QuizName        string             `json:"name" bson:"name"`
	QuizDescription string             `json:"description" bson:"description"`
	Questions       []Question         `json:"questions" bson:"questions"`
}

type Question struct {
	Description      string   `json:"description" bson:"description"`
	IsMultipleChoice bool     `json:"is_multiple_choice" bson:"is_multiple_choice"`
	Answers          []Answer `json:"answers" bson:"answers"`
}

type Answer struct {
	Description string `json:"description" bson:"description"`
	IsCorrect   bool   `json:"is_correct" bson:"is_correct"`
}
