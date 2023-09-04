package services

import (
	"github.com/Dieg657/Poc.Go/pkg/domain/dtos"
	"github.com/Dieg657/Poc.Go/pkg/interfaces/repositories"
)

type IEmployeeService interface {
	New(repository repositories.IEmployeeRepository) error
	AddEmployee(model *dtos.EmployeeAddModel) bool
	GetAllEmployees() *[]dtos.Employee
}
