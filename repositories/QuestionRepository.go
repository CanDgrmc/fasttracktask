package repositories

import (
	"github.com/CanDgrmc/gotask/models"
)

type QuestionRepository struct {
	collection string
	questions  []models.Question
}

func NewQuestionRepository() (*QuestionRepository, error) {
	return &QuestionRepository{collection: "questions"}, nil
}

func (r *QuestionRepository) Find(id int) (*models.Question, error) {
	for _, question := range r.questions {
		if question.Id == id {
			return &question, nil
		}
	}
	return nil, nil
}

func (r *QuestionRepository) FindByIdAndQuizId(id int, quizId int) (*models.Question, error) {
	for _, question := range r.questions {
		if question.Id == id && question.QuizId == quizId {
			return &question, nil
		}
	}
	return nil, nil
}

func (r *QuestionRepository) FindAll() (*[]models.Question, error) {

	return &r.questions, nil
}

func (r *QuestionRepository) Add(q models.Question) (int, error) {

	r.questions = append(r.questions, q)

	return q.Id, nil
}
