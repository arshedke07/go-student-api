package service

import (
	"database/sql"
	"fmt"
	"studentmgt/model"
)

func FindStudentById(dbconfig model.DbConfig, id string) (*model.Student, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.UserName, dbconfig.Password, dbconfig.DatabaseName)

	selectstatement := "SELECT id, firstname, lastname, emailid, attendance FROM student WHERE id = $1"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(selectstatement, id)
	student := model.Student{}
	scanErr := row.Scan(&student.Id, &student.FirstName, &student.LastName, &student.EmailId, &student.Attendance)
	if scanErr != nil {
		return nil, err
	}
	return &student, nil
}

func FindAllStudent(dbconfig model.DbConfig) (*[]model.Student, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbconfig.Host, dbconfig.Port, dbconfig.UserName, dbconfig.Password, dbconfig.DatabaseName)

	selectstatement := "SELECT id, firstname, lastname, emailid, attendance FROM student"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(selectstatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []model.Student{}
	for rows.Next() {
		student := model.Student{}
		scanErr := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.EmailId, &student.Attendance)
		if scanErr != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return &students, nil
}
