package services

import (
	"context"
	"myProject/postgrest/models"
	"myProject/postgrest/repository"
)

type ToDoService struct {
	Repository repository.ToDoRepositoryInterface
}

func NewToDoService(repo repository.ToDoRepositoryInterface) *ToDoService {
	return &ToDoService{
		Repository: repo,
	}
}

func (t *ToDoService) GetAllTasksService(ctx context.Context) (resp []models.Todo, err error) {
	resp, err = t.Repository.GetAllTasks(ctx)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *ToDoService) GetTaskByIdService(ctx context.Context, id string) (resp []models.Todo, err error) {
	resp, err = t.Repository.GetTaskById(ctx, id)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *ToDoService) CreateTaskService(ctx context.Context, input models.Todo) (err error) {
	err = t.Repository.CreateTask(ctx, input)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDoService) UpdateTaskByIdService(ctx context.Context, id string, input models.TodoUpdate) (err error) {
	err = t.Repository.UpdateTask(ctx, input, id)
	if err != nil {
		return err
	}

	return nil
}

func (t *ToDoService) DeleteTaskByIdService(ctx context.Context, id string) (err error) {
	err = t.Repository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
