package main

import "fmt"

func main() {
	var numbers = map[int]string{1: "one", 2: "two", 3: "three"} // init map and assign value
	fmt.Println("show all value : ", numbers)

	key := 2
	fmt.Printf("show value by key = %d : %s\n", key, numbers[key])

	var colors = make(map[string]string) // *make() คล้ายๆกับการใช้ new ในภาษา Java
	colors["red"] = "#f00"
	colors["green"] = "#0f0"
	colors["blue"] = "#00f"
	fmt.Println("show all colors : ", colors)
	fmt.Println("show green code : ", colors["green"])

	// 2D Map
	var courses = make(map[string]map[string]int)
	courses["android"] = make(map[string]int)
	courses["android"]["code"] = 1001
	courses["android"]["price"] = 200

	courses["ios"] = map[string]int{"code": 1002, "price": 250}

	fmt.Println("\nshow all value : ", courses)
	fmt.Println("show course android price : ", courses["android"]["price"])
}
