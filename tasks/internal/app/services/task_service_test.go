package services

import (
	"context"
	"github.com/santonov10/microservices/tasks/api/grpc/pb"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"github.com/santonov10/microservices/tasks/internal/app/task/usecase"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestUserService_GetUser(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		ctx := context.Background()
		uc := &usecase.TaskUseCaseMock{}
		taskService := NewTaskService(uc)

		UserId := "123"
		createdID := "newTaskID"
		createTaskRequest := &pb.CreateTaskRequest{
			Header:      "test",
			Description: "test",
			UserId:      UserId,
		}

		uc.On("CreateTask", &task.InsertTask{
			Header:      createTaskRequest.Header,
			Description: createTaskRequest.Description,
			UserID:      createTaskRequest.UserId,
		}).Return(createdID, nil)

		response, err := taskService.CreateTask(ctx, createTaskRequest)

		require.NoError(t, err)
		require.Equal(t, response.Id, createdID)

		updateTaskRequest := &pb.UpdateTaskRequest{
			Header:      "test2",
			Description: "test2",
			TaskId:      createdID,
		}

		uc.On("UpdateTask", createdID, &task.UpdateTask{
			Header:      updateTaskRequest.Header,
			Description: updateTaskRequest.Description,
		}).Return(nil)

		updateResponse, err := taskService.UpdateTask(ctx, updateTaskRequest)

		require.NoError(t, err)
		require.True(t, updateResponse.Success)

		GetTasksMock := []*models.Task{{
			ID:          createdID,
			UserID:      UserId,
			TimeCreated: time.Now(),
			TimeUpdated: time.Now(),
			Header:      updateTaskRequest.Header,
			Description: updateTaskRequest.Description,
		}}
		uc.On("GetTasks", UserId).Return(GetTasksMock, nil)

		GetTasksResponseResult := &pb.TasksResponse{
			Tasks: []*pb.Task{
				{
					Id:          GetTasksMock[0].ID,
					TimeCreated: timestamppb.New(GetTasksMock[0].TimeCreated),
					TimeUpdated: timestamppb.New(GetTasksMock[0].TimeUpdated),
					Header:      GetTasksMock[0].Header,
					Description: GetTasksMock[0].Description,
					UserId:      GetTasksMock[0].UserID,
				},
			},
		}
		GetTasksResponse, err := taskService.GetTasks(ctx, &pb.GetTasksRequest{UserId: UserId})

		require.NoError(t, err)
		require.EqualValues(t, GetTasksResponse.Tasks[0], GetTasksResponseResult.Tasks[0])

	})
}
