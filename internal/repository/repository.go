package repository

import (
	"context"

	"github.com/mant1COREX/pet-project/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Task interface {
	CreateTask(ctx context.Context, task entity.Task) (int, error)
	DeleteTask(ctx context.Context, id int) (int, error)
	UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error)
	GetAllTasks(ctx context.Context) (*[]entity.Task, error)
}

type Repository struct {
	Task
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		Task: NewTaskRepo(db),
	}
}
