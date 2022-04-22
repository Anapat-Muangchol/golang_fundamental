package main

import "fmt"

func main() {
	fn1()
	fn2("Test 2")
	fn3("Good morning", "Student", 1)

	const a, b = 2, 3
	fmt.Printf("%d + %d = %d\n", a, b, sum(a, b))

	x, y := swap(a, b)
	fmt.Printf("%d, %d => %d, %d\n", a, b, x, y)

	x, y = swapV2(a, b)
	fmt.Printf("%d, %d => %d, %d\n", a, b, x, y)

	_x, _y, title := swapV3(a, b)
	fmt.Printf("%d, %d => %d, %d | %s\n", a, b, _x, _y, title)
}

func fn1() {
	fmt.Println("Test")
}

func fn2(msg string) {
	fmt.Println(msg)
}

func fn3(title, subtitle string, version int) {
	fmt.Print(title + ", ")
	fmt.Println(subtitle)
	fmt.Printf("version %d\n", version)
}

func sum(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func swapV2(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func swapV3(a, b int) (int, int, string) {
	return b, a, "Result"
}
