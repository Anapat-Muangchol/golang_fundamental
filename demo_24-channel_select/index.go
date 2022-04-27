package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s : time %d\n", <-c, (i + 1))
		}
		quit <- "end"
	}()

	task(c, quit)
}

func task(c, quit chan string) {
	// select คือการสุ่มเลือกงานขึ้นมาทำ จากตัวอย่างคือ มีเคส c กับ quit ถ้าตัวไหนมีงาน แต่อีกตัวไม่มี จะทำตัวที่มีงาน, แต่ถ้ามีงานทั้งคู่ จะทำการสุ่ม ทำโดย balance กัน, ถ้าไม่มีงานทั้งคู่ จะเกิด deadlock
	for { // forever loop (while true)
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return // ออกจาก loop และจบการทำงาน
		}
	}
}
