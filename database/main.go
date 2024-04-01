package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connect to database")

	// Create
	if err := CreateEmployee(Employee{
		FullName: "ibam",
		Email:    "ibrahimker@gmail.com",
		Age:      30,
		Division: "sales",
	}); err != nil {
		log.Println(err)
	}

	employees, err := GetAllEmployee()
	if err == nil {
		for _, employee := range employees {
			fmt.Println("employee", employee)
		}
	}
}

func CreateEmployee(employee Employee) error {
	queryString := "INSERT INTO employees (full_name, email, age, division) " +
		"VALUES ($1,$2,$3,$4) "
	_, err := db.Exec(queryString, employee.FullName, employee.Email, employee.Age, employee.Division)
	if err != nil {
		return err
	}

	return nil
}

func GetAllEmployee() ([]Employee, error) {
	queryString := "SELECT id, full_name, email, age, division FROM employees LIMIT 10"
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Employee
	for rows.Next() {
		var employee Employee
		if err := rows.Scan(
			&employee.ID, &employee.FullName, &employee.Email, &employee.Age, &employee.Division,
		); err != nil {
			log.Println(err)
			continue
		}
		res = append(res, employee)
	}

	return res, nil
}
