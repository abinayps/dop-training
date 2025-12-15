package structs

import "fmt"

func GetUserFields(firstName, lastName string, age int, city, email string) (string, string, int, string, string) {
	return firstName, lastName, age, city, email
}

type student struct {
	FirstName string
	LastName  string
	Age       int
	City      string
	Email     string
	Marks     float64
	Cars      Car
}

func CreateStudent(firstName, lastName string, age int, city,
	email string, marks float64, cars Car) *student {
	newStudent := student{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		City:      city,
		Email:     email,
		Marks:     marks,
		Cars:      cars,
	}
	return &newStudent
}

type Car struct {
	Name  string
	Model string
	Year  int
}

func CreateCar(name, model string, year int) Car {
	return Car{
		Name:  name,
		Model: model,
		Year:  year,
	}
}

func CreateStudents(s student, r student) []student {
	//stringSlice := []string{"one", "two"}
	newSlice := []student{s, r}
	return newSlice
}

func UpdateStudentMarks(s *student, newMarks float64) {
	s.Marks = newMarks
	fmt.Println("Updated Marks inside function:", s.Marks)
}

func (s *student) UpdateStudentMarks(newMarks float64) {
	s.Marks = newMarks
	fmt.Println("Updated Marks inside function:", s.Marks)
}

func (s *student) GetStudentName() {
	fmt.Println("First Name:", s.FirstName)
	fmt.Println("Last Name:", s.LastName)
}

func GetStudentDetails(s student) {
	// println("Student Details:")
	fmt.Println("First Name:", s.FirstName)
	fmt.Println("Last Name:", s.LastName)
	fmt.Println("Age:", s.Age)
	fmt.Println("City:", s.City)
	fmt.Println("Email:", s.Email)
	fmt.Println("Marks:", s.Marks)
	fmt.Println("-----------------------")
}

func StudentDemo() {
	// var s student = student{
	// 	FirstName: "Ram",
	// 	LastName:  "Sharma",
	// 	Age:       20,
	// 	City:      "New York",
	// 	Email:     "ram@example.com",
	// 	Marks:     85.5,
	// }

	c := CreateCar("Toyota", "Camry", 2020)

	r := CreateStudent("Sita", "Verma", 22, "Los Angeles",
		"sita@example.com", 90.0, c)

	// s := CreateStudent("Ram", "Sharma", 20, "New York",
	// 	"example.com", 90.0, c)

	UpdateStudentMarks(r, 95.0)
	r.UpdateStudentMarks(97.0)
	r.GetStudentName()

	fmt.Println("Marks after UpdateStudentMarks call:", r.Cars.Name)

	//  r := student{
	// 	FirstName: "Ram",
	// 	LastName:  "Sharma",
	// 	Age:       20,
	// 	City:      "New York",
	// 	Email:     "ram@example.com",
	// 	Marks:     85.5,
	// }

	// var r student
	// r.FirstName = "Sita"
	// r.LastName = "Verma"
	// r.Age = 22
	// r.City = "Los Angeles"
	// r.Email = "sita@example.com"
	// r.Marks = 90.0

	// // var s1 student = student{"Ram", "Sharma", 20, "New York", "ram@example.com", 85.5}
	// // fmt.Println("Empty Student Struct:", s1)
	// fmt.Println(CreateStudents(s, r))
	// s1 = CreateStudent("Ram", "Sharma", 20, "New York", "ram@example.com", 85.5)
	// s1.GetStudentDetails()
	// s1.UpdateMarks(90.0)
	// s1.GetStudentDetails()
	// UpdateStudentMarks(s1, 95.0)
	// s1.GetStudentDetails()
}

// --------------------------------------------------------------------------------

func (S student) GetStudentDetails() {
	// println("Student Details:")
	fmt.Println("First Name:", S.FirstName)
	fmt.Println("Last Name:", S.LastName)
	fmt.Println("Age:", S.Age)
	fmt.Println("City:", S.City)
	fmt.Println("Email:", S.Email)
	fmt.Println("Marks:", S.Marks)
	fmt.Println("-----------------------")
}

func (S student) UpdateMarks(newMarks float64) {
	S.Marks = newMarks
}

// -------------------------------------------------------------------------------
