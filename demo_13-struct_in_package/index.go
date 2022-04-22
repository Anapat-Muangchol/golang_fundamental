package main

// command : go mod init demo13
import (
	"demo13/employee"
)

func main() {
	emp := employee.Employee{
		FirstName:   "Anapat",
		LastName:    "Muangchol",
		TotalLeaves: 20,
		LeavesTaken: 5,
	}

	emp.LeaveRemaining()
}
