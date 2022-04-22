package main

import "fmt"

// Slice is a dynamic array, array not fix length.
func main() {
	// *** make() คล้ายๆกับการใช้ new ในภาษา Java
	numbers1 := make([]int, 3, 5) // 3 len, 5 cap -> เพิ่มค่า default ไปแล้ว 3, แต่ความจุทั้งหมดคือ 5
	showSlice(numbers1)
	numbers1 = append(numbers1, 0)
	numbers1 = append(numbers1, 1)
	showSlice(numbers1)
	numbers1 = append(numbers1, 1) // ถ้า append เกิน cap จะเพิ่ม cap ไปอีก 1 เท่า (init cap * 2)
	showSlice(numbers1)

	fmt.Println()

	var numbers2 []int // ประกาศ slice โดยที่ไม่ระบุ len และ cap ดังนั้น default จะเป็น 0 ทั้งหมด
	showSlice(numbers2)

	numbers2 = append(numbers2, 1)
	numbers2 = append(numbers2, 2)
	showSlice(numbers2)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x) // %v = print value
}
