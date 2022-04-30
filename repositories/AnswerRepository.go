package repositories

import (
	"github.com/CanDgrmc/gotask/models"
)

type AnswerRepository struct {
	collection string
	answers    []models.Answer
}

func NewAnswerRepository() (*AnswerRepository, error) {
	return &AnswerRepository{collection: "answers"}, nil
}

func (r *AnswerRepository) Find(id int) (*models.Answer, error) {
	for _, answer := range r.answers {
		if answer.Id == id {
			return &answer, nil
		}
	}
	return nil, nil
}

func (r *AnswerRepository) FindAll() (*[]models.Answer, error) {

	return &r.answers, nil
}

func (r *AnswerRepository) FindByQuestionId(id int) (*[]models.Answer, error) {
	var answers []models.Answer
	for _, answer := range r.answers {
		if answer.QuestionId == id {
			answers = append(answers, answer)
		}
	}
	return &answers, nil
}

func (r *AnswerRepository) Add(a models.Answer) (int, error) {

	r.answers = append(r.answers, a)

	return a.Id, nil
}
