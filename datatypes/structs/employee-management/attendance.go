package main

import (
	"fmt"
	"time"
)

// AttendanceRecord represents daily attendance of an employee
type AttendanceRecord struct {
	Employee *Employee // Contains actual Employee pointer (dependency)
	Date     string
	Status   string // e.g., "Present", "Absent", "Leave"
	CheckIn  string
	CheckOut string
}

// NewAttendanceRecord is a constructor function to create a new AttendanceRecord
func NewAttendanceRecord(employee *Employee, status string) *AttendanceRecord {
	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")

	return &AttendanceRecord{
		Employee: employee,
		Date:     currentDate,
		Status:   status,
		CheckIn:  currentTime,
		CheckOut: "",
	}
}

// DisplayAttendance prints attendance record details
func (a AttendanceRecord) DisplayAttendance() {
	fmt.Println("=== Attendance Record ===")
	if a.Employee != nil {
		fmt.Printf("Employee: %s (ID: %d)\n", a.Employee.Name, a.Employee.ID)
		fmt.Printf("Position: %s\n", a.Employee.Position)
	}
	fmt.Printf("Date: %s\n", a.Date)
	fmt.Printf("Status: %s\n", a.Status)
	fmt.Printf("Check-In: %s\n", a.CheckIn)
	fmt.Printf("Check-Out: %s\n", a.CheckOut)
	fmt.Println("=========================")
}

// MarkCheckOut records the check-out time
func (a *AttendanceRecord) MarkCheckOut() {
	currentTime := time.Now().Format("15:04:05")
	a.CheckOut = currentTime
	if a.Employee != nil {
		fmt.Printf("Check-out marked at %s for %s\n", currentTime, a.Employee.Name)
	}
}

// GetWorkingHours calculates working hours (simplified version)
func (a AttendanceRecord) GetWorkingHours() string {
	if a.CheckIn == "" || a.CheckOut == "" {
		return "N/A"
	}
	return "Working hours calculated based on check-in and check-out"
}
