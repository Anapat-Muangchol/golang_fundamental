package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	// หาก send เกิน buffer ที่ตั้งไว้ แต่อยู่ใน routine จะไม่ error แต่จะทำถึงแค่บรรทัดที่ send เข้าไปเต็ม buffer เท่านั้น หลังจากนั้นจะไม่ทำงานต่อ
	go routine(ch, 1)

	// ----- ตัวอย่าง เรียกหลายๆ routine -----
	//go routine1(ch, 1)
	//go routine2(ch, 2)
	//go routine3(ch, 3)

	//fmt.Println("received: ", <-ch)
	//fmt.Println("received: ", <-ch)
	//fmt.Println("received: ", <-ch)

	time.Sleep(5 * time.Second) // หยุดรอ เพื่อดูผลลัพธ์เฉยๆ
}

func routine(c chan int, payload int) {
	c <- payload
	fmt.Println("routine : send", payload)

	payload++
	c <- payload
	fmt.Println("routine : send", payload)

	payload++
	c <- payload
	fmt.Println("routine : send", payload)
}

func routine1(c chan int, payload int) {
	c <- payload
	fmt.Println("routine1 : send", payload)
}

func routine2(c chan int, payload int) {
	c <- payload
	fmt.Println("routine2 : send", payload)
}

func routine3(c chan int, payload int) {
	c <- payload
	fmt.Println("routine3 : send", payload)
}
