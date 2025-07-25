package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id     int
	salary float64
}

type FullTimeEmployee struct {
	Employee
	Person
	isActive bool
}

func main() {
	ftEmployee1 := FullTimeEmployee{
		Employee: Employee{
			id:     1,
			salary: 50000,
		},
		Person: Person{
			name: "John Doe",
			age:  30,
		},
		isActive: true,
	}
	fmt.Println(ftEmployee1)

	ftEmployee2 := FullTimeEmployee{}
	ftEmployee2.id = 2
	ftEmployee2.salary = 60000
	ftEmployee2.name = "Jane Smith"
	ftEmployee2.age = 28
	ftEmployee2.isActive = true
	fmt.Println(ftEmployee2)
	fmt.Println("ID:", ftEmployee2.id)
}
