package routes

import (
	"database/sql"
	"fmt"
	"studentmgt/model"
	"studentmgt/service"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func UpdateStudent(c *fiber.Ctx) error {
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
		return c.Render("updatestudent", fiber.Map{
			"Title": "Edit Existing Student",
			"Data":  data,
		}, "layout")
	} else if c.Method() == "POST" {
		firstname := c.FormValue("firstname")
		lastname := c.FormValue("lastname")
		emailid := c.FormValue("emailid")
		attendance := c.FormValue("attendance")

		// fmt.Printf("%s, %s, %s", firstname, lastname, emailid)

		student := model.Student{
			FirstName:  firstname,
			LastName:   lastname,
			EmailId:    emailid,
			Attendance: attendance,
		}

		_, err := UpdateStudentToDb(student, dbconfig, id)
		if err != nil {
			return err
		}
		return c.Render("updatestudent", fiber.Map{
			"Title": "Edit Existing Student",
		}, "layout")
	}
	return nil
}

func UpdateStudentToDb(student model.Student, dbconfig model.DbConfig, id string) (*model.Student, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.UserName, dbconfig.Password, dbconfig.DatabaseName)

	updatestatement := "UPDATE student SET firstname = $1, lastname = $2, emailid = $3, attendance = $4 WHERE id = $5"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	_, updateErr := db.Exec(updatestatement, student.FirstName, student.LastName, student.EmailId, student.Attendance, id)

	if updateErr != nil {
		return nil, updateErr
	}
	return &student, nil
}
