package course

import (
	"errors"
	"fmt"

	"github.com/hosseinlashgari/IE_HW3/internal/model"
	"github.com/jinzhu/gorm"
)

var (
	ErrCourseNotFound  = errors.New("course not found")
	ErrCourseDuplicate = errors.New("course already exists")
)

type Mysql struct {
	database *gorm.DB
}

func NewMysql(database *gorm.DB) *Mysql {
	database.AutoMigrate(&model.Course{})
	return &Mysql{
		database: database,
	}
}

func (m *Mysql) GetAll() ([]model.Course, error) {
	var courses []model.Course
	result := m.database.Find(&courses)

	if result.Error != nil {
		return nil, fmt.Errorf("mysql find query failed %w", result.Error)
	}

	return courses, nil
}

func (m *Mysql) Get(name string) (model.Course, error) {
	var course model.Course

	result := m.database.Where("name = ?", name).Take(course)

	if result.Error != nil {
		return course, fmt.Errorf("mysql find query failed %w", ErrCourseNotFound)
	}

	return course, nil
}

func (m *Mysql) Set(course model.Course) error {
	result := m.database.Create(&course)

	if result.Error != nil {
		return fmt.Errorf("mysql insertion failed %w", result.Error)
	}

	return nil
}

func (m *Mysql) Delete(course model.Course) error {
	result := m.database.Delete(&course)

	if result.Error != nil {
		return fmt.Errorf("mysql Deleteion failed %w", result.Error)
	}

	return nil
}
