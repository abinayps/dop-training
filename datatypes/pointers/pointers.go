package pointers

import "fmt"

func changeRef(x *int) {
	*x = 10
	fmt.Println("x Reference: ", *x)
}
func changeValue(x int) {
	x = 10
	fmt.Println("x value: ", x)
}

func Pointers() {
	a := 5

	changeValue(a)
	fmt.Println("a value:", a)

	changeRef(&a)
	fmt.Println("a  ref: ", a)

}
