package main

// command : go mod init demo15
import (
	"demo15/employee"
)

func main() {
	emp := employee.Init("Anapat", "Muangchol", 20, 5)
	emp.LeaveRemaining()

	emp = employee.Init("Anapat", "Muangfoam", 20, 5)
	emp.LeaveRemaining() // จะยังเป็นค่าเดิม เนื่องจากที่ Constructor มีการใส่ logic ไว้
}
