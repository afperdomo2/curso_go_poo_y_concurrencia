package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	// Crea una instancia de Employee con valores por defecto
	e := Employee{}
	fmt.Println(e)

	// Asignando valores a los atributos directamente
	e.id = 1
	e.name = "John Doe"
	fmt.Println(e)

	// Usando métodos para modificar los atributos
	e.SetId(2)
	e.SetName("Pepito Perez")
	fmt.Println(e)

	// Usando métodos para obtener los valores de los atributos
	fmt.Println("ID:", e.GetId())
	fmt.Println("Name:", e.GetName())
}
