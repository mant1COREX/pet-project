package service

import (
	"context"
	"time"

	"github.com/BabyJhon/skillsrock-test-task/internal/entity"
	"github.com/BabyJhon/skillsrock-test-task/internal/repository"
)

type taskService struct {
	repo repository.Task
}

func NewTaskService(repo repository.Task) *taskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) CreateTask(ctx context.Context, task entity.Task) (int, error) {
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	return s.repo.CreateTask(ctx, task)
}

func (s *taskService) DeleteTask(ctx context.Context, id int) (int, error) {
	return s.repo.DeleteTask(ctx, id)
}

func (s *taskService) UpdateTask(ctx context.Context, task entity.Task) (entity.Task, error) {
	task.UpdatedAt = time.Now()
	return s.repo.UpdateTask(ctx, task)
}

func (s *taskService) GetAllTasks(ctx context.Context) (*[]entity.Task, error) {
	return s.repo.GetAllTasks(ctx)
}
