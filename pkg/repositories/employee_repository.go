package repositories

import (
	"errors"

	"github.com/Dieg657/Poc.Go/pkg/domain/dtos"
	"github.com/Dieg657/Poc.Go/pkg/interfaces/db"
)

type EmployeeRepository struct {
	dbAdapter db.IDatabaseAdapter
}

func (er *EmployeeRepository) New(databaseAdapter db.IDatabaseAdapter) error {
	if databaseAdapter == nil {
		return errors.New("databaseAdapter was null")
	}

	er.dbAdapter = databaseAdapter

	return nil
}

func (er *EmployeeRepository) Add(model *dtos.EmployeeAddModel) bool {

	er.dbAdapter.GetDatabase().Omit("UpdatedAt", "DeletedAt").Table("Employees").Create(model)

	return true
}

func (er *EmployeeRepository) GetAll() *[]dtos.Employee {

	var employees []dtos.Employee

	result := er.dbAdapter.GetDatabase().Omit("DeletedAt").Find(&employees)

	if result.Error != nil {
		return nil
	}

	return &employees
}
