package main

type Quiz struct {
	quiz_description string
	questions        []Question
}

type Question struct {
	description      string
	isMultibleChoice bool
	answers          []Answer
}

type Answer struct {
	description string
	isCorrect   bool
}
