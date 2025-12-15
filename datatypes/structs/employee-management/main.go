package main

import "fmt"

// EmployeeManagementSystem manages all employees, departments, projects, and attendance
type EmployeeManagementSystem struct {
	Employees   []*Employee
	Departments []*Department
	Projects    []*Project
	Attendance  []*AttendanceRecord
}

// NewEmployeeManagementSystem is a constructor for the management system
func NewEmployeeManagementSystem() *EmployeeManagementSystem {
	return &EmployeeManagementSystem{
		Employees:   []*Employee{},
		Departments: []*Department{},
		Projects:    []*Project{},
		Attendance:  []*AttendanceRecord{},
	}
}

// AddEmployee adds a new employee to the system
func (ems *EmployeeManagementSystem) AddEmployee(emp *Employee) {
	ems.Employees = append(ems.Employees, emp)
	fmt.Printf("✓ Employee %s added successfully!\n", emp.Name)
}

// RemoveEmployee removes an employee by ID
func (ems *EmployeeManagementSystem) RemoveEmployee(employeeID int) {
	for i, emp := range ems.Employees {
		if emp.ID == employeeID {
			ems.Employees = append(ems.Employees[:i], ems.Employees[i+1:]...)
			fmt.Printf("✓ Employee ID %d removed successfully!\n", employeeID)
			return
		}
	}
	fmt.Printf("✗ Employee ID %d not found!\n", employeeID)
}

// SearchEmployeeByID finds an employee by ID
func (ems EmployeeManagementSystem) SearchEmployeeByID(employeeID int) *Employee {
	for _, emp := range ems.Employees {
		if emp.ID == employeeID {
			return emp
		}
	}
	return nil
}

// SearchEmployeeByName finds employees by name (exact match)
func (ems EmployeeManagementSystem) SearchEmployeeByName(name string) []*Employee {
	var results []*Employee
	for _, emp := range ems.Employees {
		if emp.Name == name {
			results = append(results, emp)
		}
	}
	return results
}

// FilterEmployeesByPosition returns all employees with a specific position
func (ems EmployeeManagementSystem) FilterEmployeesByPosition(position string) []*Employee {
	var results []*Employee
	for _, emp := range ems.Employees {
		if emp.Position == position {
			results = append(results, emp)
		}
	}
	return results
}

// FilterEmployeesBySalary returns employees with salary above a threshold
func (ems EmployeeManagementSystem) FilterEmployeesBySalary(minSalary float64) []*Employee {
	var results []*Employee
	for _, emp := range ems.Employees {
		if emp.Salary >= minSalary {
			results = append(results, emp)
		}
	}
	return results
}

// DisplayAllEmployees shows all employees in the system
func (ems EmployeeManagementSystem) DisplayAllEmployees() {
	fmt.Println("\n========== ALL EMPLOYEES ==========")
	if len(ems.Employees) == 0 {
		fmt.Println("No employees in the system.")
		return
	}
	for _, emp := range ems.Employees {
		fmt.Printf("ID: %d | Name: %s | Position: %s | Salary: $%.2f\n",
			emp.ID, emp.Name, emp.Position, emp.Salary)
	}
	fmt.Println("===================================\n")
}

