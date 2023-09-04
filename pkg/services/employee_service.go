package services

import (
	"errors"

	"github.com/Dieg657/Poc.Go/pkg/domain/dtos"
	interfaces "github.com/Dieg657/Poc.Go/pkg/interfaces/repositories"
)

type EmployeeService struct {
	repository interfaces.IEmployeeRepository
}

func (es *EmployeeService) New(repository interfaces.IEmployeeRepository) error {
	if repository == nil {
		return errors.New("repository was null")
	}

	es.repository = repository

	return nil
}

func (es *EmployeeService) AddEmployee(model *dtos.EmployeeAddModel) bool {
	return es.repository.Add(model)
}

func (es *EmployeeService) GetAllEmployees() *[]dtos.Employee {
	return es.repository.GetAll()
}
