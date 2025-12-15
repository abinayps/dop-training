package main

import "fmt"

// Project represents a project in the organization
type Project struct {
	ID          int
	Name        string
	Description string
	Status      string      // e.g., "Active", "Completed", "On Hold"
	TeamMembers []*Employee // Contains actual Employee pointers (dependency)
	Department  *Department // Project belongs to a Department (dependency)
}

// NewProject is a constructor function to create a new Project
func NewProject(id int, name, description, status string, department *Department) *Project {
	return &Project{
		ID:          id,
		Name:        name,
		Description: description,
		Status:      status,
		TeamMembers: []*Employee{},
		Department:  department,
	}
}

// DisplayProject prints project details
func (p Project) DisplayProject() {
	fmt.Println("=== Project Details ===")
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Description: %s\n", p.Description)
	fmt.Printf("Status: %s\n", p.Status)
	if p.Department != nil {
		fmt.Printf("Department: %s\n", p.Department.Name)
	}
	fmt.Printf("Team Size: %d members\n", len(p.TeamMembers))
	if len(p.TeamMembers) > 0 {
		fmt.Println("Team Members:")
		for _, emp := range p.TeamMembers {
			fmt.Printf("  - %s (%s)\n", emp.Name, emp.Position)
		}
	}
	fmt.Println("=======================")
}

// AssignEmployee assigns an employee to the project
func (p *Project) AssignEmployee(employee *Employee) {
	p.TeamMembers = append(p.TeamMembers, employee)
	fmt.Printf("Employee %s assigned to project %s\n", employee.Name, p.Name)
}

// RemoveEmployee removes an employee from the project
func (p *Project) RemoveEmployee(employeeID int) {
	for i, emp := range p.TeamMembers {
		if emp.ID == employeeID {
			// Remove the employee by slicing
			p.TeamMembers = append(p.TeamMembers[:i], p.TeamMembers[i+1:]...)
			fmt.Printf("Employee ID %d removed from project %s\n", employeeID, p.Name)
			return
		}
	}
	fmt.Printf("Employee ID %d not found in project %s\n", employeeID, p.Name)
}

// UpdateStatus updates the project status
func (p *Project) UpdateStatus(newStatus string) {
	p.Status = newStatus
	fmt.Printf("Project %s status updated to: %s\n", p.Name, newStatus)
}

// GetTeamSize returns the number of team members
func (p Project) GetTeamSize() int {
	return len(p.TeamMembers)
}
