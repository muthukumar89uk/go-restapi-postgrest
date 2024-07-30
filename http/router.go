package http

import (
	c "myProject/postgrest/helpers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router(app *gin.Engine) {
	handler := NewToDoHandler()
	app.GET(viper.GetString(c.GetAllTasksEndpoint), handler.GetAllTasks)
	app.GET(viper.GetString(c.GetTaskByIdEndpoint), handler.GetTaskById)
	app.POST(viper.GetString(c.CreateTaskEndpoint), handler.CreateTask)
	app.PATCH(viper.GetString(c.UpdateTaskEndpoint), handler.UpdateTaskById)
	app.DELETE(viper.GetString(c.DeleteTaskEndpoint), handler.DeleteTaskById)
}
