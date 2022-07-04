package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinlashgari/IE_HW3/internal/http/request"
	"github.com/hosseinlashgari/IE_HW3/internal/model"
	"github.com/hosseinlashgari/IE_HW3/internal/store/student"
)

type Student struct {
	Store student.Mysql
}

func (s *Student) List(c *fiber.Ctx) error {
	l, err := s.Store.GetAll()
	if err != nil {
		log.Println("store.getall failed")

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(l)
}

func (s *Student) Get(c *fiber.Ctx) error {
	first_name := c.Params("first_name", "-")
	if first_name == "-" {
		log.Println("cannot get first name")
		return fiber.ErrBadRequest
	}
	last_name := c.Params("last_name", "-")
	if last_name == "-" {
		log.Println("cannot get last name")
		return fiber.ErrBadRequest
	}

	std, err := s.Store.Get(first_name, last_name)
	if err != nil {
		if errors.Is(err, student.ErrStudentNotFound) {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(std)
}

func (s *Student) Create(c *fiber.Ctx) error {
	var req request.Student

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	m := model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Course:    req.Course,
		Grade:     req.Grade,
	}

	if err := s.Store.Set(m); err != nil {
		log.Println(err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(m)
}

func (s *Student) Delete(c *fiber.Ctx) error {
	first_name := c.Params("first_name", "-")
	if first_name == "-" {
		log.Println("cannot get first name")
		return fiber.ErrBadRequest
	}
	last_name := c.Params("last_name", "-")
	if last_name == "-" {
		log.Println("cannot get last name")
		return fiber.ErrBadRequest
	}

	m := model.Student{
		FirstName: first_name,
		LastName:  last_name,
	}

	err := s.Store.Delete(m)
	if err != nil {
		if errors.Is(err, student.ErrStudentNotFound) {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(m)
}

func (s *Student) Register(g fiber.Router) {
	g.Get("/", s.List)
	g.Post("/", s.Create)
	g.Get("/:first_name-:last_name", s.Get)
	g.Post("/delete", s.Delete)
}
