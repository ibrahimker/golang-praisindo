package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var employees = []Employee{
	{
		ID:       1,
		Name:     "Airell",
		Age:      23,
		Division: "IT",
	},
	{
		ID:       2,
		Name:     "Nanda",
		Age:      23,
		Division: "Finance",
	},
	{
		ID:       3,
		Name:     "Mailo",
		Age:      20,
		Division: "IT",
	},
}

const PORT = ":8080"

func main() {
	http.HandleFunc("/employees", getEmployees)
	http.ListenAndServe(PORT, nil)
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(employees)
	case "POST":
		newEmployee := Employee{
			ID:       len(employees) + 1,
			Name:     r.FormValue("name"),
			Division: r.FormValue("division"),
		}
		age, _ := strconv.Atoi(r.FormValue("age"))
		newEmployee.Age = age
		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
