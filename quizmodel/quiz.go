package quizmodel

//TODO ID

type Quiz struct {
	ID              int        `json:"id" bson:"id"`
	QuizName        string     `json:"name" bson:"name"`
	QuizDescription string     `json:"description" bson:"description"`
	Questions       []Question `json:"questions" bson:"questions"`
}

type Question struct {
	ID               int      `json:"id" bson:"id"`
	Description      string   `json:"description" bson:"description"`
	IsMultipleChoice bool     `json:"isMultipleChoice" bson:"is_multiple_choice"`
	Answers          []Answer `json:"answers" bson:"answers"`
}

type Answer struct {
	ID          int    `json:"id" bson:"id"`
	Description string `json:"description" bson:"description"`
	IsCorrect   bool   `json:"isCorrect" bson:"is_correct"`
}
