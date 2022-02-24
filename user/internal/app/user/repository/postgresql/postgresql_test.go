package PostgreSQL

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
	"github.com/santonov10/microservices/user/internal/pkg/config"
	"github.com/santonov10/microservices/user/internal/pkg/db"
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
		pgSQL := NewUserPostgresSQL(db)
		createData := models.NewUserRegistration("testlogin", "pass")
		createdID, createErr := pgSQL.Create(ctx, createData)

		//проверка на уже существующий тестовую запись и удаление её
		if createErr != nil {
			foundUser, errGet := pgSQL.Get(ctx, createData.Login, createData.Password)
			require.NoError(t, errGet)
			require.NotEmpty(t, foundUser)
			if errGet == nil {
				err := pgSQL.Delete(ctx, foundUser.ID)
				require.NoError(t, err)
				createdID, createErr = pgSQL.Create(ctx, createData)
			}
		}

		require.NoError(t, createErr)
		require.NotEmpty(t, createdID)
		defer pgSQL.Delete(ctx, createdID)

		foundUser, err := pgSQL.GetById(ctx, createdID)
		require.NoError(t, err)
		require.NotEmpty(t, foundUser)

		foundUser, err = pgSQL.Get(ctx, createData.Login, createData.Password)
		require.NoError(t, err)
		require.NotEmpty(t, foundUser)

		// ошибка на дубль логина в базе
		id, err := pgSQL.Create(ctx, createData)
		require.Error(t, err)
		require.Empty(t, id)

		err = pgSQL.Delete(ctx, foundUser.ID)
		require.NoError(t, err)

		foundUser, err = pgSQL.Get(ctx, createData.Login, createData.Password)
		require.ErrorIs(t, err, user.ErrUserNotFound)
		require.Empty(t, foundUser)

		foundUser, err = pgSQL.GetById(ctx, createdID)
		require.ErrorIs(t, err, user.ErrUserNotFound)
		require.Empty(t, foundUser)
	})
}
