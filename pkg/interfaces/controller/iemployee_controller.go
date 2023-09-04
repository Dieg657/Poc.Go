package controller

import (
	interfaces "github.com/Dieg657/Poc.Go/pkg/interfaces/services"
	"github.com/gin-gonic/gin"
)

type IEmployeeController interface {
	New(engine *gin.Engine, service *interfaces.IEmployeeService) error
	RegisterRoutes()
}
