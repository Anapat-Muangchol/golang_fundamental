package main

// command : go mod init demo14
import (
	"demo14/employee"
)

func main() {
	emp := employee.New("Anapat", "Muangchol", 20, 5)
	emp.LeaveRemaining()
}
