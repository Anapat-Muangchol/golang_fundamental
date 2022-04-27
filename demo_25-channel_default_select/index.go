package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	quit := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%s : time %d\n", <-c, (i + 1))
		}
		//quit <- "end"	// Comment ออก เพื่อทดสอบว่า ถ้าไม่มีงานของ quit ส่งไป แล้วงานของ c หมดแล้ว จะทำงานในส่วนของ default
	}()

	task(c, quit)
}

func task(c, quit chan string) {
	// select คือการสุ่มเลือกงานขึ้นมาทำ จากตัวอย่างคือ มีเคส c กับ quit ถ้าตัวไหนมีงาน แต่อีกตัวไม่มี จะทำตัวที่มีงาน, แต่ถ้ามีงานทั้งคู่ จะทำการสุ่ม ทำโดย balance กัน, ถ้าไม่มีงานทั้งคู่ จะเข้าเคส default ทำวนในนี้เรื่อยๆ จนกว่าจะมีงานอื่นๆเขามา
	for { // forever loop (while true)
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return // ออกจาก loop และจบการทำงาน
		default:
			// ถ้าไม่มีงานอื่นๆเลย จะมาทำในนี้
			fmt.Println("waiting. . .")
			time.Sleep(1 * time.Second)
		}
	}
}
