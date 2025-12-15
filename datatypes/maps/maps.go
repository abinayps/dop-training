package maps

import "fmt"

func Maps() {
	var m map[int]int
	m = make(map[int]int)
	m[1] = 100
	m[2] = 200
	m[3] = 300
	fmt.Println("Map m:", m)
	fmt.Println("Length of map m:", len(m))
	m[2] = 250
	m[4] = 400
	m[4] = 450
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	d := make(map[int]string)
	d[1] = "Mango"
	d[2] = "Banana"
	fmt.Println("Map d:", d)
	fmt.Println("Map d[2]:", d[2])

	e := make(map[string]int)
	e["Apple"] = 1
	e["Banana"] = 2
	e["Banana"] = 3
	fmt.Println("Map e:", e)

	f := make(map[int][]int)
	f[1] = []int{10, 20, 30}
	f[2] = []int{40, 50, 60}
	fmt.Println("Map f:", f)

	g := make(map[string]int)
	g["Apple"] = 1

	fmt.Println("Value: ", d[g["Apple"]])

	// h := make(map[string]interface{})
	h := make(map[string]any)
	h["Name"] = "John"
	h["Age"] = 30
	h["IsStudent"] = false
	fmt.Println("Map h:", h)
}
