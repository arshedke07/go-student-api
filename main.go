package main

import (
	"studentmgt/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: false,
	})

	app.Get("/", routes.Home)
	app.Get("/addstudent", routes.AddStudent)
	app.Post("/addstudent", routes.AddStudent)
	app.Get("/updatestudent/:id", routes.UpdateStudent)
	app.Post("/updatestudent/:id", routes.UpdateStudent)
	app.Get("/deletestudent/:id", routes.DeleteStudent)
	app.Post("/deletestudent/:id", routes.DeleteStudent)
	app.Listen(":3000")
}
