package model

import "github.com/jinzhu/gorm"

type Student struct {
	gorm.Model
	FirstName string  `db:"first_name" json:"first_name"`
	LastName  string  `db:"last_name" json:"last_name"`
	Email     string  `db:"email" json:"email"`
	Course    Course  `db:"course" json:"course"`
	Grade     float64 `db:"grade" json:"grade"`
}
