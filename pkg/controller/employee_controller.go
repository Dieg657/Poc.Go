package controller

import (
	"errors"
	"net/http"

	"github.com/Dieg657/Poc.Go/pkg/domain/dtos"
	"github.com/Dieg657/Poc.Go/pkg/interfaces/services"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	routerGroup *gin.RouterGroup
	service     services.IEmployeeService
}

func (controller *EmployeeController) New(routerGroup *gin.RouterGroup, service services.IEmployeeService) error {
	if routerGroup == nil {
		return errors.New("gin engine instance was null")
	}

	if service == nil {
		return errors.New("service instance was null")
	}

	controller.routerGroup = routerGroup
	controller.service = service
	return nil
}

func (controller *EmployeeController) RegisterRoutes() {
	controller.routerGroup = controller.routerGroup.Group("/employee")
	controller.PostEmployee()
	controller.GetAllEmployees()
}

// @BasePath /api/v1/employee
// @Tags employee
// @Summary Insert employee
// @Schemes
// @Param employeeModel body dtos.EmployeeAddModel yes "Employee Form Data"
// @Description Add employee to database
// @Produce json
// @Success 200 {object} string
// @Router /employee [post]
func (controller *EmployeeController) PostEmployee() {
	controller.routerGroup.POST("/", func(gin *gin.Context) {
		var model dtos.EmployeeAddModel
		if err := gin.BindJSON(&model); err != nil {
			return
		}
		controller.service.AddEmployee(&model)
		gin.JSON(http.StatusCreated, nil)
	})
}

// @BasePath /api/v1/employee
// @Tags employee
// @Summary Get all persisted employees on database
// @Schemes
// @Description List of employees
// @Produce json
// @Success 200 {array} dtos.Employee
// @Router /employee [get]
func (controller *EmployeeController) GetAllEmployees() {
	controller.routerGroup.GET("/", func(gin *gin.Context) {
		gin.JSON(http.StatusOK, controller.service.GetAllEmployees())
	})
}
