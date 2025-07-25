package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	// 1. Declaración de variables
	// Declaración e inicialización de variables
	var x int
	x = 8

	// Declaración e inicialización de otra variable
	y := 7

	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// 2. Captura de errores
	newNumber, err := strconv.ParseInt("NaN", 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed number:", newNumber)
	}

	// 3. Uso de Maps
	phoneDirectory := make(map[string]int)
	phoneDirectory["Alice"] = 123456789
	phoneDirectory["Bob"] = 987654321
	fmt.Println("\nPhone Directory:", phoneDirectory)

	// 4. Uso de Slices (números primos)
	fmt.Println("\nNúmeros primos:")
	primes := []int{2, 3, 5, 7, 11, 13}

	primes = append(primes, 17, 19) // Agregar más números primos

	for i, prime := range primes {
		fmt.Printf("Prime number at index %d: %d\n", i, prime)
	}

	// 5. Uso Goroutines y canales
	fmt.Println("\nIniciando goroutine...")
	c := make(chan int) // Canal para comunicación entre goroutines
	go doSomething(c)   // Esta goroutine permite crearla pero no esperar su finalización
	<-c                 // Espera a que la goroutine envíe un valor al canal
	fmt.Println("Fin")

	// 6. Uso de Punteros
	var a int = 10
	var b *int = &a
	fmt.Println("\nValor de a:", a)
	fmt.Println("Dirección de a:", &a)
	fmt.Println("Valor de b:", *b)
	fmt.Println("Dirección de b:", b)
	fmt.Println(("A es igual a B?"), a == *b)
}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second) // Simula una tarea que toma tiempo
	fmt.Println("Tarea completada")
	c <- 42 // Envía un valor al canal
}
