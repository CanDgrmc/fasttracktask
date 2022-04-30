package models

type Question struct {
	Id       int      `json:"id" yaml:"id"`
	Question string   `json:"question" yaml:"question"`
	Answer   int      `json:"-" yaml:"answer"`
	Answers  []Answer `json:"answers"`
	QuizId   int      `json:"quizId" yaml:"answer"`
}