func main() {
	fmt.Println("🏢 Employee Management System 🏢")
	fmt.Println("=================================\n")

	// Create the management system using constructor
	ems := NewEmployeeManagementSystem()

	// 1. CREATE EMPLOYEES using constructors
	fmt.Println("--- 1. Adding Employees ---")
	emp1 := NewEmployee(101, "John Doe", "john@company.com", "Software Engineer", 75000)
	emp2 := NewEmployee(102, "Jane Smith", "jane@company.com", "Manager", 95000)
	emp3 := NewEmployee(103, "Bob Johnson", "bob@company.com", "Developer", 70000)

	ems.AddEmployee(emp1)
	ems.AddEmployee(emp2)
	ems.AddEmployee(emp3)

	// 2. DISPLAY ALL EMPLOYEES
	ems.DisplayAllEmployees()

	// 3. CREATE DEPARTMENTS using constructors
	fmt.Println("\n--- 2. Creating Departments ---")
	dept1 := NewDepartment(1, "Engineering", "Building A")
	dept2 := NewDepartment(2, "Management", "Building B")

	ems.Departments = append(ems.Departments, dept1, dept2)
	dept1.DisplayDepartment()

	// 4. ASSIGN EMPLOYEES TO DEPARTMENT (showing dependency)
	fmt.Println("\n--- 3. Assigning Employees to Departments ---")
	dept1.AddEmployee(emp1) // Department now contains actual Employee
	dept1.AddEmployee(emp3)
	dept2.AddEmployee(emp2)

	fmt.Println("\nDepartment after adding employees:")
	dept1.DisplayDepartment()

	// 5. CREATE AND MANAGE PROJECTS (showing dependency with Department)
	fmt.Println("\n--- 4. Creating Projects ---")
	project1 := NewProject(1, "Website Redesign", "Redesign company website", "Active", dept1)

	ems.Projects = append(ems.Projects, project1)
	project1.DisplayProject()

	fmt.Println("\n--- 5. Assigning Employees to Projects ---")
	project1.AssignEmployee(emp1) // Project now contains actual Employees
	project1.AssignEmployee(emp3)

	fmt.Println("\nProject after adding employees:")
	project1.DisplayProject()

	// 6. UPDATE EMPLOYEE
	fmt.Println("\n--- 6. Updating Employee Information ---")
	emp1.UpdateEmployee("John Doe", "john.doe@company.com", "Senior Software Engineer", 85000)
	emp1.DisplayEmployee()

	// 7. SEARCH EMPLOYEES
	fmt.Println("\n--- 7. Searching Employees ---")
	foundEmp := ems.SearchEmployeeByID(102)
	if foundEmp != nil {
		fmt.Println("Employee found:")
		foundEmp.DisplayEmployee()
	}

	// 8. FILTER EMPLOYEES
	fmt.Println("\n--- 8. Filtering Employees by Position ---")
	developers := ems.FilterEmployeesByPosition("Developer")
	fmt.Printf("Developers found: %d\n", len(developers))
	for _, emp := range developers {
		fmt.Printf("  - %s: $%.2f\n", emp.Name, emp.Salary)
	}

	fmt.Println("\n--- 9. Filtering Employees by Salary ---")
	highEarners := ems.FilterEmployeesBySalary(80000)
	fmt.Printf("Employees earning >= $80,000: %d\n", len(highEarners))
	for _, emp := range highEarners {
		fmt.Printf("  - %s: $%.2f\n", emp.Name, emp.Salary)
	}

	// 9. ATTENDANCE MANAGEMENT (showing dependency with Employee)
	fmt.Println("\n--- 10. Managing Attendance ---")
	attendance1 := NewAttendanceRecord(emp1, "Present") // Attendance contains actual Employee
	attendance1.MarkCheckOut()
	attendance1.DisplayAttendance()

	attendance2 := NewAttendanceRecord(emp2, "Present")
	ems.Attendance = append(ems.Attendance, attendance1, attendance2)

	// 10. REMOVE EMPLOYEE FROM PROJECT
	fmt.Println("\n--- 11. Removing Employee from Project ---")
	project1.RemoveEmployee(103)
	project1.DisplayProject()

	// 11. REMOVE EMPLOYEE FROM SYSTEM
	fmt.Println("\n--- 12. Removing Employee from System ---")
	ems.RemoveEmployee(103)
	ems.DisplayAllEmployees()

	// 12. PROJECT STATUS UPDATE
	fmt.Println("\n--- 13. Updating Project Status ---")
	project1.UpdateStatus("Completed")

	fmt.Println("\n✅ Employee Management System Demo Complete!")
	fmt.Println("\n📝 Key Points Demonstrated:")
	fmt.Println("  ✓ Each struct has its own constructor (New...)")
	fmt.Println("  ✓ Department contains Employee pointers (dependency)")
	fmt.Println("  ✓ Project contains Employee pointers and Department pointer (dependency)")
	fmt.Println("  ✓ AttendanceRecord contains Employee pointer (dependency)")
	fmt.Println("  ✓ All structs work together showing real relationships!")
}
