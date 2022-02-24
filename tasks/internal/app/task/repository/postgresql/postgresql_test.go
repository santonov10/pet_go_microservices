package PostgreSQL

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"github.com/santonov10/microservices/tasks/internal/pkg/config"
	"github.com/santonov10/microservices/tasks/internal/pkg/db"
	"github.com/stretchr/testify/require"
	"testing"
)

func getConn() (*sql.DB, error) {
	config.SetupPaths("../../../../../config/", "config", "../../../../../config/dev.env")
	err := config.Init()

	if err != nil {
		return nil, err
	}

	dbDSN := config.GetPostgreDSN()

	ctx := context.Background()

	return db.PostgreSQLConnect(ctx, dbDSN)
}

func TestUserPostgresSQL(t *testing.T) {
	t.Run("CRUD", func(t *testing.T) {
		db, err := getConn()
		require.NoError(t, err)
		ctx := context.Background()
		pgSQL := NewTaskPostgresSQL(db)
		testUserID := "123e4567-e89b-12d3-a456-426655440000"
		createData := &task.InsertTask{
			Header:      "test",
			Description: "test",
			UserID:      testUserID,
		}
		createdID, createErr := pgSQL.Create(ctx, createData)

		require.NoError(t, createErr)
		require.NotEmpty(t, createdID)
		defer pgSQL.Delete(ctx, createdID)

		updateData := &task.UpdateTask{
			Header:      "modified",
			Description: "modified",
		}

		err = pgSQL.Update(ctx, createdID, updateData)

		require.NoError(t, err)

		wrongData := &task.UpdateTask{
			Header: `tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_
				tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_
tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_
tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_tooLong_`,
			Description: "testWrongData",
		}

		err = pgSQL.Update(ctx, createdID, wrongData)

		require.Error(t, err)

		findedTasks, err := pgSQL.GetAllForUser(ctx, testUserID)

		require.NoError(t, createErr)
		require.NotEmpty(t, findedTasks)
		foundCreated := false
		for _, task := range findedTasks {
			if createdID == task.ID {
				foundCreated = true
				require.Equal(t, task.Header, updateData.Header)
				require.Equal(t, task.Description, updateData.Description)
			}
			err := pgSQL.Delete(ctx, task.ID)
			require.NoError(t, err)
		}

		require.True(t, foundCreated)
	})
}
