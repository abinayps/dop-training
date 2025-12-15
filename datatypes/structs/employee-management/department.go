package main

import "fmt"

// Department represents a department in the organization
type Department struct {
	ID        int
	Name      string
	Location  string
	Employees []*Employee // Contains actual Employee pointers (dependency)
}

// NewDepartment is a constructor function to create a new Department
func NewDepartment(id int, name, location string) *Department {
	return &Department{
		ID:        id,
		Name:      name,
		Location:  location,
		Employees: []*Employee{}, // Initialize empty slice
	}
}

// DisplayDepartment prints department details
func (d Department) DisplayDepartment() {
	fmt.Println("=== Department Details ===")
	fmt.Printf("ID: %d\n", d.ID)
	fmt.Printf("Name: %s\n", d.Name)
	fmt.Printf("Location: %s\n", d.Location)
	fmt.Printf("Number of Employees: %d\n", len(d.Employees))
	if len(d.Employees) > 0 {
		fmt.Println("Employees:")
		for _, emp := range d.Employees {
			fmt.Printf("  - %s (%s)\n", emp.Name, emp.Position)
		}
	}
	fmt.Println("==========================")
}

// AddEmployee adds an employee to the department
func (d *Department) AddEmployee(employee *Employee) {
	d.Employees = append(d.Employees, employee)
	fmt.Printf("Employee %s added to department %s\n", employee.Name, d.Name)
}

// RemoveEmployee removes an employee from the department
func (d *Department) RemoveEmployee(employeeID int) {
	for i, emp := range d.Employees {
		if emp.ID == employeeID {
			// Remove the employee by slicing
			d.Employees = append(d.Employees[:i], d.Employees[i+1:]...)
			fmt.Printf("Employee ID %d removed from department %s\n", employeeID, d.Name)
			return
		}
	}
	fmt.Printf("Employee ID %d not found in department %s\n", employeeID, d.Name)
}

// GetEmployeeCount returns the number of employees in the department
func (d Department) GetEmployeeCount() int {
	return len(d.Employees)
}
