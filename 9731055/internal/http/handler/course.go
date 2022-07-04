package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/hosseinlashgari/IE_HW3/internal/http/request"
	"github.com/hosseinlashgari/IE_HW3/internal/model"
	"github.com/hosseinlashgari/IE_HW3/internal/store/course"
)

type Course struct {
	Store course.Mysql
}

func (s *Course) List(c *fiber.Ctx) error {
	l, err := s.Store.GetAll()
	if err != nil {
		log.Println("store.getall failed")

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(l)
}

func (s *Course) Get(c *fiber.Ctx) error {
	name := c.Params("name", "-")
	if name == "-" {
		log.Println("cannot get course name")
		return fiber.ErrBadRequest
	}

	crs, err := s.Store.Get(name)
	if err != nil {
		if errors.Is(err, course.ErrCourseNotFound) {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(crs)
}

func (s *Course) Create(c *fiber.Ctx) error {
	var req request.Course

	if err := c.BodyParser(&req); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	if err := req.Validate(); err != nil {
		log.Println(err)

		return fiber.ErrBadRequest
	}

	m := model.Course{
		Name:     req.Name,
		Lecturer: req.Lecturer,
	}

	if err := s.Store.Set(m); err != nil {
		log.Println(err)

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusCreated).JSON(m)
}

func (s *Course) Delete(c *fiber.Ctx) error {
	name := c.Params("name", "-")
	if name == "-" {
		log.Println("cannot get name")
		return fiber.ErrBadRequest
	}
	m := model.Course{
		Name: name,
	}
	err := s.Store.Delete(m)
	if err != nil {
		if errors.Is(err, course.ErrCourseNotFound) {
			return fiber.ErrNotFound
		}

		return fiber.ErrInternalServerError
	}

	return c.Status(http.StatusOK).JSON(m)
}

func (s *Course) Register(g fiber.Router) {
	g.Get("/", s.List)
	g.Post("/", s.Create)
	g.Get("/:name", s.Get)
	g.Post("/delete", s.Delete)
}
