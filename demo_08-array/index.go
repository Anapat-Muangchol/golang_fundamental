package main

import "fmt"

func main() {
	var array1 []int = []int{1, 2, 3, 4, 5}
	fmt.Print("array1 : ")
	printIntegerArray(array1)

	var array2 = []int{1, 2, 3}
	fmt.Print("array2 : ")
	printIntegerArray(array2)

	var array3 [2]string
	fmt.Print("array3 : ")
	for _, item := range array3 {
		fmt.Print(item, ", ")
	}

	array3[0], array3[1] = "android", "ios"
	fmt.Print("\narray3 : ")
	for _, item := range array3 {
		fmt.Print(item, ", ")
	}

	// printStringArray(array3)
}

func printIntegerArray(array []int) {
	fmt.Print("{")
	for i, value := range array {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(value)
	}
	fmt.Print("}\n")
}

// func printStringArray(array []string) {
// 	fmt.Print("{")
// 	for i, value := range array {
// 		if i > 0 {
// 			fmt.Print(", ")
// 		}
// 		fmt.Print(value)
// 	}
// 	fmt.Print("}\n")
// }
