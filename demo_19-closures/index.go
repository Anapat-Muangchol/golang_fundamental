package main

import "fmt"

// Closures คือ Function ที่ return function โดยที่สามารถมี local variable ภายใน function นั้นๆได้ด้วย ดังตัวอย่าง
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newNextInt := intSeq()    // คือว่าเป็นคนละ Instance กับตัวด้านบน ตัวแปรคนละตัวกัน ทำให้ไม่นับต่อ
	fmt.Println(newNextInt()) // 1
	fmt.Println(newNextInt()) // 2

	nextString := stringSeq()
	fmt.Println(nextString()) // y = 1
	fmt.Println(nextString()) // y = 2
	fmt.Println(nextString()) // y = 3

	fmt.Println(stringSeq()()) // y = 1
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		return fmt.Sprintf("y = %d", y)
	}
}
