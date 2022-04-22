package employee

import "fmt"

type employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func New(firstName string, lastName string, totalLeaves int, leavesTaken int) employee {
	return employee{
		FirstName:   firstName,
		LastName:    lastName,
		TotalLeaves: totalLeaves,
		LeavesTaken: leavesTaken,
	}
}

func (e employee) LeaveRemaining() {
	fmt.Printf("%s %s has %d leaves remining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
