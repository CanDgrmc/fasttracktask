package api

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/CanDgrmc/gotask/models"
	"github.com/CanDgrmc/gotask/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
)

func New(
	enableCORS bool,
	config *Configuration,
	questionRepository *repositories.QuestionRepository,
	answerRepository *repositories.AnswerRepository,
	statisticRepository *repositories.StatisticRepository,
) (*chi.Mux, error) {

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Use(render.SetContentType(render.ContentTypeJSON))

	if enableCORS {
		r.Use(corsConfig().Handler)
	}

	for _, question := range config.Questions {
		questionRepository.Add(question)
	}

	for _, answer := range config.Answers {
		answerRepository.Add(answer)
	}
	r.Get("/questions", func(w http.ResponseWriter, r *http.Request) {
		listQuestions(w, r, questionRepository, answerRepository)
	})

	r.Get("/answers/{questionId}", func(w http.ResponseWriter, r *http.Request) {
		listAnswersByQuestionId(w, r, answerRepository)
	})

	r.Post("/answer", func(w http.ResponseWriter, r *http.Request) {
		answerQuestions(w, r, answerRepository, questionRepository, statisticRepository)
	})

	return r, nil
}

func corsConfig() *cors.Cors {

	return cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           86400, // Maximum value not ignored by any of major browsers
	})
}

func listQuestions(w http.ResponseWriter, r *http.Request, questionRepository *repositories.QuestionRepository, answerRepository *repositories.AnswerRepository) {
	var (
		questions         *[]models.Question
		questionsResponse []models.Question
		err               error
	)
	if questions, err = questionRepository.FindAll(); err != nil {
		render.Respond(w, r, nil)
	}

	for _, question := range *questions {
		if answers, err := answerRepository.FindByQuestionId(question.Id); err == nil {
			question.Answers = *answers
			questionsResponse = append(questionsResponse, question)
		}
	}

	response := GetAllQuestions{
		Success: true,
		Data:    &questionsResponse,
	}
	render.Respond(w, r, response)

}

func listAnswersByQuestionId(w http.ResponseWriter, r *http.Request, answerRepository *repositories.AnswerRepository) {
	var (
		answers    *[]models.Answer
		err        error
		questionId int
	)
	questionIdStr := chi.URLParam(r, "questionId")
	if questionId, err = strconv.Atoi(questionIdStr); err == nil {
		fmt.Printf("i=%d, type: %T\n", questionId, questionId)
	}

	if answers, err = answerRepository.FindByQuestionId(questionId); err != nil {
		render.Respond(w, r, nil)
	}

	response := GetAllAnswers{
		Success: true,
		Data:    answers,
	}
	render.Respond(w, r, response)

}

func answerQuestions(
	w http.ResponseWriter,
	r *http.Request,
	answerRepository *repositories.AnswerRepository,
	questionRepository *repositories.QuestionRepository,
	statisticRepository *repositories.StatisticRepository,
) {
	var (
		err               error
		correct           int = 0
		wrong             int = 0
		stats             *[]models.Statistics
		newStat           models.Statistics
		successPercentage int
	)
	data := &AnswersRequest{}
	if err = render.Bind(r, data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(http.ErrAbortHandler)
	}

	if exist, err := statisticRepository.FindByQuizIdAndUserId(data.QuizId, data.UserId); err != nil || exist != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		panic(http.ErrAbortHandler)
	}
	var arr = make(map[int]int)

	for _, answer := range data.Answers {
		if _, ok := arr[answer.QuestionId]; ok {
			http.Error(w, "duplicate-answer", http.StatusAlreadyReported)
			panic(http.ErrAbortHandler)
		} else {
			arr[answer.QuestionId] = 1
			var (
				err      error
				question *models.Question
			)
			if question, err = questionRepository.FindByIdAndQuizId(answer.QuestionId, data.QuizId); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(http.ErrAbortHandler)
			}

			if question.Answer == answer.AnswerId {
				correct++
			} else {
				wrong++
			}
		}

	}
	newStat.Id = uuid.New().String()
	newStat.UserId = data.UserId
	newStat.QuizId = data.QuizId
	newStat.Correct = correct
	newStat.Wrong = wrong

	if _, err = statisticRepository.Add(newStat); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if stats, err = statisticRepository.FindByQuizId(data.QuizId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(http.ErrAbortHandler)
	}
	sort.Sort(models.ByCorrectAnswers(*stats))
	statsLen := len(*stats)
	others := statsLen - 1
	index := findIndexOfStats(*stats, newStat)

	if index == -1 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		panic(http.ErrAbortHandler)
	}
	position := index + 1
	if others > 0 {
		successPercentage = 100 / others * (statsLen - position)

	} else {
		successPercentage = 100
	}

	response := PostAnswer{
		Success: true,
		Data:    newStat,
		Message: fmt.Sprintf("You scored higher than %d%s of all quizzers", successPercentage, "%"),
	}
	render.Respond(w, r, response)

}

func findIndexOfStats(sortedStats []models.Statistics, stat models.Statistics) int {
	for index, item := range sortedStats {
		if item.Id == stat.Id {
			return index
		}
	}

	return -1
}
