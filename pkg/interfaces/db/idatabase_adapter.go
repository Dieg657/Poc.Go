package db

import "gorm.io/gorm"

type IDatabaseAdapter interface {
	New(connectionString string) error
	GetDatabase() *gorm.DB
}
