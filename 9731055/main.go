package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/hosseinlashgari/IE_HW3/internal/http/handler"
	"github.com/hosseinlashgari/IE_HW3/internal/store/course"
	"github.com/hosseinlashgari/IE_HW3/internal/store/student"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "root:@/ie_hw3?parseTime=true")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	defer db.Close()

	std_hnd := handler.Student{
		Store: *student.NewMysql(db),
	}

	crs_hnd := handler.Course{
		Store: *course.NewMysql(db),
	}

	app := fiber.New()

	std_hnd.Register(app.Group("/students"))
	crs_hnd.Register(app.Group("/courses"))

	if err := app.Listen("0.0.0.0:1379"); err != nil {
		log.Println("cannot listen")
	}
}
