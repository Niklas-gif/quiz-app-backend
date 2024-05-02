package quizmodel

var (
	SampleQuiz = Quiz{
		QuizName:        "Quiz Example",
		QuizDescription: "Sample Quiz",
		Questions: []Question{{
			Description:      "Was ist die Hauptstadt von Deutschland?",
			IsMultipleChoice: false,
			Answers: []Answer{
				{Description: "Paris", IsCorrect: false},
				{Description: "London", IsCorrect: false},
				{Description: "Berlin", IsCorrect: true},
			},
		}, {
			Description:      "Was ist 2 + 2?",
			IsMultipleChoice: false,
			Answers: []Answer{
				{Description: "4", IsCorrect: true},
				{Description: "42", IsCorrect: false},
				{Description: "Banana", IsCorrect: false},
			},
		},
		},
	}
)
