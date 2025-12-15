package variables

import (
	"fmt"
)

// 1) Return user's first name, last name, age, city, and email as multiple return values.
func GetUserFields(firstName, lastName string, age int, city, email string) (string, string, int, string, string) {
	return firstName, lastName, age, city, email
}

// 2) Validate age; return a boolean and an error if invalid (<0 or >120).
func ValidateAge(age int) (bool, error) {
	if age < 0 || age > 120 {
		return false, fmt.Errorf("age %d is out of range [0, 120]", age)
	}
	return true, nil
}

// 3) Return user's full name and the total number of characters in it (Unicode-aware).
func FullNameAndLength(firstName, lastName string) (string, int) {
	// first := strings.TrimSpace(firstName)
	// last := strings.TrimSpace(lastName)

	var fullName string
	switch {
	case firstName != "" && lastName != "":
		fullName = firstName + " " + lastName
	case firstName != "":
		fullName = firstName
	default:
		fullName = lastName
	}

	return fullName, len(fullName)
}

// 4) Convert user's age from int to string and return both values.
func AgeIntAndString(age int) (int, string) {
	return age, fmt.Sprintf("%d", age)
}

// 5) Demonstrate usage of all the above functions.
func DemonstrateUserFunctions() {
	// Example values
	firstName := "John"
	lastName := "Doe"
	age := 25
	city := "New York"
	email := "john.doe@example.com"

	// Call GetUserFields
	fn, ln, a, c, e := GetUserFields(firstName, lastName, age, city, email)
	fmt.Printf("User Fields: %s, %s, %d, %s, %s\n", fn, ln, a, c, e)

	// Call ValidateAge
	isValid, err := ValidateAge(age)
	if err != nil {
		fmt.Printf("Age validation failed: %v\n", err)
	} else {
		fmt.Printf("Age is valid: %v\n", isValid)
	}

	// Call FullNameAndLength
	fullName, nameLength := FullNameAndLength(firstName, lastName)
	fmt.Printf("Full Name: %s, Length: %d\n", fullName, nameLength)

	// Call AgeIntAndString
	ageInt, ageStr := AgeIntAndString(age)
	fmt.Printf("Age (int): %d, Age (string): %s\n", ageInt, ageStr)
}
