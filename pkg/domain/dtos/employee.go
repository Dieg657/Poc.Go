package dtos

import "time"

type Employee struct {
	ID uint64 `gorm:"primarykey" json:"Id"`
	EmployeeAddModel
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:CreatedAt"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:UpdatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty" gorm:"column:DeletedAt"`
}

type EmployeeAddModel struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   bool    `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}
