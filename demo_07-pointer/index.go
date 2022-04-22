package main

import "fmt"

func main() {
	msg := "Some Message"
	fmt.Printf("Message is '%s'\n", msg)

	var msgPointer *string = &msg
	fmt.Printf("Pointer address : %p, Value : '%s'\n", msgPointer, *msgPointer)

	changeMessage(msgPointer)
	fmt.Printf("Pointer address : %p, Value : '%s'\n", msgPointer, *msgPointer)

	changeMessageV2(&msg, "New Message V2")
	fmt.Printf("New Value : '%s'\n", msg)
	fmt.Printf("Pointer address : %p, Value : '%s'\n", msgPointer, *msgPointer)
}

func changeMessage(pointer *string) {
	*pointer = "New Message"
}

func changeMessageV2(pointer *string, newMessage string) {
	*pointer = newMessage
}
