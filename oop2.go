package main

import "fmt"

type Employee struct {
	name string
	ID   int
}

func (e Employee) Work() {
	fmt.Printf("Employee name:%s, ID:%d\n", e.name, e.ID)
}

type Manager struct {
	Employee
	Department string
}

func main() {
	m := Manager{
		Employee: Employee{
			name: "Kamila",
			ID:   1603,
		},
		Department: "Information Systems",
	}
	m.Work()
}
