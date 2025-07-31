package main

import "time"

type Person struct {
	dni  string
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
}

var GetPersonByDni = func(dni string) (Person, error) {
	time.Sleep(2 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(2 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee = FullTimeEmployee{}
	person, err := GetPersonByDni(dni)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = person

	employee, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = employee

	return ftEmployee, nil
}
