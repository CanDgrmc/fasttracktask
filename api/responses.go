package api

import "github.com/CanDgrmc/gotask/models"

type GetAllQuestions struct {
	Data    *[]models.Question `json:"data"`
	Success bool               `json:"success"`
}

type GetAllAnswers struct {
	Data    *[]models.Answer `json:"data"`
	Success bool             `json:"success"`
}

type GetAnswer struct {
	Data    *models.Answer `json:"data"`
	Success bool           `json:"success"`
}

type PostAnswer struct {
	Success bool              `json:"success"`
	Data    models.Statistics `json:"data"`
	Message string            `json:"message"`
}

type GetQuestionStats struct {
	Success bool                 `json:"success"`
	Data    *[]models.Statistics `json:"data"`
}
