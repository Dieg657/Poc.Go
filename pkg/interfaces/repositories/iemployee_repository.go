package repositories

import (
	"github.com/Dieg657/Poc.Go/pkg/domain/dtos"
	"github.com/Dieg657/Poc.Go/pkg/interfaces/db"
)

type IEmployeeRepository interface {
	New(databaseAdapter db.IDatabaseAdapter) error
	Add(model *dtos.EmployeeAddModel) bool
	GetAll() *[]dtos.Employee
}
