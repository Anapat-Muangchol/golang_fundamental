package main

// command : go mod init demo16
import (
	"demo17/shape"
	"fmt"
	"strconv"
)

func main() {
	testCast()

	r := shape.NewRectangle(10, 20)
	shape.CastToRectangle(r)

	c := shape.NewCircle(10)
	shape.CastToRectangle(c)
}

func testCast() {
	var index int8 = 15
	var bigIndex int32 = int32(index)
	fmt.Println(bigIndex)

	yesterdayString := "50" // Success Case
	// yesterdayString := "50asdf" // Error Case
	yesterday, err := strconv.Atoi(yesterdayString)
	if err != nil {
		fmt.Printf("Error cast string to int : %s\n", yesterdayString)
	} else {
		fmt.Printf("Success cast string to int : %d\n", yesterday)
	}
}
