package main

import "fmt"

func main() {
	isEqualTen(10)
	fmt.Println()

	isBetweenZeroToTen(11)
	fmt.Println()

	number := 9
	if result := isOddEven(number); result == "Even" {
		fmt.Printf("%d : Is Even\n", number)
	} else if result == "Odd" {
		fmt.Printf("%d : Is Odd\n", number)
	} else {
		fmt.Printf("%d : Is failed\n", number)
	}
	fmt.Println()

	for i := 0; i < 7; i++ {
		fnSwitchCase(i)
	}
}

func isEqualTen(value int) {
	if value == 10 {
		fmt.Println("Equal 10")
	} else {
		fmt.Println("Not Equal 10")
	}
}

func isBetweenZeroToTen(value int) {
	if value >= 0 && value <= 10 {
		fmt.Printf("%d : Is Between 0 - 10\n", value)
	} else {
		fmt.Printf("%d : Is Not Between 0 - 10\n", value)
	}
}

func isOddEven(value int) string {
	if value%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func fnSwitchCase(number int) {
	switch number {
	case 0:
		fmt.Printf("%d : Is Zero\n", number)
	case 1:
		fmt.Printf("%d : Is One\n", number)
	case 2:
		fmt.Printf("%d : Is Two\n", number)
	case 3:
		fmt.Printf("%d : Is Three\n", number)
	case 4:
		fmt.Printf("%d : Is Four\n", number)
	case 5:
		fmt.Printf("%d : Is Five\n", number)
	default:
		fmt.Printf("%d : Is Not Know!\n", number)
	}
}
