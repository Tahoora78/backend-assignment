package request

import (
	"fmt"
	// "regexp"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hosseinlashgari/IE_HW3/internal/model"
)

const MaxGrade = 20

type Student struct {
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Course    model.Course `json:"course"`
	Grade     float64      `json:"grade"`
}

func (s Student) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		// validation.Field(&s.Email, validation.Required, validation.Match(regexp.MustCompile("/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:.[a-zA-Z0-9-]+)*$/"))),
		// validation.Field(&s.Grade, validation.Required, validation.Min(0), validation.Max(MaxGrade)),
	); err != nil {
		return fmt.Errorf("student validation failed %w", err)
	}

	return nil
}
