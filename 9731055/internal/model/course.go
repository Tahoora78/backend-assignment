package model

import "github.com/jinzhu/gorm"

type Course struct {
	gorm.Model
	Name     string `db:"name" json:"name"`
	Lecturer string `db:"lecturer" json:"lecturer"`
}
