package server

import (
	"context"
	"fmt"
	"github.com/santonov10/microservices/api-gateway/internal/app/handlers/auth"
	task "github.com/santonov10/microservices/api-gateway/internal/app/handlers/tasks"
	"github.com/santonov10/microservices/api-gateway/internal/app/handlers/user"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct { // TODO
	server *http.Server
}

func NewServer() *Server {
	r := gin.Default()

	authHandler := auth.NewAuthHandler()
	authHandler.RegisterHTTPEndpoints(r)
	userHandler := user.NewUserHandler()
	userHandler.RegisterHTTPEndpoints(r)
	taskHandler := task.NewTaskHandler()
	taskHandler.RegisterHTTPEndpoints(r)

	s := &http.Server{
		Addr:           viper.GetString("app_port"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		server: s,
	}
}

func (s *Server) Start(ctx context.Context) {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()
	go func() {
		<-ctx.Done()
		s.Stop(ctx)
	}()
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	fmt.Println("закрываем сервер")
	return err
}
