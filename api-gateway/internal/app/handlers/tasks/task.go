package task

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/santonov10/microservices/api-gateway/api/grpc/pb"
	"github.com/santonov10/microservices/api-gateway/internal/app/handlers"
	"github.com/santonov10/microservices/api-gateway/internal/app/middleware"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net/http"
	"time"
)

type Handler struct {
	taskServiceAddress string
}

func NewTaskHandler() *Handler {
	return &Handler{
		taskServiceAddress: viper.GetString("task_service_port"),
	}
}

type taskResponse struct {
	ID          string    `json:"id"`
	TimeCreated time.Time `json:"time_created"`
	TimeUpdated time.Time `json:"time_updated"`
	Header      string    `json:"header"`
	Description string    `json:"description"`
	UserID      string    `json:"user_id"`
}

type getTaskRequest struct {
	UserID string `form:"user_id"`
}

type getTaskResponse struct {
	Tasks []*taskResponse `json:"tasks"`
}

type createTaskRequest struct {
	UserID      string `form:"user_id"`
	Header      string `form:"header"`
	Description string `form:"description"`
}

type createTaskResponse struct {
	TaskID string `json:"task_id"`
}

type updateTaskRequest struct {
	TaskID      string `form:"task_id"`
	Header      string `form:"header"`
	Description string `form:"description"`
}

type updateTaskResponse struct {
	Success bool `json:"success"`
}

func (h *Handler) GetTasks(ginC *gin.Context) {
	inp := new(getTaskRequest)

	if err := ginC.ShouldBindQuery(inp); err != nil ||
		inp.UserID == "" {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	conn, err := grpc.Dial(h.taskServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewTaskServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.GetTasks(ctx, &pb.GetTasksRequest{
		UserId: inp.UserID,
	})

	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	tasksResponse := make([]*taskResponse, len(serviceResp.Tasks))
	for i, v := range serviceResp.Tasks {
		tasksResponse[i] = &taskResponse{
			ID:          v.Id,
			TimeCreated: v.TimeCreated.AsTime(),
			TimeUpdated: v.TimeUpdated.AsTime(),
			Header:      v.Header,
			Description: v.Description,
			UserID:      v.UserId,
		}
	}

	handlers.NewResponseOk(
		getTaskResponse{Tasks: tasksResponse},
	).GinJson(ginC)
}

func (h *Handler) CreateTask(ginC *gin.Context) {
	inp := new(createTaskRequest)

	if err := ginC.ShouldBind(inp); err != nil ||
		inp.UserID == "" || inp.Header == "" {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	conn, err := grpc.Dial(h.taskServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewTaskServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.CreateTask(ctx, &pb.CreateTaskRequest{
		Header:      inp.Header,
		Description: inp.Description,
		UserId:      inp.UserID,
	})

	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	handlers.NewResponseOk(
		&createTaskResponse{TaskID: serviceResp.Id},
	).GinJson(ginC)
}

func (h *Handler) UpdateTask(ginC *gin.Context) {
	inp := new(updateTaskRequest)

	if err := ginC.ShouldBind(inp); err != nil ||
		inp.TaskID == "" || inp.Header == "" {
		handlers.NewResponseError(http.StatusBadRequest, handlers.ErrWrongDataFormat).GinJson(ginC)
		return
	}

	conn, err := grpc.Dial(h.taskServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		handlers.NewResponseError(http.StatusServiceUnavailable, handlers.ErrServiceIsDown).GinJson(ginC)
		return
	}
	defer conn.Close()

	serviceClient := pb.NewTaskServiceClient(conn)

	ctx := context.Background()
	serviceResp, err := serviceClient.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Header:      inp.Header,
		Description: inp.Description,
		TaskId:      inp.TaskID,
	})

	if err != nil {
		if status.Code(err) == codes.Unavailable {
			handlers.NewResponseError(http.StatusBadRequest, handlers.ErrServiceIsDown).GinJson(ginC)
		} else {
			handlers.NewResponseError(http.StatusBadRequest, err).GinJson(ginC)
		}
		return
	}

	handlers.NewResponseOk(
		&updateTaskResponse{Success: serviceResp.Success},
	).GinJson(ginC)
}

func (h *Handler) RegisterHTTPEndpoints(router *gin.Engine) {
	group := router.Group("/api/task")
	group.Use(middleware.NewAuthMiddleware())
	{
		group.GET("", h.GetTasks)
		group.POST("", h.CreateTask)
		group.PUT("", h.UpdateTask)
	}
}
