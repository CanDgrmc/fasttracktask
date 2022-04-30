package models

type Answer struct {
	Id         int    `json:"id" yaml:"id"`
	QuestionId int    `json:"questionId" yaml:"questionId"`
	Answer     string `json:"answer" yaml:"answer"`
}
