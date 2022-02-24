package services

import (
	"context"
	"github.com/santonov10/microservices/tasks/api/grpc/pb"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer

	taskUC task.UseCase
}

func (t *TaskService) CreateTask(ctx context.Context, request *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	taskID, err := t.taskUC.CreateTask(ctx, &task.InsertTask{
		Header:      request.Header,
		Description: request.Description,
		UserID:      request.UserId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTaskResponse{Id: taskID}, nil
}

func (t *TaskService) UpdateTask(ctx context.Context, request *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	err := t.taskUC.UpdateTask(ctx, request.TaskId, &task.UpdateTask{
		Header:      request.Header,
		Description: request.Description,
	})
	if err != nil {
		return &pb.UpdateTaskResponse{Success: false}, err
	}
	return &pb.UpdateTaskResponse{Success: true}, nil
}

func (t *TaskService) GetTasks(ctx context.Context, request *pb.GetTasksRequest) (*pb.TasksResponse, error) {
	tasks, err := t.taskUC.GetTasks(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	respTasks := make([]*pb.Task, 0, len(tasks))

	for _, t := range tasks {
		respTasks = append(respTasks, &pb.Task{
			Id:          t.ID,
			TimeCreated: timestamppb.New(t.TimeCreated),
			TimeUpdated: timestamppb.New(t.TimeUpdated),
			Header:      t.Header,
			Description: t.Description,
			UserId:      t.UserID,
		})
	}
	return &pb.TasksResponse{Tasks: respTasks}, nil
}

func NewTaskService(taskUC task.UseCase) *TaskService {
	return &TaskService{
		taskUC: taskUC,
	}
}
