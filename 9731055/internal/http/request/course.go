package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Course struct {
	Name     string `json:"name"`
	Lecturer string `json:"lecturer"`
}

func (s Course) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required),
		validation.Field(&s.Lecturer, validation.Required),
	); err != nil {
		return fmt.Errorf("course validation failed %w", err)
	}

	return nil
}
