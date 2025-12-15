package arrays

import "fmt"

// 4. Find Maximum value in an array of integers
func FindMax(nums []int) int {
	if len(nums) == 0 {
		fmt.Println("empty array")
		return 0
	}
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

// 5. Print each character of "Hello" on a new line
func PrintChars(s string) {
	for _, ch := range s {
		fmt.Println(string(ch))
	}
}

// 6. Count how many times 2 appears in the array
func CountTwos(arr []int) int {
	count := 0
	for _, v := range arr {
		if v == 2 {
			count++
			// count=count+1
		}
	}
	return count
}

// 7. Return the first odd value in the array using break
func FirstOdd(arr []int) (int, bool) {
	for _, v := range arr {
		if v%2 != 0 {
			// break
			return v, true
		}
	}

	return 0, false // no odd value found
}

// 8. Print all odd values in the array using continue
func PrintAllOdds(arr []int) {
	for _, v := range arr {
		if v%2 == 0 {
			continue
		}
		fmt.Println(v)
	}
}

// Sum all numbers from 1 to 100 and return the total sum.
func sumNumbers() int {
	total := 0
	for i := 1; i <= 100; i++ {
		total += i
	}
	return total
}

// Print all even numbers from 1 to 100 using a loop.
func printEvenNumbers() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Print(" ", i)
		}
	}
}

// Print all odd numbers from 1 to 100 using a loop.
func printOddNumbers() {
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			fmt.Print(" ", i)
		}
	}
}

func LoopsSolutions() {
	// Call the functions to demonstrate their functionality
	fmt.Println("Sum of numbers from 1 to 100:", sumNumbers())
	fmt.Println("")
	fmt.Print("Even numbers from 1 to 100: ")
	printEvenNumbers()
	fmt.Println("")
	fmt.Println("")
	fmt.Print("Odd numbers from 1 to 100: ")
	printOddNumbers()
}

// ArraysSolutions demonstrates the latest array-related functions (Q4-Q8)
func ArraysSolutions() {
	// 4. Find Maximum value in an array of integers
	nums := []int{5, 3, 9, 1, 7, 2}
	fmt.Println("Maximum value:", FindMax(nums))

	// 5. Print each character of "Hello" on a new line
	fmt.Println("Characters in 'Hello':")
	PrintChars("Hello")

	// 6. Count how many times 2 appears in the array
	arr := []int{1, 2, 3, 4, 5, 2, 2, 1, 3, 2}
	fmt.Println("Count of 2:", CountTwos(arr))

	// 7. Return the first odd value in the array using break
	arr2 := []int{2, 4, 6, 8, 9, 10, 12}
	if odd, ok := FirstOdd(arr2); ok {
		fmt.Println("First odd value:", odd)
	} else {
		fmt.Println("No odd value found")
	}

	// 8. Print all odd values in the array using continue
	fmt.Println("Odd values:")
	arr3 := []int{2, 4, 6, 8, 9, 10, 11, 13, 14}
	PrintAllOdds(arr3)
}
