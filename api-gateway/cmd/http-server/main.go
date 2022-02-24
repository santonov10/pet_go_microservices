package main

import (
	"context"
	"github.com/santonov10/microservices/api-gateway/internal/app/server"
	"github.com/santonov10/microservices/api-gateway/internal/pkg/config"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	serv := server.NewServer()
	serv.Start(ctx)

	<-ctx.Done()
}
