package models

type Statistics struct {
	Id      string `json:"id"`
	UserId  int    `json:"-"`
	QuizId  int    `json:"quizId"`
	Correct int    `json:"correct"`
	Wrong   int    `json:"wrong"`
}

type ByCorrectAnswers []Statistics

func (a ByCorrectAnswers) Len() int           { return len(a) }
func (a ByCorrectAnswers) Less(i, j int) bool { return a[i].Correct > a[j].Correct }
func (a ByCorrectAnswers) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
