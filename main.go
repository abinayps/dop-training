package main

import (
	"log"
	"training/datatypes/api"
)

func main() {
	// variables.Variables()
	// variables.DemonstrateUserFunctions()
	// arrays.Arrays()
	// arrays.LoopsSolutions()
	// arrays.ArraysSolutions()
	// slices.Slices()
	// slices.SliceExamples()
	// maps.Maps()

	// fmt.Println(slices.ReverseSlice([]int{1, 2, 3, 4, 5}))
	// fmt.Println(slices.CombineSlices([]int{1, 2, 3}, []int{4, 5, 6}))
	// fmt.Println(slices.InsertAt([]int{10, 20, 30, 40, 50}, 2, 99))
	// fmt.Println(slices.SplitFixed([]int{10, 20, 30, 40, 50, 60, 70}))

	// maps.StudentMap()

	// structs.StudentDemo()
	// pointers.Pointers()

	// structs.StructExamples()
	// err:=api.StartServerGin(":8080")
	// if err!=nil{
	// 	log.Fatalf("server error: %v", err)
	// }

	if err := api.StartServerGin(":8080"); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
