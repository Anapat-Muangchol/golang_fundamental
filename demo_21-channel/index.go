package main

import (
	"fmt"
	"time"
)

func main() {
	// Channel คือตัวแปร สำหรับส่งข้อมูลข้าม Concurrency

	ch := make(chan int, 1) // จอง buffer ไว้ 1 (จะสามารถ send ได้แค่ 1 ครั้ง) ห้าม send เกิน buffer, จะต้อง receive ออกก่อนเพื่อไม่ให้ buffer เต็ม ถึงจะส่งเข้าไปเพิ่มได้, ห้าม receive ตอนที่ยังไม่มีข้อมูลใน buffer | หากไม่ทำตามกฎนี้ จะเกิดปัญหา deadlock
	ch <- 1                 // send (คล้ายๆการเอา value เข้าคิวไว้)

	i := <-ch // receive (คล้ายๆการเอา value ออกมาจากคิว, เรียงลำดับ เข้าก่อน ออกก่อน)
	fmt.Println(i)

	ch <- 2           // send
	fmt.Println(<-ch) // สามารถ receive แบบนี้ได้เลย เป็น short form

	time.Sleep(1 * time.Second) // หยุดรอ เพื่อดูผลลัพธ์เฉยๆ
}
