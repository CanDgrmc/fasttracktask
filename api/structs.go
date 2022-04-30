package api

import (
	"github.com/CanDgrmc/gotask/models"
)

type Configuration struct {
	Questions []models.Question `yaml:"questions"`
	Answers   []models.Answer   `yaml:"answers"`
	Port      string            `yaml:"port"`
	LogLevel  string            `yaml:"log_level"`
}
