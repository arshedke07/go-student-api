package routes

import (
	"database/sql"
	"fmt"
	"studentmgt/model"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func AddStudent(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("addstudent", fiber.Map{
			"Title": "Add Student",
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

		dbconfig := model.DbConfig{
			Host:         "localhost",
			Port:         "5432",
			UserName:     "postgres",
			Password:     "1234",
			DatabaseName: "studentmanagement",
		}
		_, err := AddStudentToDb(student, dbconfig)
		if err != nil {
			return err
		}
		return c.Render("addstudent", fiber.Map{
			"Title": "Add Student",
		}, "layout")
	}
	return nil
}

func AddStudentToDb(student model.Student, dbconfig model.DbConfig) (*model.Student, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.UserName, dbconfig.Password, dbconfig.DatabaseName)

	// insertstatement := "INSERT INTO student(firstname, lastname, emailid) VALUES($1,$2,$3)"
	insertstatement := "INSERT INTO student(firstname, lastname, emailid, attendance) VALUES($1,$2,$3, $4)"

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	defer db.Close()
	_, inserterr := db.Exec(insertstatement, student.FirstName, student.LastName, student.EmailId, student.Attendance)
	if inserterr != nil {
		return nil, inserterr
	}
	return &student, nil
}
