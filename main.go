package main

import (
	"fmt"
	"myProject/postgrest/config"
	"myProject/postgrest/http"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error while getting current working directory")
		return
	}

	config.Configuration(workingDir)
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	http.Router(app)

	fmt.Println("Server running at 8080...")

	if err := app.Run(":8080"); err != nil {
		panic(err)
	}
}
