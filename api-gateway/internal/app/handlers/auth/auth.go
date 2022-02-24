package auth

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
)

type Handler struct {
	userServiceAddress string
}

func NewAuthHandler() *Handler {
	return &Handler{
		userServiceAddress: viper.GetString("user_service_port"),
	}
}

type loginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type regInput struct {
	loginInput
}

type authResponse struct {
	Token string `json:"token"`
}

func (h *Handler) Registration(ginC *gin.Context) {
	inp := new(regInput)

	if err := ginC.ShouldBindJSON(inp); err != nil {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	conn, err := grpc.Dial(h.userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewUserServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.Registration(ctx, &pb.RegistrationRequest{
		Login:    inp.Login,
		Password: inp.Password,
	})

	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	ginC.Header("Authorization", fmt.Sprintf("%s %s", "Bearer", serviceResp.Token))
	handlers.NewResponseOk(&authResponse{Token: serviceResp.Token}).GinJson(ginC)
}

func (h *Handler) Login(ginC *gin.Context) {
	inp := new(loginInput)

	if err := ginC.ShouldBindJSON(inp); err != nil {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	conn, err := grpc.Dial(h.userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewUserServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.Login(ctx, &pb.LoginRequest{
		Login:    inp.Login,
		Password: inp.Password,
	})
	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	ginC.Header("Authorization", fmt.Sprintf("%s %s", "Bearer", serviceResp.Token))
	handlers.NewResponseOk(&authResponse{Token: serviceResp.Token}).GinJson(ginC)
}

func (h *Handler) RegisterHTTPEndpoints(router *gin.Engine) {
	authEndpoints := router.Group("/api/auth")
	{
		authEndpoints.POST("/registration", h.Registration)
		authEndpoints.POST("/login", h.Login)
	}
}
