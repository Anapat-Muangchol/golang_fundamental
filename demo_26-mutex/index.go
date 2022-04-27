package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to use concurrently.
// Concept เดียวกับ Thread safe ใน Java
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Increase(key string) {
	c.mux.Lock() // สั่ง Lock ไว้เพื่อให้มีเพียงแค่ go-routine ตัวที่เรียกใช้ function นี้ ตัวเดียวเท่านั้น เข้าถึง map ได้
	c.v[key]++
	c.mux.Unlock() // สั่งปลดล็อค ให้ go-routine ตัวอื่นๆ สามารถใช้งาน map ต่อได้
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()         // สั่ง Lock ไว้เพื่อให้มีเพียงแค่ go-routine ตัวที่เรียกใช้ function นี้ ตัวเดียวเท่านั้น เข้าถึง map ได้
	defer c.mux.Unlock() // สั่งปลดล็อค ให้ go-routine ตัวอื่นๆ สามารถใช้งาน map ต่อได้ (มี defer ทำให้คำสั่งนี้ทำงานทีหลัง หลังจาก return)
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Increase("some-key")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("some-key"))
}
