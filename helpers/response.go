package helpers

import (
	"myProject/postgrest/models"

	"github.com/gin-gonic/gin"
)

func GenerateSuccessResponse(g *gin.Context, StatusCode int, data interface{}) {
	g.JSON(StatusCode, models.Response{
		Success: true,
		Data:    data,
		Error:   nil,
	})
}

func GenerateErrorResponse(g *gin.Context, StatusCode int, err error) {
	g.JSON(StatusCode, models.Response{
		Success: false,
		Error:   err.Error(),
	})
}
