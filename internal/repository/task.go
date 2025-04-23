package repository

import (
	"context"
	"fmt"

	"github.com/mant1COREX/pet-project/internal/entity"
	"github.com/mant1COREX/pet-project/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type taskRepo struct {
	db *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) *taskRepo {
	return &taskRepo{
		db: db,
	}
}

func (t *taskRepo) CreateTask(ctx context.Context, task entity.Task) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, description, created_at, updated_at) values ($1, $2, $3, $4) RETURNING id", postgres.TasksTable)

	row := t.db.QueryRow(ctx, query, task.Title, task.Description, task.CreatedAt, task.UpdatedAt)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (t *taskRepo) DeleteTask(ctx context.Context, id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgres.TasksTable)
	row, err := t.db.Exec(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected := row.RowsAffected()
	if rowsAffected == 0 {
		return 0, ErrTaskNotFound
	}

	return id, err
}

func (t *taskRepo) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	var updatedTask entity.Task

	query := fmt.Sprintf("UPDATE %s SET title = $1, description = $2, status = $3, updated_at = $4 WHERE id = $5 RETURNING id, title, description, status, created_at, updated_at", postgres.TasksTable)
	row := t.db.QueryRow(ctx, query, task.Title, task.Description, task.Status, task.UpdatedAt, task.Id)
	if err := row.Scan(&updatedTask.Id, &updatedTask.Title, &updatedTask.Description, &updatedTask.Status, &updatedTask.CreatedAt, &updatedTask.UpdatedAt); err != nil {
		return entity.Task{}, ErrTaskNotFound
	}

	return updatedTask, nil
}

func (t *taskRepo) GetAllTasks(ctx context.Context) (*[]entity.Task, error) {
	var tasks []entity.Task
	query := fmt.Sprintf("SELECT * FROM %s", postgres.TasksTable)

	rows, err := t.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entity.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return &tasks, nil
}
