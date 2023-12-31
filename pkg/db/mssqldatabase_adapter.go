package db

import (
	"errors"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type MssqlDatabaseAdapter struct {
	database *gorm.DB
}

func (mssqlAdapter *MssqlDatabaseAdapter) New(connectionString string) error {
	if len(connectionString) == 0 {
		return errors.New("connectionString was empty")
	}

	db, err := gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err == nil {
		mssqlAdapter.database = db
		return err
	}

	defer mssqlAdapter.Close()

	return err
}

func (mssqlAdapter *MssqlDatabaseAdapter) Close() {
	dbInstance, _ := mssqlAdapter.database.DB()
	_ = dbInstance.Close()
}

func (mssqlAdapter *MssqlDatabaseAdapter) GetDatabase() *gorm.DB {
	return mssqlAdapter.database
}
