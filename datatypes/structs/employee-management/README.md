# Employee Management System

A beginner-friendly Go project demonstrating struct composition, dependencies, and object-oriented programming concepts.

## 📁 Project Structure

```
employee-management/
├── employee.go       # Employee struct with constructor
├── department.go     # Department struct (contains Employees)
├── project.go        # Project struct (contains Employees & Department)
├── attendance.go     # AttendanceRecord struct (contains Employee)
├── main.go          # Main program demonstrating all features
└── README.md        # This file
```

## 🎯 Key Concepts Demonstrated

### 1. **Constructors for Each Struct**
Each struct has a constructor function (e.g., `NewEmployee()`, `NewDepartment()`) that creates and initializes instances properly.

```go
emp := NewEmployee(101, "John Doe", "john@company.com", "Software Engineer", 75000)
```

### 2. **Struct Dependencies (Composition)**

#### Employee (Base struct)
- Independent struct with basic employee information
- No dependencies on other structs

#### Department (Contains Employees)
```go
type Department struct {
    ID        int
    Name      string
    Location  string
    Employees []*Employee  // ← Contains Employee pointers
}
```

#### Project (Contains Employees & Department)
```go
type Project struct {
    ID          int
    Name        string
    Description string
    Status      string
    TeamMembers []*Employee  // ← Contains Employee pointers
    Department  *Department  // ← Belongs to a Department
}
```

#### AttendanceRecord (Contains Employee)
```go
type AttendanceRecord struct {
    Employee *Employee  // ← Contains Employee pointer
    Date     string
    Status   string
    CheckIn  string
    CheckOut string
}
```

## 🚀 Features

### Employee Operations
- ✅ Create employees using constructor
- ✅ Update employee information
- ✅ Display employee details
- ✅ Search by ID or name
- ✅ Filter by position or salary

### Department Operations
- ✅ Create departments using constructor
- ✅ Add employees to department (actual Employee objects)
- ✅ Remove employees from department
- ✅ Display department with employee list
- ✅ Get employee count

### Project Operations
- ✅ Create projects with department assignment
- ✅ Assign employees to project team
- ✅ Remove employees from project
- ✅ Update project status
- ✅ Display project with team members

### Attendance Operations
- ✅ Mark attendance for an employee
- ✅ Record check-in and check-out times
- ✅ Display attendance with employee details

## 💡 Running the Program

```bash
# Navigate to the directory
cd datatypes/structs/employee-management

# Run the program
go run .
```

## 📚 Learning Points

1. **Constructors**: Each struct has a `New...()` function to create instances
2. **Pointers**: Using `*Employee` instead of `Employee` for efficient memory usage
3. **Composition**: Structs contain other structs showing real-world relationships
4. **Methods**: Each struct has methods that operate on its data
5. **Slices of Pointers**: `[]*Employee` allows multiple references to same data
6. **Encapsulation**: Data and methods are grouped together

## 🎓 For Beginners

This code is intentionally kept simple:
- Clear variable names
- Simple logic without complex algorithms
- Comments explaining each section
- Step-by-step demonstration in main()
- Real-world example that's easy to understand

## 📝 Code Highlights

### Before (Independent Structs)
```go
type Department struct {
    EmployeeIDs []int  // Just IDs - no actual relationship
}
```

### After (Dependent Structs)
```go
type Department struct {
    Employees []*Employee  // Contains actual Employee objects!
}
```

This shows **real relationships** between data structures, just like in real applications!
