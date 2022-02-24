package PostgreSQL

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/santonov10/microservices/tasks/internal/app/models"
	"github.com/santonov10/microservices/tasks/internal/app/task"
	"time"
)

var _ task.Repository = (*TaskPostgresSQL)(nil)

type TaskPostgresSQL struct {
	db *sql.DB
}

func (s *TaskPostgresSQL) Create(ctx context.Context, task *task.InsertTask) (id string, err error) {
	q := `INSERT INTO "task" (user_id,header,description,time_created,time_updated)
							VALUES ($1, $2, $3, $4, $5) RETURNING id;`

	timeCreated := time.Now()

	row := s.db.QueryRowContext(ctx, q,
		task.UserID, task.Header, task.Description, timeCreated, timeCreated,
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

func (s *TaskPostgresSQL) Update(ctx context.Context, id string, task *task.UpdateTask) (err error) {
	q := `UPDATE "task" SET header = $1, description = $2, time_updated = $3
							WHERE id = $4;`

	timeCreated := time.Now()

	_, err = s.db.ExecContext(ctx, q,
		task.Header, task.Description, timeCreated, id,
	)
	return err
}

func (s *TaskPostgresSQL) GetAllForUser(ctx context.Context, userID string) ([]*models.Task, error) {
	q := `SELECT user_id,time_updated,id,time_created,description,header 
			FROM task WHERE user_id = $1`

	rows, err := s.db.QueryContext(ctx, q,
		userID,
	)

	if err != nil {
		return nil, fmt.Errorf("cannot select: %w", err)
	}
	defer rows.Close()

	var tasks []*models.Task
	for rows.Next() {
		t := &models.Task{}

		if err := rows.Scan(
			&t.UserID,
			&t.TimeUpdated,
			&t.ID,
			&t.TimeCreated,
			&t.Description,
			&t.Header,
		); err != nil {
			return nil, fmt.Errorf("cannot scan: %w", err)
		}

		tasks = append(tasks, t)
	}
	return tasks, rows.Err()
}

func (s *TaskPostgresSQL) Delete(ctx context.Context, id string) error {
	q := `DELETE FROM task WHERE id = $1`
	_, err := s.db.ExecContext(ctx, q, id)
	return err
}

func NewTaskPostgresSQL(db *sql.DB) *TaskPostgresSQL {
	return &TaskPostgresSQL{
		db: db,
	}
}
