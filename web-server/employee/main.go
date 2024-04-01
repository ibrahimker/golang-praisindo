package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	ID       int    `json:"id`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Division string `json:"division`
	Sekolah  []string
}

var employees = []Employee{
	{
		ID:       1,
		Name:     "Airell",
		Age:      23,
		Division: "IT",
		Sekolah:  []string{"SD 1", "SMP 2"},
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

type ErrMessage struct {
	Error string `json:"error"`
}

const PORT = ":8080"

func main() {
	http.HandleFunc("/employees", handleEmployees)
	http.ListenAndServe(PORT, nil)
}

func handleEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(employees)
	case "POST":
		// if read from form url
		// newEmployee := Employee{
		// 	ID:       len(employees) + 1,
		// 	Name:     r.FormValue("name"),
		// 	Division: r.FormValue("division"),
		// }
		// age, _ := strconv.Atoi(r.FormValue("age"))
		// newEmployee.Age = age

		// if read from json
		newEmployee := Employee{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&newEmployee); err != nil {
			log.Println(err)
			// data, _ := json.Marshal(ErrMessage{Error: err.Error()})
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
		newEmployee.ID = len(employees) + 1

		employees = append(employees, newEmployee)
		json.NewEncoder(w).Encode(newEmployee)
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
	}
}
