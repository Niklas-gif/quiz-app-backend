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
	ExampleQuizzes = []Quiz{
		{
			QuizName:        "*SEEDER QUIZ* Science Quiz *SEEDER QUIZ*",
			QuizDescription: "Test your science knowledge",
			Questions: []Question{
				{
					Description:      "What is the chemical symbol for water?",
					IsMultipleChoice: true,
					Answers: []Answer{
						{Description: "H2O", IsCorrect: true},
						{Description: "O2", IsCorrect: false},
						{Description: "CO2", IsCorrect: false},
						{Description: "HO", IsCorrect: false},
					},
				},
				{
					Description:      "What planet is known as the Red Planet?",
					IsMultipleChoice: true,
					Answers: []Answer{
						{Description: "Mars", IsCorrect: true},
						{Description: "Venus", IsCorrect: false},
						{Description: "Jupiter", IsCorrect: false},
						{Description: "Saturn", IsCorrect: false},
					},
				},
			},
		},
		{
			QuizName:        "*SEEDER QUIZ* Math Quiz *SEEDER QUIZ*",
			QuizDescription: "Test your math knowledge",
			Questions: []Question{
				{
					Description:      "What is 2 + 2?",
					IsMultipleChoice: true,
					Answers: []Answer{
						{Description: "4", IsCorrect: true},
						{Description: "22", IsCorrect: false},
						{Description: "3", IsCorrect: false},
						{Description: "5", IsCorrect: false},
					},
				},
				{
					Description:      "What is the square root of 9?",
					IsMultipleChoice: true,
					Answers: []Answer{
						{Description: "3", IsCorrect: true},
						{Description: "81", IsCorrect: false},
						{Description: "9", IsCorrect: false},
						{Description: "27", IsCorrect: false},
					},
				},
			},
		},
	}
)
