package usecase

import (
	"context"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"github.com/santonov10/microservices/tasks/internal/app/task/repository/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserUseCase(t *testing.T) {
	t.Run("normal data", func(t *testing.T) {

		repoMock := new(mock.TaskRepoMock)
		UC := NewTaskUseCase(repoMock)

		ctx := context.Background()

		UserID := "123"
		taskData := &task.InsertTask{
			Header:      "test",
			Description: "test",
			UserID:      UserID,
		}
		createdMockID := "taskID"

		repoMock.On("Create", taskData).Return(createdMockID, nil)

		taskID, err := UC.CreateTask(ctx, taskData)
		require.NoError(t, err)
		require.NotEmpty(t, taskID)

		taskUpdateData := &task.UpdateTask{
			Header:      "test2",
			Description: "test2",
		}

		repoMock.On("Update", createdMockID, taskUpdateData).Return(nil)
		err = UC.UpdateTask(context.Background(), createdMockID, taskUpdateData)

		require.NoError(t, err)

		mockTasks := []*models.Task{
			{
				ID:          createdMockID,
				UserID:      UserID,
				TimeCreated: time.Time{},
				TimeUpdated: time.Time{},
				Header:      taskUpdateData.Header,
				Description: taskUpdateData.Description,
			},
		}

		repoMock.On("GetAllForUser", UserID).Return(mockTasks, nil)
		tasks, err := UC.GetTasks(context.Background(), UserID)
		require.NoError(t, err)
		require.Equal(t, tasks[0].ID, createdMockID)
		require.Len(t, tasks, 1)

	})

	t.Run("wrong data", func(t *testing.T) {
		repoMock := new(mock.TaskRepoMock)
		UC := NewTaskUseCase(repoMock)

		taskData := &task.InsertTask{
			Header:      "",
			Description: "NoHeader",
			UserID:      "test",
		}

		userID, err := UC.CreateTask(context.Background(), taskData)
		require.ErrorIs(t, err, task.ErrEmptyHeader)
		require.Empty(t, userID)

		taskData = &task.InsertTask{
			Header:      "test",
			Description: "NoUserID",
			UserID:      "",
		}

		userID, err = UC.CreateTask(context.Background(), taskData)
		require.ErrorIs(t, err, task.ErrEmptyUserID)
		require.Empty(t, userID)

		taskUpdateData := &task.UpdateTask{
			Header:      "",
			Description: "NoHeader",
		}

		err = UC.UpdateTask(context.Background(), "userID", taskUpdateData)
		require.ErrorIs(t, err, task.ErrEmptyHeader)
		require.Empty(t, userID)

	})
}
