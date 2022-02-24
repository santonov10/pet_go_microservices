package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/santonov10/microservices/api-gateway/api/grpc/pb"
	"github.com/santonov10/microservices/api-gateway/internal/app/handlers"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"strings"
)

type Handler struct {
	userServiceAddress string
}

func NewUserHandler() *Handler {
	return &Handler{
		userServiceAddress: viper.GetString("user_service_port"),
	}
}

type getUserInput struct {
	Authorization string `form:"Authorization"`
}

type userResponse struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}

func (h *Handler) GetUser(ginC *gin.Context) {
	inp := new(getUserInput)

	if err := ginC.ShouldBindHeader(inp); err != nil ||
		inp.Authorization == "" {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	headerParts := strings.Split(inp.Authorization, " ")
	fmt.Println(headerParts)
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}
	headerToken := headerParts[1]

	conn, err := grpc.Dial(h.userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewUserServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.GetUser(ctx, &pb.Token{
		Token: headerToken,
	})

	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	handlers.NewResponseOk(&userResponse{
		ID:    serviceResp.Id,
		Login: serviceResp.Login,
	}).GinJson(ginC)
}

func (h *Handler) RegisterHTTPEndpoints(router *gin.Engine) {
	router.GET("/api/user", h.GetUser)
}
