package main

import (
	"context"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/santonov10/microservices/user/api/grpc/pb"
	"github.com/santonov10/microservices/user/internal/app/services"
	UserRepo "github.com/santonov10/microservices/user/internal/app/user/repository/postgresql"
	"github.com/santonov10/microservices/user/internal/app/user/usecase"
	"github.com/santonov10/microservices/user/internal/pkg/config"
	"github.com/santonov10/microservices/user/internal/pkg/db"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	lsn, err := net.Listen("tcp", viper.GetString("app_service_port"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	server := grpc.NewServer()

	dbDSN := config.GetPostgreDSN()
	pgDB, err := db.PostgreSQLConnect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("ошибка соединения с БД dsn = %s \r\n %s", dbDSN, err)
	}
	defer pgDB.Close()

	UserPGRepo := UserRepo.NewUserPostgresSQL(pgDB)
	UserUC := usecase.NewUserUseCase(UserPGRepo)

	pb.RegisterUserServiceServer(server, services.NewUserService(UserUC))

	log.Printf("starting server on %s", lsn.Addr().String())
	if err := server.Serve(lsn); err != nil {
		log.Fatal(err)
	}
}
