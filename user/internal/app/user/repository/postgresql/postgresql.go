package PostgreSQL

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/santonov10/microservices/user/internal/app/models"
	"github.com/santonov10/microservices/user/internal/app/user"
)

type UserPostgresSQL struct {
	db *sql.DB
}

func (s *UserPostgresSQL) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM "user" WHERE id = $1`
	_, err := s.db.ExecContext(ctx, q, id)
	return err
}

func (s *UserPostgresSQL) Create(ctx context.Context, user *models.User) (id string, err error) {
	q := `INSERT INTO "user" ("login", "password")
							VALUES ($1, $2) RETURNING id;`

	row := s.db.QueryRowContext(ctx, q,
		user.Login, user.Password,
	)

	if row.Err() != nil {
		return "", fmt.Errorf("create: %w", row.Err())
	}

	dbErr := row.Scan(&id)

	if dbErr != nil {
		return "", fmt.Errorf("create:%w", dbErr)
	}

	return id, nil
}

func (s *UserPostgresSQL) Get(ctx context.Context, login, password string) (*models.User, error) {
	q := `SELECT id,login,password FROM "user" WHERE login = $1 AND password = $2;`

	row := s.db.QueryRowContext(ctx, q, login, password)
	if row.Err() != nil {
		return nil, fmt.Errorf("get:%w", row.Err())
	}

	var foundUser models.User
	err := row.Scan(&foundUser.ID, &foundUser.Login, &foundUser.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf("get:%w", err)
	}
	return &foundUser, nil
}

func (s *UserPostgresSQL) GetById(ctx context.Context, id string) (*models.User, error) {
	q := `SELECT id,login,password FROM "user" WHERE id = $1;`

	row := s.db.QueryRowContext(ctx, q, id)
	if row.Err() != nil {
		return nil, fmt.Errorf("GetByID:%w", row.Err())
	}

	var foundUser models.User
	err := row.Scan(&foundUser.ID, &foundUser.Login, &foundUser.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user.ErrUserNotFound
		}
		return nil, fmt.Errorf("GetByID:%w", err)
	}
	return &foundUser, nil
}

//func (s *EventPostgresqlStorage) Delete(ctx context.Context, eventID string) error {
//	q := `DELETE FROM event WHERE id = $1`
//	_, err := s.db.ExecContext(ctx, q,
//		eventID,
//	)
//	return err
//}

func NewUserPostgresSQL(db *sql.DB) *UserPostgresSQL {
	return &UserPostgresSQL{
		db: db,
	}
}
