package web

import (
	"context"
	"os"

	"github.com/Dieg657/Poc.Go/pkg/controller"
	"github.com/Dieg657/Poc.Go/pkg/db"
	"github.com/Dieg657/Poc.Go/pkg/docs/swagger"
	"github.com/Dieg657/Poc.Go/pkg/repositories"
	"github.com/Dieg657/Poc.Go/pkg/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type WebApi struct {
}

func (webapi *WebApi) Start(ctx context.Context) {

	// 1 - Initialize database
	database := db.MssqlDatabaseAdapter{}
	err := database.New(os.Getenv("MSSQL"))
	if err != nil {
		return
	}

	// 2 - Initialize repository
	employeeRepository := repositories.EmployeeRepository{}
	err = employeeRepository.New(&database)
	if err != nil {
		return
	}

	// 3 - Initialize service
	employeeService := services.EmployeeService{}
	err = employeeService.New(&employeeRepository)
	if err != nil {
		return
	}

	// 4 - Initialize controller
	employeeController := controller.EmployeeController{}
	router := gin.Default()
	v1 := router.Group("/api/v1")
	err = employeeController.New(v1, &employeeService)
	if err != nil {
		return
	}

	// 5 - Register Routes
	employeeController.RegisterRoutes()

	// 6 - Setup Swagger
	webapi.ConfigureSwagger(router)

	// Run Gin-Gonic
	router.Run()
}

func (webApi *WebApi) ConfigureSwagger(gin *gin.Engine) {
	swagger.SwaggerInfo.Title = "POC WebAPI Golang"
	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
