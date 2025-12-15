package slices

import "fmt"

// 1. Create a slice of length 10 and capacity 20. Print its length and capacity.
func question1() {
	s := make([]int, 10, 20)
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
}

// 2. Append 15 integers to an empty slice and print the slice, its length, and capacity after each append operation.
func question2() {
	s := []int{}
	for i := 1; i <= 15; i++ {
		s = append(s, i)
		fmt.Printf("After appending %d: slice=%v, length=%d, capacity=%d\n", i, s, len(s), cap(s))
	}
	// b := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// for i, v := range b {
	// 	s = append(s, v)
	// 	fmt.Printf("After appending index %d, value=%d, slice=%v, length=%d, capacity=%d\n", i, v, s, len(s), cap(s))
	// }
}

// 3. Write a function that takes a slice of integers and returns a new slice containing only the even numbers from the original slice.
func filterEven(nums []int) []int {
	evens := []int{}
	for _, n := range nums {
		if n%2 == 0 {
			evens = append(evens, n)
		}
	}
	return evens
}

func SliceExamples() {
	// question1()
	// question2()
	fmt.Println("Even numbers:", filterEven([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}

// Write a function that takes a slice of integers and returns a new slice with the elements reversed.
func ReverseSlice(nums []int) []int {
	reversed := make([]int, 0, len(nums))
	// Loop in reverse order
	for i := len(nums) - 1; i >= 0; i-- {
		reversed = append(reversed, nums[i])
	}
	return reversed
}

// Given two slices, create a new slice that combines elements from both.
// a := []int{1, 2, 3}
// b := []int{4, 5, 6}
func CombineSlices(a, b []int) []int {
	// Preallocate result with the total length
	result := make([]int, 0, len(a)+len(b))

	// Append elements from a using a loop
	for i := 0; i < len(a); i++ {
		result = append(result, a[i])
	}

	// Append elements from b using a loop
	for i := 0; i < len(b); i++ {
		result = append(result, b[i])
	}

	return result
}

// Write a function that inserts a value at a given index in a slice of integers.
// The function must return a new slice with the value inserted and all elements shifted accordingly.
// nums := []int{10, 20, 30, 40, 50}
// index := 2
// value := 99
func InsertAt(nums []int, index int, value int) []int {
	if index < 0 || index > len(nums) {
		// If index is out of bounds, return the original slice
		return nums
	}

	result := make([]int, 0, len(nums)+1)

	// Add elements before the index
	for i := 0; i < index; i++ {
		result = append(result, nums[i])
	}

	// Add the new value at the index
	result = append(result, value)

	// Append remaining items
	for i := index; i < len(nums); i++ {
		result = append(result, nums[i])
	}

	return result
}

// Write a function that divides a slice into multiple smaller slices (chunks) of a given size.
// The last chunk may be smaller if elements don’t fit evenly.
// nums := []int{10, 20, 30, 40, 50, 60, 70}
// chunkSize := 3
func SplitFixed(nums []int) ([]int, []int, []int) {
	s1 := nums[:3]
	s2 := nums[3:6]
	s3 := nums[6:]
	s3[0] = 0
	fmt.Println("Nums: ", nums)
	return s1, s2, s3
}

// func SplitFixed(nums []int) [][]int {
// 	chunkSize := 3
// 	result := make([][]int, 0)

// 	for i := 0; i < len(nums); i += chunkSize {
// 		end := i + chunkSize
// 		if end > len(nums) {
// 			end = len(nums)
// 		}
// 		result = append(result, nums[i:end])
// 	}

// 	return result
// }
