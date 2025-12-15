package slices

import "fmt"

func Slices() {
	// var b [5]float32
	// b[0] = 10.5
	// fmt.Println("Array b:", b)
	// var a = make([]int, 5, 10)
	// fmt.Println("Slice a:", a)
	// a = append(a, 10)
	// fmt.Println("Slice a after appending 10:", a)
	// fmt.Println("Length of slice a:", len(a))
	// fmt.Println("Capacity of slice a:", cap(a))
	// a = make([]int, 5, 10)
	// fmt.Println("Slice a after make:", a)
	// fmt.Println("Length of slice a:", len(a))
	// fmt.Println("Capacity of slice a:", cap(a))
	// a[0] = 10
	// fmt.Println("Slice a after assigning first element:", a)
	// a = append(a, 60, 70, 80, 90, 100)
	// fmt.Println("Slice a after assigning sixth element:", a)
	// fmt.Println("Length of slice a after append:", len(a))
	// fmt.Println("Capacity of slice a after append:", cap(a))
	// a = append(a, 110)
	// fmt.Println("Slice a after assigning eleventh element:", a)
	// fmt.Println("Length of slice a after appending eleventh element:", len(a))
	// fmt.Println("Capacity of slice a after appending eleventh element:", cap(a))

	// var d = make([]int, 5)
	// d = append(d, 1, 2, 3)
	// fmt.Println("Slice d after appending elements:", d)
	// fmt.Println("Length of slice d after appending elements:", len(d))
	// fmt.Println("Capacity of slice d after appending elements:", cap(d))

	var a = []int{0, 1, 2, 3, 4, 5}

	b := a[:3]  //index 0,1,2 (3rd element)
	c := a[3:]  //index 3,4,5 (from 4th element to end)
	d := a[2:5] //index 2,3,4    [included:excluded]
	e := []int{1, 2, 3, 4, 5, 6}
	f := e[:3]
	g := e[3:]
	h := e[1:4]
	i := e[:6]

	fmt.Println("Original slice a:", a)
	fmt.Println("Slice b:", b)
	fmt.Println("Slice c:", c)
	fmt.Println("Slice d:", d)
	fmt.Println("Original slice e:", e)
	fmt.Println("Slice f:", f)
	fmt.Println("Slice g:", g)
	fmt.Println("Slice h:", h)
	fmt.Println("Slice i:", i)
}
