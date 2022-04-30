package repositories

import (
	"github.com/CanDgrmc/gotask/models"
)

type StatisticRepository struct {
	collection string
	statistics []models.Statistics
}

func NewStatisticRepository() (*StatisticRepository, error) {
	return &StatisticRepository{collection: "statistics"}, nil
}

func (r *StatisticRepository) Find(id string) (*models.Statistics, error) {
	for _, statistic := range r.statistics {
		if statistic.Id == id {
			return &statistic, nil
		}
	}
	return nil, nil
}

func (r *StatisticRepository) FindAll() (*[]models.Statistics, error) {

	return &r.statistics, nil
}

func (r *StatisticRepository) FindByQuizId(quizId int) (*[]models.Statistics, error) {
	var statistics []models.Statistics
	for _, statistic := range r.statistics {
		if statistic.QuizId == quizId {
			statistics = append(statistics, statistic)
		}
	}
	return &statistics, nil
}

func (r *StatisticRepository) FindByQuizIdAndUserId(quizId int, userId int) (*models.Statistics, error) {
	for _, statistic := range r.statistics {
		if statistic.QuizId == quizId && statistic.UserId == userId {
			return &statistic, nil
		}
	}
	return nil, nil
}

func (r *StatisticRepository) Add(a models.Statistics) (string, error) {

	r.statistics = append(r.statistics, a)

	return a.Id, nil
}
