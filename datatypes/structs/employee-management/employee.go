package main

import "fmt"

// Employee represents an employee in the system
type Employee struct {
	ID       int
	Name     string
	Email    string
	Position string
	Salary   float64
}

// NewEmployee is a constructor function to create a new Employee
func NewEmployee(id int, name, email, position string, salary float64) *Employee {
	return &Employee{
		ID:       id,
		Name:     name,
		Email:    email,
		Position: position,
		Salary:   salary,
	}
}

// DisplayEmployee prints employee details
func (e Employee) DisplayEmployee() {
	fmt.Println("=== Employee Details ===")
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Email: %s\n", e.Email)
	fmt.Printf("Position: %s\n", e.Position)
	fmt.Printf("Salary: $%.2f\n", e.Salary)
	fmt.Println("========================")
}

// UpdateEmployee updates employee information
func (e *Employee) UpdateEmployee(name, email, position string, salary float64) {
	if name != "" {
		e.Name = name
	}
	if email != "" {
		e.Email = email
	}
	if position != "" {
		e.Position = position
	}
	if salary > 0 {
		e.Salary = salary
	}
	fmt.Println("Employee updated successfully!")
}
