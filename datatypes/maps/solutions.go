package maps

import "fmt"

// Create a map where keys are student names and values are their scores.
// Write a function to:
// -add a student
// -update a student's score
// -print all students and scores

func addStudent(scores map[string]int, name string, score int) {
	_, exists := scores[name]
	if exists {
		fmt.Printf("student %q already exists\n", name)
		fmt.Println("---", scores)
		return
	}
	scores[name] = score
}

func updateScore(scores map[string]int, name string, score int) {
	_, exists := scores[name]
	if !exists {
		fmt.Printf("student %q not found\n", name)
		return
	}
	scores[name] = score
}

func printAll(scores map[string]int) {
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}
}

func StudentMap() {
	scores := make(map[string]int)

	addStudent(scores, "Alice", 85)
	addStudent(scores, "Bob", 90)
	updateScore(scores, "Alice", 88)
	updateScore(scores, "Charlie", 70) // not found

	printAll(scores)
}
