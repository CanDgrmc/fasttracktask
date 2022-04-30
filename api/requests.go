package api

import (
	"net/http"
)

type answerRequest struct {
	QuestionId int `json:"questionId"`
	AnswerId   int `json:"answerId"`
}
type AnswersRequest struct {
	UserId  int             `json:"userId"`
	QuizId  int             `json:"quizId"`
	Answers []answerRequest `json:"answers"`
}

func (d *AnswersRequest) Bind(r *http.Request) error {
	return nil
}
