package main

type PrintInfo interface {
	getMessage() string
}

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
	endDate string
}

func (e FullTimeEmployee) getMessage() string {
	return "Full-time employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (e TemporaryEmployee) getMessage() string {
	return "Temporary employee"
}

func getMessage(info PrintInfo) string {
	return info.getMessage()
}

func main() {
	ftEmployee2 := FullTimeEmployee{}
	ftEmployee2.id = 2
	ftEmployee2.salary = 60000
	ftEmployee2.name = "Jane Smith"
	ftEmployee2.age = 28

	tEmployee := TemporaryEmployee{}
	tEmployee.Person = Person{name: "Alice Johnson", age: 25}
	tEmployee.Employee = Employee{id: 3, salary: 40000.0}
	tEmployee.taxRate = 15

	println(getMessage(ftEmployee2))
	println(getMessage(tEmployee))
}
