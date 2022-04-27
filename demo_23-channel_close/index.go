package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go routine1(ch, 10)

	/*
		// ตัวอย่างการหยุดรอ จนกว่า routine จะ close ด้วย forever loop
		for true {
			value, ok := <-ch	// callback รอจนกว่าจะมีการ send
			if !ok {	// ถ้า close channel แล้วจะเข้า if นี้
				fmt.Println("No more data")
				break
			}

			fmt.Println(value)
		}
	*/

	// ตัวอย่างการหยุดรอ จนกว่า routine จะ close ด้วย range loop (ควรใช้แบบนี้มากกว่า โค้ดดู clean กว่าด้านบน)
	for value := range ch {
		fmt.Println(value)
	}
	fmt.Println("No more data")

}

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	close(c)
}
