package repository

import (
	"context"
	"myProject/postgrest/models"
)

type ToDoRepositoryInterface interface {
	GetAllTasks(ctx context.Context) (resp []models.Todo, err error)
	GetTaskById(ctx context.Context, id string) (resp []models.Todo, err error)
	CreateTask(ctx context.Context, task models.Todo) (err error)
	UpdateTask(ctx context.Context, task models.TodoUpdate, id string) (err error)
	DeleteTask(ctx context.Context, id string) (err error)
}
