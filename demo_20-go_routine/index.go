package main

import (
	"fmt"
	"time"
)

func main() {
	// Concurrency คือ ยังมีแค่ thread เดียว หรือ cpu core เดียวอยู่ แต่หากงานที่รันอยู่มีการรอ เช่น วนลูป ในระหว่างรอนี้ จะแบ่งไปให้ process อื่นๆทำงานต่อได้ (ตัวที่จะคอยจัดการให้คือ go-routine)
	go run1()
	go run2()

	time.Sleep(5 * time.Second) // หยุดรอ เพื่อดูผลลัพธ์เฉยๆ
}

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run something 1 :", i)
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run something 2 :", i)
	}
}
