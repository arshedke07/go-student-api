package routes

import (
	"database/sql"
	"fmt"
	"studentmgt/model"
	"studentmgt/service"

	"github.com/gofiber/fiber/v2"
)

func DeleteStudent(c *fiber.Ctx) error {
	id := c.Params("id")

	dbconfig := model.DbConfig{
		Host:         "localhost",
		Port:         "5432",
		UserName:     "postgres",
		Password:     "1234",
		DatabaseName: "studentmanagement",
	}

	if c.Method() == "GET" {
		data, err := service.FindStudentById(dbconfig, id)
		if err != nil {
			return err
		}

		return c.Render("deletestudent", fiber.Map{
			"Title": "Delete Student",
			"Data":  data,
		}, "layout")

	} else if c.Method() == "POST" {

		err := DeleteStudentFromDb(dbconfig, id)
		if err != nil {
			return err
		}

		data, finderr := service.FindAllStudent(dbconfig)
		if finderr != nil {
			return finderr
		}
		return c.Render("home", fiber.Map{
			"Title": "Student Management",
			"Data":  data,
		}, "layout")
	}
	return nil
}

func DeleteStudentFromDb(dbconfig model.DbConfig, id string) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.UserName, dbconfig.Password, dbconfig.DatabaseName)

	deletestatement := "DELETE FROM student WHERE id = $1"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, deleteErr := db.Exec(deletestatement, id)

	if deleteErr != nil {
		return deleteErr
	}
	return nil
}
