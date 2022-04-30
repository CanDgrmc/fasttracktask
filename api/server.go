package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/CanDgrmc/gotask/repositories"

	"github.com/spf13/viper"
)

type Server struct {
	*http.Server
}

func NewServer(config *Configuration) (*Server, error) {

	// create repositories
	// singleton db should be created here and pass it to repositories
	questionRepository, err := repositories.NewQuestionRepository()
	answerRepository, err := repositories.NewAnswerRepository()
	statisticRepository, err := repositories.NewStatisticRepository()

	if err != nil {
		return nil, err
	}

	// create api
	// should configure cors depending on parameter in configuration
	api, err := New(viper.GetBool("enable_cors"), config, questionRepository, answerRepository, statisticRepository)
	if err != nil {
		return nil, err
	}

	port := config.Port
	addr := ":" + port

	srv := http.Server{
		Addr:    addr,
		Handler: api,
	}

	return &Server{&srv}, nil
}

func (srv *Server) Start() {
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()
	log.Printf("Listening on %s\n", srv.Addr)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Println("Shutting down server... Reason:", sig)

	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err)
	}
	log.Println("stopped")
}
