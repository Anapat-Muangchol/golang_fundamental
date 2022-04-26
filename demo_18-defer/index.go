package main

import "fmt"

func main() {
	// Defer จะทำงานทีหลังสุด (คล้ายๆ OnDestroy หรือ finally) โดยสามารถมีได้หลายบรรทัด การทำงานจะทำเป็น Stack อันไหนถูกเรียกก่อน จะทำงานทีหลัง

	for i := 0; i < 10; i++ {
		defer printFinish(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println("Start:", i)
	}

	defer fmt.Println("End")

}

func printFinish(i int) {
	fmt.Println("Finish:", i)
}
