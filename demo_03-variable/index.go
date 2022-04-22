package main

import (
	"fmt"
)

var count = 0

func main() {
	fmt.Println("Begin")

	// Explicit Declaration (ระบุ Type ของตัวแปร)
	var temp1 int = 0
	temp1 = 10
	var temp2 string = "hello"
	var temp3 bool = true
	const temp4 int = 0
	// temp4 = 10

	// Implicit Declaration (ไม่ระบุ Type ของตัวแปร)
	temp5 := 0
	temp6 := "Hello"

	fmt.Println(temp1)
	fmt.Println(temp2)
	fmt.Println(temp3)
	fmt.Println(temp4)

	fmt.Println(temp5)
	fmt.Println(temp6)

	plus()
	fmt.Println(count)
}

func plus() {
	count++
	count++
	count++
	count += 10
}
