package http

import (
	h "myProject/postgrest/helpers"
	"myProject/postgrest/models"
	"myProject/postgrest/repository"
	"myProject/postgrest/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ToDoService struct {
	Service *services.ToDoService
}

func NewToDoHandler() *ToDoService {
	return &ToDoService{
		Service: services.NewToDoService(
			&repository.ToDoRepositoryImp{},
		),
	}
}

func (t *ToDoService) GetAllTasks(g *gin.Context) {
	ctx := g.Request.Context()

	resp, err := t.Service.GetAllTasksService(ctx)
	if err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
		return
	}

	h.GenerateSuccessResponse(g, http.StatusOK, resp)
}

func (t *ToDoService) GetTaskById(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")

	resp, err := t.Service.GetTaskByIdService(ctx, id)
	if err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
		return
	}

	h.GenerateSuccessResponse(g, http.StatusOK, resp)
}

func (t *ToDoService) CreateTask(g *gin.Context) {
	ctx := g.Request.Context()

	var task models.Todo

	if err := g.Bind(&task); err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
	}

	err := t.Service.CreateTaskService(ctx, task)
	if err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
		return
	}

	h.GenerateSuccessResponse(g, http.StatusOK, nil)
}

func (t *ToDoService) UpdateTaskById(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")

	var task models.TodoUpdate

	if err := g.Bind(&task); err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
	}

	err := t.Service.UpdateTaskByIdService(ctx, id, task)
	if err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
		return
	}

	h.GenerateSuccessResponse(g, http.StatusOK, nil)
}

func (t *ToDoService) DeleteTaskById(g *gin.Context) {
	ctx := g.Request.Context()

	id := g.Param("id")

	err := t.Service.DeleteTaskByIdService(ctx, id)
	if err != nil {
		h.GenerateErrorResponse(g, http.StatusBadRequest, err)
		return
	}

	h.GenerateSuccessResponse(g, http.StatusOK, nil)
}
