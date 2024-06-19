package routes

import (
	"studentmgt/model"
	"studentmgt/service"

	"github.com/gofiber/fiber/v2"
)

func Home(c *fiber.Ctx) error {
	dbconfig := model.DbConfig{
		Host:         "localhost",
		Port:         "5432",
		UserName:     "postgres",
		Password:     "1234",
		DatabaseName: "studentmanagement",
	}

	data, err := service.FindAllStudent(dbconfig)
	if err != nil {
		return err
	}
	return c.Render("home", fiber.Map{
		"Title": "Student Management",
		"Data":  data,
	}, "layout")
}
