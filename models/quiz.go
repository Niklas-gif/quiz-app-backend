package models

//TODO ID

type Quiz struct {
	QuizDescription string     `json: "description" bson: "description"`
	Questions       []Question `json:"questions" bson:"questions"`
}

type Question struct {
	Description      string   `json: "description" bson: "description"`
	IsMultipleChoice bool     `json: "isMultipleChoice bson: is_multiple_choice"`
	Answers          []Answer `json:"answers" bson:"answers"`
}

type Answer struct {
	Description string `json: "description" bson: "description"`
	IsCorrect   bool   `json: "isCorrect" bson: "is_correct"`
}
