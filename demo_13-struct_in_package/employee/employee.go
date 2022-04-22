package employee

import "fmt"

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeaveRemaining() {
	fmt.Printf("%s %s has %d leaves remining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
