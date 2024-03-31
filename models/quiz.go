package models

//TODO

type Quiz struct {
	QuizDescription string `json: "description" bson: "description"`
	Questions       []Question
}

type Question struct {
	Description      string `json: "description" bson: "description"`
	IsMultipleChoice bool   `json: "isMultipleChoice bson: is_multiple_choice"`
	Answers          []Answer
}

type Answer struct {
	Description string `json: "description" bson: "description"`
	IsCorrect   bool   `json: "isCorrect" bson: "is_correct"`
}
