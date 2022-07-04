package student

import (
	"errors"
	"fmt"

	"github.com/hosseinlashgari/IE_HW3/internal/model"
	"github.com/jinzhu/gorm"
)

var (
	ErrStudentNotFound  = errors.New("student not found")
	ErrStudentDuplicate = errors.New("student already exists")
)

type Mysql struct {
	database *gorm.DB
}

func NewMysql(database *gorm.DB) *Mysql {
	database.AutoMigrate(&model.Student{})
	return &Mysql{
		database: database,
	}
}

func (m *Mysql) GetAll() ([]model.Student, error) {
	var students []model.Student
	result := m.database.Find(&students)

	if result.Error != nil {
		return nil, fmt.Errorf("mysql find query failed %w", result.Error)
	}

	return students, nil
}

func (m *Mysql) Get(first_name string, last_name string) (model.Student, error) {
	var std model.Student

	result := m.database.Where("first_name = ? AND last_name = ?", first_name, last_name).Take(&std)

	if result.Error != nil {
		return std, fmt.Errorf("mysql find query failed %w", ErrStudentNotFound)
	}

	return std, nil
}

func (m *Mysql) Set(std model.Student) error {
	result := m.database.Create(&std)

	if result.Error != nil {
		return fmt.Errorf("mysql insertion failed %w", result.Error)
	}

	return nil
}

func (m *Mysql) Delete(std model.Student) error {
	result := m.database.Delete(&std)

	if result.Error != nil {
		return fmt.Errorf("mysql Deletion failed %w", result.Error)
	}

	return nil
}
