package main

import "fmt"

func main() {
	fnFor()
	fnWhile()
	fnEach()
	fnWhileUsingBreak()
}

func fnFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("for index %d\n", i)
	}
}

func fnWhile() {
	index := 0
	for index < 5 {
		fmt.Printf("while index %d\n", index)
		index++
	}
}

func fnEach() {
	courses := []string{"Android", "iOS", "React", "Java"}

	// with index
	for index, item := range courses {
		fmt.Printf("%d. %s\n", index+1, item)
	}

	// not use index, (shortcut : forr)
	for _, item := range courses {
		fmt.Printf("%s\n", item)
	}

}

func fnWhileUsingBreak() {
	index := 0
	// while true
	for {
		if index > 5 {
			break
		}
		fmt.Printf("while Using Break index %d\n", index)
		index++
	}
}
