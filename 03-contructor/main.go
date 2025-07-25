package main

import "fmt"

type Employee struct {
	id       int
	name     string
	isActive bool
}

func NewEmployee(id int, name string, isActive bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		isActive: isActive,
	}
}

func main() {
	fmt.Println("1. Creando una instancia de Employee con valores por defecto")
	e := Employee{}
	fmt.Println(e)

	fmt.Println("\n2. Asignando valores directamente")
	e2 := Employee{
		id:       1,
		name:     "John Doe",
		isActive: true,
	}
	fmt.Println(e2)

	fmt.Println("\n3. Usando new para crear una instancia")
	e3 := new(Employee)
	fmt.Println(*e3)

	fmt.Println("\n4. Usando el constructor NewEmployee")
	e4 := NewEmployee(2, "Pepito Perez", true)
	fmt.Println(*e4)
}
