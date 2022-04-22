package main

import "fmt"

// Slice is a dynamic array, array not fix length.
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	showSlice(numbers)

	// leading remove
	numbers = numbers[1:len(numbers)]
	showSlice(numbers)
	numbers = numbers[1:len(numbers)]
	showSlice(numbers)

	// trailing remove
	numbers = numbers[0 : len(numbers)-1]
	showSlice(numbers)
	numbers = numbers[0 : len(numbers)-1]
	showSlice(numbers)
	fmt.Println()

	// remove specific index
	numbers = []int{1, 2, 3, 4, 5, 6}
	showSlice(numbers)
	numbers = removeIndex(numbers, 2)
	showSlice(numbers)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) // %v = print value
}

func removeIndex(array []int, index int) []int {
	// append() โดยปกติแล้ว จะเพิ่มได้ทีละ value
	// แต่ถ้าต้องการ append array จะต้องเพิ่ม ... ต่อท้าย array ตัวนั้นๆ ทำให้ได้แบบนี้
	return append(array[:index], array[index+1:]...)
}
