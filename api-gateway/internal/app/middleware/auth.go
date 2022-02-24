package middleware

import (
	"context"
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

const ctxUserIDKey = "userID"

type userAuth struct {
	Authorization string `form:"Authorization"`
}

type AuthMiddleware struct {
	userServiceAddress string
}

func NewAuthMiddleware() gin.HandlerFunc {
	return (&AuthMiddleware{
		userServiceAddress: viper.GetString("user_service_port"),
	}).Handle
}

func (m *AuthMiddleware) Handle(ginC *gin.Context) {
	inp := new(userAuth)

	if err := ginC.ShouldBindHeader(inp); err != nil ||
		inp.Authorization == "" {
		handlers.NewResponseError(http.StatusUnauthorized, ErrAuth).GinJson(ginC)
		ginC.Abort()
		return
	}

	headerParts := strings.Split(inp.Authorization, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		handlers.NewResponseError(http.StatusUnauthorized, ErrAuth).GinJson(ginC)
		ginC.Abort()
		return
	}
	headerToken := headerParts[1]

	conn, err := grpc.Dial(m.userServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, ErrAuth).GinJson(ginC)
		ginC.Abort()
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
			handlers.NewResponseError(http.StatusUnauthorized, err).GinJson(ginC)
		}
		ginC.Abort()
		return
	} else {
		ginC.Set(ctxUserIDKey, serviceResp.Id)
		ginC.Next()
	}
}
