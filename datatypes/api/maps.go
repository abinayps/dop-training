package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	City      string `json:"city"`
	Email     string `json:"email"`
}

var UserMap = make(map[string]UserDetails)

func CreateUserDetails(c *gin.Context) {
	var user UserDetails
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation Error",
		})
		return
	}

	UserMap[user.Email] = user

	c.JSON(http.StatusCreated, gin.H{
		"message": "User details created successfully",
		"user":    user,
	})
}

type GetUserDetailsRequest struct {
	Email string `uri:"email"`
}

func GetUserDetails(c *gin.Context) {
	var req GetUserDetailsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation Error",
		})
		return
	}

	user, ok := UserMap[req.Email]
	if ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "User details retrieved successfully",
			"success": true,
			"user":    user,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"success": false,
		})
	}
}

func ListAllUserDetails(c *gin.Context) {
	if len(UserMap) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"message": "No users found",
			"users":   []UserDetails{},
		})
		return
	}
	users := make([]UserDetails, 0, len(UserMap))
	for _, user := range UserMap {
		users = append(users, user)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User details retrieved successfully",
		"success": true,
		"users":   users,
	})
}

type DeleteUserDetailsRequest struct {
	Email string `form:"email"`
}

func DeleteUserDetails(c *gin.Context) {
	var req DeleteUserDetailsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation Error",
		})
		return
	}
	_, ok := UserMap[req.Email]
	if ok {
		delete(UserMap, req.Email)
		c.JSON(http.StatusOK, gin.H{
			"message": "User deleted successfully",
			"success": true,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"success": false,
		})
	}
}

type UpdateUserDetailsRequest struct {
	Email     string `uri:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func UpdateUserDetails(c *gin.Context) {
	var req UpdateUserDetailsRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation Error",
		})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation Error",
		})
		return
	}
	_, ok := UserMap[req.Email]
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
			"success": false,
		})
		return
	}
	UserMap[req.Email] = UserDetails{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		City:      req.City,
		Age:       req.Age,
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User details updated successfully",
		"success": true,
		"user":    UserMap[req.Email],
	})
}

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
